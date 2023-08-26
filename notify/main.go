package notify

import (
	"fmt"
	"gosocket/notify/fac"
	"runtime"

	. "gosocket/app"
)

var system interface {
	Notify(item Notification) error
}

func init() {
	sys := runtime.GOOS
	fmt.Println("System: ", sys)
	system = fac.GetSystem()
}

func Notify(item Notification) error {
	return system.Notify(item)
}
