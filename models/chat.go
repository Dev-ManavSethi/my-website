package models

type (
	ChatMessage struct {
		IP string `json:"ip"`
		Name    string `json:"name"`
		Message string `json:"message"`
		Time int64 `json:"time"`
	}

	User struct {
		IP string `json:"ip"`
		Name string `json:"name"`
		Chats []ChatMessage `json:"chats"`
		VisitCount int `json:"visit_count"`
		VisitMoreThanOnce bool `json:"visit_more_than_once"`
	}

	FirstTime struct {
		FirstTime bool
	}

)


var(
	Chats map[string]User
)

