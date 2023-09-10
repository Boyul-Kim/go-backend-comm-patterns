package util

import (
	"github.com/google/uuid"
)

func RandomIdGenerator() string {
	id := uuid.New()
	return id.String()
}
