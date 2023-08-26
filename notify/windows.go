//go:build windows
// +build windows

package notify

import (
	"bytes"
	"gosocket/util"
	"os"
	"os/exec"
	"syscall"
	"text/template"
)

type Windows struct{}

var ToastTemplate *template.Template

func (n *Windows) Notify(item Notification) error {
	script, err := buildTemplate(item)
	if err != nil {
		return err
	}
	if err := sendNotification(script); err != nil {
		return err
	}
	return nil
}

func sendNotification(script []byte) error {
	path, err := util.WriteTempFile(script, "ps1")
	if err != nil {
		return err
	}
	cmd := exec.Command("PowerShell", "-ExecutionPolicy", "Bypass", "-File", path)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func buildTemplate(item Notification) ([]byte, error) {
	if ToastTemplate == nil {
		err := initTemplate()
		if err != nil {
			return nil, err
		}
	}
	bytes := bytes.Buffer{}
	bytes.Write([]byte{0xEF, 0xBB, 0xBF})
	err := ToastTemplate.Execute(&bytes, item)
	if err != nil {
		return nil, err
	}
	return bytes.Bytes(), nil
}

func initTemplate() error {
	ToastTemplate = template.New("toast")
	path, _ := os.Getwd()
	bytes, err := os.ReadFile(path + "/NotifyBody.ps1")
	if err != nil {
		return err
	}
	ToastTemplate.Parse(string(bytes))
	return nil
}
