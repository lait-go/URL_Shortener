package domain

const (
	NOTFOUND  string = "NOT_FOUND"
	KEYEXISTS string = "KEY_EXISTS"
)

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type KeyExistsError struct{}

func (e *KeyExistsError) Error() string {
	return "key already exists"
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return "key not found"
}
