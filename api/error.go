package api

type Error struct {
	Error  string `json:"error"`
	Status string `json:"status"`
}

func NewError(msg string) *Error {
	return &Error{Error: msg, Status: "fail"}
}
