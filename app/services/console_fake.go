package services

import (
	"fmt"
)

type FakeConsole struct {
}

func NewFakeConsole() *FakeConsole {
	return &FakeConsole{}
}

func (fc *FakeConsole) Block(host string) (*string, error) {
	fmt.Printf("Blocking... %s ", host)
	return &host, nil
}
