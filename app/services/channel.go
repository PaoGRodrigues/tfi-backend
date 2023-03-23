package services

type Channel interface {
	SendMessage(string) error
}

type Telegram struct {
}

func NewTelegramInterface() *Telegram {
	return &Telegram{}
}

func (t *Telegram) SendMessage(e string) error {

	return nil
}
