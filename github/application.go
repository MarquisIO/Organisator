package github

import (
	"fmt"

	"regexp"

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

	var repos []models.Repository

	if err := network.GetJSON(url, &repos); err != nil {
		return nil, err
	}

	return repos, nil
}

// Start begins logic
func (app *Application) Start() error {
	fmt.Println("Start github logic.")
	repositories, err := app.getAllRepositoriesFromOrganisation()
	if err != nil {
		return err
	}

	for _, repository := range repositories {
		// Current labels of a repository
		fmt.Println("Actions on repository : ", repository.Name)
		app.labelsActions(repository)
	}

	return nil
}

func (app *Application) labelsActions(repository models.Repository) {

	currentLabels := repository.GetAllLabels(app.Token)

	for _, labelToAdd := range app.Config.Labels {
		labelToAdd.Information.Repository = repository.Name
		if len(labelToAdd.Repositories) == 0 { // No rule, accept all repo
			app.updateLabel(currentLabels, &labelToAdd)
		} else {
			for _, repositoyRegex := range labelToAdd.Repositories {
				if matched, err := regexp.MatchString(repositoyRegex, repository.Name); matched {
					app.updateLabel(currentLabels, &labelToAdd)
					break
				} else if err != nil {
					fmt.Printf("\tError matching regex %v : %v\n", repositoyRegex, err)
					continue
				}
				fmt.Printf("\tRepository %v not matching any regexp\n", repository.Name)
			}
		}
	}
}

func (app *Application) updateLabel(currentLabels map[string]models.Label, label *models.Label) {
	if _, isThere := currentLabels[label.Name]; isThere {
		fmt.Println("\tLabel", label.Name, "exists in", label.Information.Repository, "update it.")
		label.Update(app.Token)
	} else {
		fmt.Println("\tCreate label", label.Name, "in", label.Information.Repository)
		label.Create(app.Token)
	}
}
