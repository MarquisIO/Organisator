package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Application is our current Github implementation.
type Application struct {
	Token string
}

const (
	rootURL = "https://api.github.com"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func (app *Application) getJSON(url string, object interface{}) error {
	resp, err := httpClient.Get(fmt.Sprintf("%v?access_token=%v", url, app.Token))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(object)
}

// GetAllRepositoriesFromOrganisation gets a list of all repositories.
func (app *Application) GetAllRepositoriesFromOrganisation(name string) (*[]Repository, error) {

	url := fmt.Sprintf("%v/orgs/%v/repos", rootURL, name)

	fmt.Println(url)

	var repos []Repository

	app.getJSON(url, &repos)

	fmt.Printf("Repo : %v\n", repos)
	return &repos, nil
}
