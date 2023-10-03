package chat

import "strings"

type Message struct {
	from *User
	Text []byte
}

type JsonMessage struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func NewMessage(message JsonMessage, user *User) *Message {
	return &Message{
		from: user,
		Text: []byte(strings.TrimSpace(strings.Replace(message.Text, string(newline), string(space), -1))),
	}
}
