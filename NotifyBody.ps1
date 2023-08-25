[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
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
