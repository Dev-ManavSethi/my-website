package models

type (
	ChatMessage struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		//Address string `json:"address"`
		Time int64 `json:"time"`
	}

	ChatUser struct {
		IP string `json:"ip"`
		Name string `json:"name"`
		Chats []ChatMessage `json:"chats"`
	}
)
