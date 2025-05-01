package telegram

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	Token       string
	Username    string
	ChatID      int64
	ChatIDMutex *sync.Mutex
	TelegramBot *tgbotapi.BotAPI
}

func NewTelegramInterface() *Telegram {
	return &Telegram{
		ChatIDMutex: &sync.Mutex{},
	}
}

func (t *Telegram) SendMessage(message string) error {
	t.ChatIDMutex.Lock()
	defer t.ChatIDMutex.Unlock()

	// To send messages, Telegram need the ChatID that must be get when the user
	// sends a message to the bot.
	if t.ChatID == 0 {
		return nil
	}
	msg := tgbotapi.NewMessage(t.ChatID, message)

	if _, err := t.TelegramBot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (t *Telegram) Configure(token string, username string) error {
	t.Token = token
	t.Username = username
	t.ChatID = 0

	bot, err := tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		return err
	}
	t.TelegramBot = bot

	// This func waits the user to send a msg to the bot. When the user sends the message,
	// takes the ChatID and puts it as an attr in the struct.
	go t.waitForTelegramChatID()

	return nil
}

// Func to Get the ChatID getting updates from the chat with the user. Listener.
func (t *Telegram) waitForTelegramChatID() {
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
		t.ChatIDMutex.Lock()
		t.ChatID = update.Message.Chat.ID
		t.ChatIDMutex.Unlock()
		break
	}
}
