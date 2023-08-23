package tcp_test

import (
	"fmt"
	"gosocket/config"
	"gosocket/data"
	"net"
	"time"

	"github.com/jinzhu/copier"
)

var BUFFER_SIZE int = 2048

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
	server_config, err := config.GetConfig("server.connection")
	BUFFER_SIZE = config.YMLConfig.Server.Socket.BufferSize
	if err != nil {
		panic(err)
	}
	tcpopt := TCPOpt{}
	copier.Copy(&tcpopt, server_config)
	return &TCPServer{
		TCPOpt:   tcpopt,
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
	fmt.Printf("buffer size: %d", BUFFER_SIZE)
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
	conn.SetDeadline(time.Now().Add(time.Second * 10))
	bufchan := make(chan []byte)
	buf := make([]byte, BUFFER_SIZE)
	go data.ParseSocketData(bufchan)
	for {
		read, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(buf)
		bufchan <- buf[:read]
	}
}
