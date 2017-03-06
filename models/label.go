package models

import (
	"fmt"

	"github.com/MarquisIO/Organisator/network"
)

// Label represents Github label object
type Label struct {
	Name         string   `json:"name" validate:"nonzero"`
	Color        string   `json:"color" validate:"regexp=^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"`
	Repositories []string `json:"repositories"`
	Information  struct {
		Organisation string
		Repository   string
	}
}

// Create creates new label on github repository
func (label *Label) Create(token string) {
	url := fmt.Sprintf("/repos/%v/%v/labels?access_token=%v", label.Information.Organisation, label.Information.Repository, token)
	if err := network.PostJSON(url, label); err != nil {
		fmt.Printf("Error creating label %v for repository %v : %v\n", label.Name, label.Information.Repository, err)
	}
}

// Update updates existing label on github repository
func (label *Label) Update(token string) {
	url := fmt.Sprintf("/repos/%v/%v/labels/%v?access_token=%v", label.Information.Organisation, label.Information.Repository, label.Name, token)
	if err := network.PatchJSON(url, label); err != nil {
		fmt.Printf("Error creating label %v for repository %v : %v\n", label.Name, label.Information.Repository, err)
	}
}
