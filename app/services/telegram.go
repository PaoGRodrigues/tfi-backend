package services

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Telegram struct {
	Token       string
	Username    string
	TelegramBot *tgbotapi.BotAPI
}

func NewTelegramInterface() *Telegram {
	return &Telegram{}
}

func (t *Telegram) SendMessage(message string) error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := t.TelegramBot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.From.UserName != t.Username {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		if _, err := t.TelegramBot.Send(msg); err != nil {
			return err
		}
	}
	return nil
}

func (t *Telegram) Configure(token string, username string) error {
	t.Token = token
	t.Username = username

	bot, err := tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		return err
	}
	t.TelegramBot = bot
	return nil
}
