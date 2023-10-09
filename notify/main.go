package notify

import (
	"fmt"
	"gosocket/notify/fac"
	"runtime"

	. "gosocket/app"
)

var system System

func init() {
	sys := runtime.GOOS
	fmt.Println("System: ", sys)
	system = fac.GetSystem()
}

func Notify(item Notification) (string, error) {
	return system.Notify(item)
}

func NotifyRaw(s string) (string, error) {
	return system.Notify(Notification{
		AppID:    "com.raw.notify",
		Title:    "Raw message",
		Content:  s,
		IconPath: "",
	})

}
