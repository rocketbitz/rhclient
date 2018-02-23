package rhclient

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (rh *Robinhood) Login(username, password string) error {
	uri := fmt.Sprintf("%v/api-token-auth/", baseURL)
	m := map[string]interface{}{
		"username": username,
		"password": password,
	}
	resp, err := rh.post(uri, m)
	if err != nil {
		return err
	}
	m = map[string]interface{}{}
	if err = json.Unmarshal(resp.Body(), &m); err != nil {
		return err
	}
	if m["token"] == nil {
		return errors.New(string(resp.Body()))
	}
	rh.token = m["token"].(string)
	return nil
}

func (rh *Robinhood) Logout() error {
	uri := fmt.Sprintf("%v/api-token-auth/", baseURL)
	_, err := rh.post(uri, nil)
	if err != nil {
		return err
	}
	rh.token = ""
	return nil
}
