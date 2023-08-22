package tcpt_test

import (
	"fmt"
	"net"

	"github.com/jinzhu/copier"
)

type (
	OptFunc func(*TCPServer)
	TCPOpt  struct {
		Network string `yaml:"network"`
		Address string `yaml:"address"`
	}
)

type TCPServer struct {
	QuitChan chan struct{}
	Listener net.Listener
	TCPOpt
}

func DefaultServer() (*TCPServer, error) {
	server_config, err := GetConfig("server")
	if err != nil {
		panic(err)
	}
	opt := TCPOpt{}
	copier.Copy(&opt, server_config)
	return &TCPServer{
		TCPOpt:   opt,
		QuitChan: make(chan struct{}),
	}, nil
}

func NewServer(funcs ...OptFunc) (*TCPServer, error) {
	server, err := DefaultServer()
	if err != nil {
		return nil, err
	}
	for _, fn := range funcs {
		fn(server)
	}

	fmt.Printf("%+v\n", server)
	return server, nil
}

func (s *TCPServer) Start() error {
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.Listener = ln

	go s.loopAccept()
	<-s.QuitChan
	return nil
}

func (s *TCPServer) loopAccept() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			fmt.Printf("!!!Connection Accept Error: %s", err)
			s.QuitChan <- struct{}{}
			panic(err)
		}
		fmt.Println(conn.RemoteAddr())
		go readBuf(conn)
	}
}

func readBuf(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	read, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("!!!ReadBuf Error: %s", err)
		return
	}
}
