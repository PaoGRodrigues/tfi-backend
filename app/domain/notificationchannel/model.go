package notificationchannel

type NotificationChannel interface {
	Configure(string, string) error
}
