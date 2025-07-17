package providers

type FakePinGenerator struct {
	ExpectedPin string
}

func NewFakePinGenerator() *FakePinGenerator {
	return &FakePinGenerator{}
}

func (fpg *FakePinGenerator) Generate() string {
	return fpg.ExpectedPin
}
