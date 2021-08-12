package chat

import "bytes"

type Message struct {
	from *User
	Text [] byte
}

var (
	newline = []byte{'\n'}
	space 	= []byte{' '}
)

func NewMessage (message []byte, user *User) *Message {
	return &Message{
		from: user,
		Text: bytes.TrimSpace(bytes.Replace(message, newline, space, -1)),
	}
}