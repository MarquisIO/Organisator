package models

import (
	"fmt"

	"github.com/Shakarang/Orgmanager/network"
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

func (label *Label) Create(token string) {
	url := fmt.Sprintf("/repos/%v/%v/labels?access_token=%v", label.Information.Organisation, label.Information.Repository, token)
	fmt.Println(url)
	if err := network.PostJSON(url, label); err != nil {
		fmt.Printf("Error creating label %v for repository %v : %v\n", label.Name, label.Information.Repository, err)
	}
}

func (label *Label) Update(token string) {
	url := fmt.Sprintf("/repos/%v/%v/labels/%v?access_token=%v", label.Information.Organisation, label.Information.Repository, label.Name, token)
	fmt.Println(url)
	if err := network.PatchJSON(url, label); err != nil {
		fmt.Printf("Error creating label %v for repository %v : %v\n", label.Name, label.Information.Repository, err)
	}
}
