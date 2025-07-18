package providers

import "github.com/google/uuid"

type FakeUuidGenerator struct {
	ExpectedUuid uuid.UUID
}

func NewFakeUuidGenerator() *FakeUuidGenerator {
	return &FakeUuidGenerator{}
}

func (fug *FakeUuidGenerator) Generate() uuid.UUID {
	return fug.ExpectedUuid
}
