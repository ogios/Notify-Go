package app

type Notification struct {
	AppID    string
	Title    string
	Content  string
	IconPath string
}

type System interface {
	Notify(item Notification) (string, error)
}
