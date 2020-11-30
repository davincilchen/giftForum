package gentoken

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func GenUUIDv4() uuid.UUID {
	id := uuid.NewV4()
	return id
}

func GenUUIDv4String() string {
	token := GenUUIDv4()
	return fmt.Sprintf("%s", token)
}

func GenToken() string {
	return GenUUIDv4String()
}
