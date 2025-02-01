package models

type GitHubIssue struct {
	ID         int    `json:"id"`
	Number     int    `json:"number"`
	Title      string `json:"title"`
	Body       string `json:"body"`
	State      string `json:"state"`
	HTMLURL    string `json:"html_url"`
	Repository struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Owner struct {
			AvatarURL string `json:"avatar_url"`
		} `json:"owner"`
	} `json:"repository"`
	User struct {
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
}
