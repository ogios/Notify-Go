package app

import (
	"os"
)

type Notification struct {
	AppID    string
	Title    string
	Content  string
	IconPath string
}

func (n *Notification) Clear() error {
	if err := os.Remove(n.IconPath); err != nil {
		return err
	}
	return nil
}

type System interface {
	Notify(item Notification) (string, error)
}
