package dto

type exception struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
}

// NewException return new exception dto
func NewException(code int16, message string) *exception {
	return &exception{
		Code:    code,
		Message: message,
	}
}
