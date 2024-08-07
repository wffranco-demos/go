package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Get(uri string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", base, uri))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetUser(id int) (*User, error) {
	body, err := Get(fmt.Sprintf("users/%d", id))
	if err != nil {
		return nil, err
	}
	var data UserData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data.User, nil
}
