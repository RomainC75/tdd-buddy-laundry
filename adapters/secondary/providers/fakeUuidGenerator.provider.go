package providers

type FakeUuidGenerator struct {
	ExpectedUuid string
}

func NewFakeUuidGenerator() *FakeUuidGenerator {
	return &FakeUuidGenerator{}
}

func (fug *FakeUuidGenerator) Generate() string {
	return fug.ExpectedUuid
}
