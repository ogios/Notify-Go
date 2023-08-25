package notify

import (
	"fmt"
	"runtime"
)

type Notification struct {
	AppID    string
	Title    string
	Content  string
	IconPath string
}

// type System
var System interface {
	Notify(item Notification) error
}

func init() {
	sys := runtime.GOOS
	fmt.Println("System: ", sys)
	switch sys {
	case "windows":
		System = &Windows{}
	case "linux":
		System = &Linux{}
	}
}
