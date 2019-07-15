package models

type (
	ChatMessage struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Address string `json:"address"`
	}
)
