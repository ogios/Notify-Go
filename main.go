package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	_ "gosocket/notify"
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
	// test()
	util.UnmarshalConfig()
	err := util.CreateTempDir()
	defer util.RemoveTempDir()
	if err != nil {
		panic(err)
	}
	server, err := tcp_test.NewServer()
	if err != nil {
		panic(err)
	}
	fmt.Println("start listening on " + server.Address)
	server.Start()
}
