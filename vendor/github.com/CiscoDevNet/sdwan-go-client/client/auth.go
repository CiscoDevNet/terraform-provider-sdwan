package client

import (
	"fmt"
	"log"
	"net/http"
)

type auth struct {
	valid      bool
	jsessionID string
	token      string
}

func (client *Client) injectAuthenticationHeader(req *http.Request, path string) (*http.Request, error) {
	log.Println("[DEBUG] Begin Injection")

	client.mutex.Lock()
	if !client.authClient.valid {
		err := client.authenticate()
		if err != nil {
			if err.Error() == "Invalid Session" {
				err = client.authenticate()
				if err != nil {
					client.mutex.Unlock()
					return nil, err
				}
			} else {
				client.mutex.Unlock()
				return nil, err
			}
		}
	}
	client.mutex.Unlock()

	sessionID := fmt.Sprintf("JSESSIONID=%s", client.authClient.jsessionID)
	req.Header.Set("Cookie", sessionID)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-XSRF-TOKEN", client.authClient.token)

	return req, nil
}
