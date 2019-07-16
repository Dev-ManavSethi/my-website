package models

type (
	ChatMessage struct {
		Name    string `json:"name"`
		Message string `json:"message"`
		Time int64 `json:"time"`
	}

	ChatUser struct {
		IP string `json:"ip"`
		Name string `json:"name"`
		Chats []ChatMessage `json:"chats"`
	}



)


var(
	Chats map[string]ChatUser
)

