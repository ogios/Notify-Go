package data

import (
	"fmt"

	"gosocket/notify"
	"gosocket/util"

	"github.com/ogios/sutils"
	"golang.org/x/exp/slog"

	. "gosocket/app"
)

type NoIn struct {
	Si   *sutils.SBodyIN
	Item *Notification
}

func (n *NoIn) i() {
	if n.Item == nil {
		n.Item = new(Notification)
	}
}

func (n *NoIn) AppID() error {
	n.i()
	length, err := n.Si.Next()
	if err != nil {
		return err
	}
	if length > 255 {
		return fmt.Errorf("package name too long: %d", length)
	}
	name, err := n.Si.GetSec()
	if err != nil {
		return err
	}
	n.Item.AppID = string(name)
	return nil
}

func (n *NoIn) Title() error {
	n.i()
	length, err := n.Si.Next()
	if err != nil {
		return err
	}
	if length > 1024*4 {
		return fmt.Errorf("title too long: %d", length)
	}
	title, err := n.Si.GetSec()
	if err != nil {
		return err
	}
	n.Item.Title = string(title)
	return nil
}

func (n *NoIn) Content() error {
	n.i()
	_, err := n.Si.Next()
	if err != nil {
		return err
	}
	content, err := n.Si.GetSec()
	if err != nil {
		return err
	}
	n.Item.Title = string(content)
	return nil
}

func (n *NoIn) Pic() error {
	f, err := util.GetTempFile(n.Item.AppID, "png")
	if err != nil {
		return err
	}
	length, err := n.Si.Next()
	if err != nil {
		return err
	}
	buf := make([]byte, 1024)
	for length > 0 {
		read, err := n.Si.Read(buf)
		if err != nil {
			return err
		}
		f.Write(buf[:read])
		length -= read
	}
	n.Item.IconPath = f.Name()
	return nil
}

func ParseSocketData(n *NoIn) error {
	slog.Debug("getting AppID")
	err := n.AppID()
	if err != nil {
		return err
	}
	slog.Debug("getting Title")
	err = n.Title()
	if err != nil {
		return err
	}
	slog.Debug("getting Content")
	err = n.Content()
	if err != nil {
		return err
	}
	slog.Debug("getting Picture")
	err = n.Pic()
	if err != nil {
		return err
	}
	return nil
}

func Notify(n *NoIn) error {
	output, err := notify.Notify(*n.Item)
	slog.Info(output, "type", "OUTPUT")
	if err != nil {
		return err
	}
	return nil
}
