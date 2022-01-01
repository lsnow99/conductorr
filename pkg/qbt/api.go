package qbt

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"sync"
)

type Client struct {
	baseUrl string
	cookies []*http.Cookie
	sync.RWMutex
}

const urlSubpath = "/api/v2"
var HTTPClient = http.DefaultClient

func NewClient(baseUrl string) *Client {
	c := new(Client)
	c.baseUrl = baseUrl
	return c
}

func (c *Client) Login(username, password string) error {
	params := url.Values{}
	params.Add("username", username)
	params.Add("password", password)
	return c.doRequest("auth", "login", params, nil)
}

// func (c *Client) AddTorrent(torrentData []byte) error {
	
// }

func (c *Client) doRequest(apiName, methodName string, params url.Values, obj *interface{}) error {
	c.RLock()
	u, err := url.Parse(path.Join(c.baseUrl, urlSubpath, apiName, methodName))
	c.RUnlock()

	if err != nil {
		return err
	}
	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if len(resp.Cookies()) > 0 {
		c.Lock()
		c.cookies = resp.Cookies()
		c.Unlock()
	}

	if obj != nil {
		return json.NewDecoder(resp.Body).Decode(obj)
	}
	return nil
}
