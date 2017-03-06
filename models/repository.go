package models

import "fmt"
import "github.com/MarquisIO/Organisator/network"

// Repository represents Github repository and its data.
type Repository struct {
	Name      string `json:"name"`
	IsPrivate bool   `json:"private"`
	IsFork    bool   `json:"private"`
	Owner     struct {
		Name string `json:"login"`
	}
}

// GetAllLabels gets all existing labels on a Github repository
func (repo *Repository) GetAllLabels(token string) map[string]Label {

	url := fmt.Sprintf("/repos/%v/%v/labels?access_token=%v", repo.Owner.Name, repo.Name, token)

	var labels []Label

	if err := network.GetJSON(url, &labels); err != nil {
		fmt.Printf("Error getting labels for repository %v : %v\n", repo.Name, err)
	}

	labelsMap := make(map[string]Label)
	for _, element := range labels {
		labelsMap[element.Name] = element
	}

	return labelsMap
}
