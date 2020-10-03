package notification

type Notificator interface {
	Notify(id string, message string) error
}
