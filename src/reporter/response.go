package reporter

type CommitResponse struct {
	Message  string         `json:"message"`
	Author   AuthorResponse `json:"author"`
	Added    []string       `json:"added"`
	Modified []string       `json:"modified"`
	Removed  []string       `json:"removed"`
}

type SimpleCommitResponse struct {
	Message string         `json:"message"`
	Author  AuthorResponse `json:"author"`
}

type ProjectResponse struct {
	Name string `json:"name"`
}

type AuthorResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BuildResponse struct {
	Stage  string `json:"stage"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
