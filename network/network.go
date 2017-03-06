package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

const (
	rootURL = "https://api.github.com"
)

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

func PatchJSON(url string, object interface{}) error {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(object)

	request, err := http.NewRequest("PATCH", fmt.Sprintf("%v%v", rootURL, url), b)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	res, err := httpClient.Do(request)

	if err != nil {
		return err
	}

	io.Copy(os.Stdout, res.Body)
	return nil
}

func PostJSON(url string, object interface{}) error {

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(object)

	res, err := httpClient.Post(fmt.Sprintf("%v%v", rootURL, url), "application/json", b)

	if err != nil {
		return err
	}

	io.Copy(os.Stdout, res.Body)
	return nil
}
