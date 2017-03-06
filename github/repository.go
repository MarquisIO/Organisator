package github

// Repository represents Github repository and its data.
type Repository struct {
	Name      string `json:"name"`
	IsPrivate bool   `json:"private"`
	IsFork    bool   `json:"private"`
}
