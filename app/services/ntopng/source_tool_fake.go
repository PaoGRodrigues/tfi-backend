package ntopng

type FakeTool struct{}

func NewFakeTool() *FakeTool {
	return &FakeTool{}
}

func (d *FakeTool) SetInterfaceID() error {
	return nil
}
