package plex

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type AuthResponse struct {
	User struct {
		AuthToken string `json:"authToken"`
	} `json:"user"`
}

func RetrievePlexToken(username, password string) (string, error) {
	u := url.URL{}
	u.Host = "plex.tv"
	u.Scheme = "https"
	u.Path = "/users/sign_in.json"

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("X-Plex-Client-Identifier", "conductorr")
	req.Header.Add("Authorization", "Basic "+base64.RawStdEncoding.EncodeToString([]byte(username+":"+password)))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("received non-200 level status code %d", resp.StatusCode)
	}

	authResp := AuthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return "", err
	}

	return authResp.User.AuthToken, nil
}
