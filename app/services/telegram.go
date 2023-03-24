package services

type Telegram struct {
	Token    string
	Username string
}

func NewTelegramInterface() *Telegram {
	return &Telegram{}
}

func (t *Telegram) SendMessage(e string) error {

	return nil
}

func (t *Telegram) Configure(token string, username string) error {
	return nil
}
