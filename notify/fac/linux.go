//go:build linux
// +build linux

package fac

import "gosocket/app"

type Linux struct{}

func (n *Linux) Notify(item app.Notification) error {
	return nil
}

func GetSystem() app.System {
	return &Linux{}
}
