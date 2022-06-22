package reporter

type CommitResponse struct {
	Title    string   `json:"title"`
	Message  string   `json:"message"`
	Author   Author   `json:"author"`
	Added    []string `json:"added"`
	Modified []string `json:"modified"`
	Removed  []string `json:"removed"`
}

type ProjectResponse struct {
	Name string `json:"name"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
