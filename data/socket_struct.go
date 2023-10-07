package data

import (
	"fmt"
	"gosocket/notify"
	"gosocket/util"
	"math"

	. "gosocket/app"
)

type NotificationRaw struct {
	data       chan []byte
	step       int
	left       int
	size       int
	currentBuf []byte
}

func byteToInt32(bytes []byte) int32 {
	var length int32 = 0
	for ind, byte := range bytes {
		length += (int32(byte) * int32(math.Pow(255, float64(len(bytes)-1-ind))))
	}
	return length
}

func ParseSocketData(SocketData chan []byte) {
	item := Notification{}
	itemRaw := NotificationRaw{
		data:       SocketData,
		step:       0,
		size:       util.YMLConfig.Server.Socket.BufferSize,
		left:       0,
		currentBuf: []byte{},
	}

	// 包名
	err := itemRaw.next(func(bytes []byte) error {
		item.AppID = string(bytes)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// 标题
	err = itemRaw.next(func(bytes []byte) error {
		item.Title = string(bytes)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// 内容
	err = itemRaw.next(func(bytes []byte) error {
		item.Content = string(bytes)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	// 图标
	_ = itemRaw.next(func(bytes []byte) error {
		path, WriteFileErr := util.WriteTempFile(bytes, item.AppID, "png")
		if WriteFileErr != nil {
			return WriteFileErr
		}
		item.IconPath = path
		return nil
	})
	output, err := notify.Notify(item)
	fmt.Println("output:", output)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("done")
}

func (n *NotificationRaw) next(fun func(bytes []byte) error) error {
	var num []byte
	var length int32
	num = n.read(4)
	length = byteToInt32(num)
	if length > 0 {
		fmt.Println(num)
		fmt.Printf("\n字段长度: %d\n", length)
		// 分隔
		num = n.read(1)
		fmt.Println("字段分隔: ", num)
		// 内容
		num = n.read(length)
		// 分隔
		fmt.Println("下一字段分隔: ", n.read(2))

		err := fun(num)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("%s", "no buf to read")
}

func (n *NotificationRaw) read(length int32) []byte {
	total := []byte{}
	for length > 0 {
		if length > int32(n.left) {
			length = length - int32(len(n.currentBuf[n.step:]))
			total = append(total, n.currentBuf[n.step:]...)
			n.currentBuf = <-n.data
			n.left = len(n.currentBuf)
			n.step = 0
		} else {
			total = append(total, n.currentBuf[n.step:n.step+int(length)]...)
			n.step = n.step + int(length)
			return total
		}
	}
	return total
}
