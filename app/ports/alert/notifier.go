package alert

type Notifier interface {
	SendMessage(string) error
}
