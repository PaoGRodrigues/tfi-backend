package services

type FakeTool struct{}

func NewFakeTool() *FakeTool {
	return &FakeTool{}
}
