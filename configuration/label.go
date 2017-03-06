package configuration

// Label represents the label we want to set on Github.
// Name corresponds to the label name.
// Color corresponds to the color of the label.
// Repositories allowds us to set regular expressions for matching our labels.
// If repositories is empty, it will affect all repositories by default.
type Label struct {
	Name         string   `json:"name" validate:"nonzero"`
	Color        string   `json:"color" validate:"regexp=^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"`
	Repositories []string `json:"repositories"`
}
