package main

import (
	"fmt"

	tcp_test "gosocket/tcp_test"
)

func main() {
	tcp_test.UnmarshalConfig()
	server, err := tcp_test.NewServer()
	if err != nil {
		panic(err)
	}
	fmt.Println("start listening on " + server.Address)
	server.Start()
}
