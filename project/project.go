package project

type Project struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	RepoURL []string `mapstructure:"repo_url"`
}
