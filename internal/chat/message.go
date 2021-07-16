package chat

type Message struct {
	from *User
	room *Room
	Text string `json:"text"`
}