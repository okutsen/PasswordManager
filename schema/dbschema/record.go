package dbschema

// TODO: more fields (dateCreated, url, description)
type Record struct {
	ID       uint64
	Name     string
	Login    string
	Password string
}
