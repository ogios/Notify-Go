//go:build linux
// +build linux

package notify

type Linux struct{}

func (n *Linux) Notify(item Notification) error {
	return nil
}
