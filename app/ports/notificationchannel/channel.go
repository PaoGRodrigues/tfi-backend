package alert

type NotificationChannel interface {
	Configure(string, string) error
}
