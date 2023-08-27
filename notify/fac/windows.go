//go:build windows
// +build windows

package fac

import (
	"bytes"
	"gosocket/app"
	. "gosocket/app"
	"gosocket/util"
	"os/exec"
	"syscall"
	"text/template"
)

type Windows struct{}

func GetSystem() app.System {
	return &Windows{}
}

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
	path, err := util.WriteTempFile(script, "powershell", "ps1")
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
		initTemplate()
	}
	bytes := bytes.Buffer{}
	bytes.Write([]byte{0xEF, 0xBB, 0xBF})
	err := ToastTemplate.Execute(&bytes, item)
	if err != nil {
		return nil, err
	}
	return bytes.Bytes(), nil
}

func initTemplate() {
	ToastTemplate = template.New("toast")
	str := `[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.UI.Notifications.ToastNotification, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
[Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime] | Out-Null

$APP_ID = '{{if .AppID}}{{.AppID}}{{else}}Windows App{{end}}'

$template = @"
<toast activationType="protocol" duration="short">
    <visual>
        <binding template="ToastGeneric">
            {{if .IconPath}}
            <image placement="appLogoOverride" src="{{.IconPath}}" />
            {{end}}
            {{if .Title}}
            <text><![CDATA[{{.Title}}]]></text>
            {{end}}
            {{if .Content}}
            <text><![CDATA[{{.Content}}]]></text>
            {{end}}
        </binding>
    </visual>
	<audio silent="true" />
</toast>
"@

$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
$xml.LoadXml($template)
$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier($APP_ID).Show($toast)
`
	ToastTemplate.Parse(str)
}
