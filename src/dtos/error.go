package dtos

type APIError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"error,omitempty"`
}

func (error APIError) Error() string {
	return error.Message
}
