package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

const (
	rootURL = "https://api.github.com"
)

// GetJSON performs a get call on specific url and fills the object structure.
func GetJSON(url string, object interface{}) error {
	resp, err := httpClient.Get(fmt.Sprintf("%v%v", rootURL, url))
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 404:
		return fmt.Errorf("Ressource for URL %v not found.", url)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(object)
}

// PatchJSON performs a patch call on specific url with the given object.
func PatchJSON(url string, object interface{}) error {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(object)

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%v%v", rootURL, url), b)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	_, err = httpClient.Do(request)

	if err != nil {
		return err
	}

	return nil
}

// PatchJSON performs a post call on specific url with the given object.
func PostJSON(url string, object interface{}) error {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(object)

	_, err := httpClient.Post(fmt.Sprintf("%v%v", rootURL, url), "application/json", b)

	if err != nil {
		return err
	}

	return nil
}
