package notify

type Linux struct{}

func (n *Linux) Notify(item Notification) error {
	return nil
}
