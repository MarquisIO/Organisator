package github

import (
	"fmt"
	"os"

	"github.com/Shakarang/Orgmanager/configuration"
	"github.com/Shakarang/Orgmanager/models"
	"github.com/Shakarang/Orgmanager/network"
)

// Application is our current Github implementation.
type Application struct {
	Token        string
	Organisation string
	Config       *configuration.Configuration
}

// getAllRepositoriesFromOrganisation gets a list of all repositories.
func (app *Application) getAllRepositoriesFromOrganisation() ([]models.Repository, error) {

	url := fmt.Sprintf("/orgs/%v/repos?access_token=%v", app.Organisation, app.Token)

	fmt.Println(url)

	var repos []models.Repository

	if err := network.GetJSON(url, &repos); err != nil {
		return nil, err
	}

	fmt.Printf("Repo : %v\n", repos)
	return repos, nil
}

// Start begins logic
func (app *Application) Start() error {

	repositories, err := app.getAllRepositoriesFromOrganisation()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, repository := range repositories {
		// Current labels of a repository
		currentLabels := repository.GetAllLabels(app.Token)
		fmt.Println(repository)
		for _, labelToAdd := range app.Config.Labels {

			// TODO Check regexp here

			labelToAdd.Information.Repository = repository.Name
			if _, isThere := currentLabels[labelToAdd.Name]; isThere {
				// Label already present, will update it
				fmt.Printf("Update label %v\n", labelToAdd.Name)
				labelToAdd.Update(app.Token)
			} else {
				// Create label
				fmt.Printf("Create label %v\n", labelToAdd.Name)
				labelToAdd.Create(app.Token)
			}
		}

	}
	return nil
}
