package apischema

// TODO: add validator
// TODO: add uuid (request id)
type Record struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Error struct {
	Message string `json:"message"`
}
