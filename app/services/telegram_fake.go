package services

import "fmt"

type FakeBot struct {
}

func NewFakeBot() *FakeBot {
	return &FakeBot{}
}

func (bot *FakeBot) SendMessage(e string) error {
	fmt.Printf("Mensaje enviado")
	return nil
}
