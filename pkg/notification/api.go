package notification

// Notificator is API to provide to share message
type Notificator interface {
	Notify(id string, message string) error
}
