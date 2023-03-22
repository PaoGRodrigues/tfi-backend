package services

import "fmt"

type FakeBot struct {
}

func (bot *FakeBot) SendMessage(e string) error {
	fmt.Printf("Mensaje enviado")
	return nil
}
