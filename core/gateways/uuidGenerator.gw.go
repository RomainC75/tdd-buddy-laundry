package gateways

import "github.com/google/uuid"

type IUuidGenerator interface {
	Generate() uuid.UUID
}
