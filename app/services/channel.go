package services

type NotificationChannel interface {
	Configure(string, string) error
	SendMessage(string) error
}
