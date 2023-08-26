//go:build linux
// +build linux

package fac

import (
	"gosocket/app"
	"os/exec"
)

type Linux struct{}

func (n *Linux) Notify(item app.Notification) error {
	params := []string{
		"--app-name=",
		item.AppID,
		item.Title,
		item.Content,
	}
	if len(item.IconPath) > 0 {
		params = append(
			[]string{
				"--icon=",
				item.IconPath,
			},
			params...,
		)
	}
	cmd := exec.Command("notify-send", params...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func GetSystem() app.System {
	return &Linux{}
}
