package services

import (
	"fmt"
)

type FakeConsole struct {
}

func NewFakeConsole() *FakeConsole {
	return &FakeConsole{}
}

func (fc *FakeConsole) BlockHost(host string) error {
	fmt.Printf("Blocking... %s ", host)
	return nil
}
