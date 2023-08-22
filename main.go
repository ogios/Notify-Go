package main

import (
	"fmt"

	tcpt_test "gosocket/tcp_test"
)

func main() {
	tcpt_test.UnmarshalConfig()
	server_config, err := tcpt_test.GetConfig("server.address")
	if err != nil {
		panic(err)
	}
	fmt.Println(server_config)
	// fmt.Println("start")
	// server, err := tcpt_test.NewServer()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("start listening on " + server.Address)
	// server.Start()
}
