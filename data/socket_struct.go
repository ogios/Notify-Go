package data

import (
	"fmt"
	"gosocket/config"
)

type Notification struct {
	Title    string
	Content  string
	IconPath string
}

type NotificationRaw struct {
	data       chan []byte
	step       int
	left       int
	size       int
	currentBuf []byte
}

func ParseSocketData(SocketData chan []byte) {
	itemRaw := NotificationRaw{
		data:       SocketData,
		step:       0,
		size:       config.YMLConfig.Server.Socket.BufferSize,
		left:       0,
		currentBuf: []byte{},
	}
	num := itemRaw.read(4)
	fmt.Println(num)
	// itemRaw.next()
}

func (n *NotificationRaw) read(length int32) []byte {
	total := []byte{}
	for length > 0 {
		if length > int32(n.left) {
			length = length - int32(len(n.currentBuf))
			total = append(total, n.currentBuf...)
			n.currentBuf = <-n.data
			n.left = len(n.currentBuf)
		} else {
			total = append(total, n.currentBuf[n.step:n.step+int(length)]...)
			n.step = n.step + int(length)
			return total
		}
	}
	return total
}

// func (n *NotificationRaw) next() {
// 	current := 0
// 	num := make([]byte, 4)
//
// 	for {
// 		temp := n.data[current]
// 		current++
//
// 		if string([]byte{temp}) == "\n" {
// 			break
// 		} else {
// 			num = append(num, temp)
// 		}
// 		fmt.Printf("%d", num)
// 	}
// }
