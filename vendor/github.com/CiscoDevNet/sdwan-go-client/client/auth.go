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
	if !client.authClient.valid {
		err := client.authenticate()
		if err != nil {
			if err == fmt.Errorf("Invalid Session") {
				err = client.authenticate()
				if err != nil {
					return nil, err
				}
			}
		}
	}

	sessionID := fmt.Sprintf("JESSIONID=%s", client.authClient.jsessionID)
	req.Header.Set("Cookie", sessionID)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-XSRF-TOKEN", client.authClient.token)

	return req, nil
}
