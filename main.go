package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"gosocket/config"
	"gosocket/tcp_test"
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
	config.UnmarshalConfig()
	server, err := tcp_test.NewServer()
	if err != nil {
		panic(err)
	}
	fmt.Println("start listening on " + server.Address)
	server.Start()
}
