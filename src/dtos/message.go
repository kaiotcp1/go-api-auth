package dtos

type Message struct {
	Message string `json:"message,omitempty"`
}

type Error struct {
	Message string `json:"message"`
}
