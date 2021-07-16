package chat

import "github.com/google/uuid"

type Room struct {
	id uuid.UUID
	users map[*User]bool

}