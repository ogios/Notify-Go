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
			s.QuitChan <- struct{}{}
			return
		}
		fmt.Println(conn.RemoteAddr())
		conn.Close()
	}
}
