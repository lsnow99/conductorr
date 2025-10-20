package integration

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

/*
 * Follow a download URL, stopping when we reach a magnet URL or when we
 * find a torrent file.
 *
 * Either:
 * - We find a magnet URL: nextURL.Scheme == "magnet", and *respBuf should
 *   be discarded, or,
 * - We find a torrent file: nextURL.Scheme != "magnet", and *respBuf
 *   contains the torrent file.
 */
func FollowDownloadURL(nextURL *url.URL, respBuf *[]byte) error {
	// override default http redirect follow behavior
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	// perform redirect following manually. stop on magnet url.
	var resp *http.Response = nil
	maxFollows := 100
	for i := 0; i < maxFollows; i++ {
		if nextURL.Scheme == "magnet" {
			break
		}

		if nextURL.Scheme != "http" && nextURL.Scheme != "https" {
			return fmt.Errorf("Unknown protocol '%s' while following redirect", nextURL.Scheme)
		}

		resp, err := client.Get(nextURL.String())
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != 301 && resp.StatusCode != 308 {
			*respBuf, err = io.ReadAll(resp.Body)
			if err == nil {
				return err
			}

			break
		}

		nextURL, err = url.Parse(resp.Header.Get("Location"))
		if err != nil {
			return err
		}

	}

	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return nil
	} else {
		return fmt.Errorf("Unssuccessful HTTP response")
	}
}
