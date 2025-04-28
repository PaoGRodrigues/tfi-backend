package alert

type AlertsSender interface {
	SendLastAlertMessages() error
}

type Notifier interface {
	SendMessage(string) error
}
