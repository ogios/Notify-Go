//go:build linux
// +build linux

package fac

import (
	"fmt"
	"os/exec"

	"gosocket/app"
)

type Linux struct{}

func (n *Linux) Notify(item app.Notification) (string, error) {
	params := []string{
		"--app-name=" + item.AppID,
		item.Title,
		item.Content,
	}
	if len(item.IconPath) > 0 {
		params = append(
			[]string{
				"--icon=" + item.IconPath,
			},
			params...,
		)
	}
	cmd := exec.Command("notify-send", params...)
	fmt.Println(cmd.Args)
	output, err := cmd.Output()
	return string(output), err
}

func GetSystem() app.System {
	return &Linux{}
}
