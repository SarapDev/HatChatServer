package chat

import "github.com/google/uuid"

type User struct {
	id uuid.UUID
	name string
	isRegistered bool
}