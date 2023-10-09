package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"gosocket/notify"
	"gosocket/tcp_test"
	"gosocket/util"
)

func test() {
	c := int32(4)
	// buf := []byte{}
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, c)

	fmt.Println(buf.Bytes())
}

func main() {
	SetupLog()
	s, e := notify.NotifyRaw("notify started")
	fmt.Println(s)
	if e != nil {
		panic(e)
	}

	// test()
	err := util.CreateTempDir()
	defer util.RemoveTempDir()
	if err != nil {
		panic(err)
	}
	server, err := tcp_test.NewServer()
	if err != nil {
		panic(err)
	}
	fmt.Println("start listening on " + server.TCPOpt.Connection.Address)
	server.Start()
}
