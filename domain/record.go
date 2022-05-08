package domain

// TODO: more fields (dateCreated, url, description)
// TODO: add validator
type Record struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
