package tcp_test

import (
	"fmt"
	"gosocket/util"
	"net"
	"time"

	"gosocket/data"
)

var BUFFER_SIZE int = 2048

type (
	OptFunc func(*TCPServer)
)

type TCPServer struct {
	QuitChan chan struct{}
	Listener net.Listener
	TCPOpt   util.ServerOpt
}

func DefaultServer() *TCPServer {
	BUFFER_SIZE = util.YMLConfig.Server.Socket.BufferSize
	return &TCPServer{
		TCPOpt:   util.YMLConfig.Server,
		QuitChan: make(chan struct{}),
	}
}

func NewServer(funcs ...OptFunc) (*TCPServer, error) {
	server := DefaultServer()
	for _, fn := range funcs {
		fn(server)
	}

	fmt.Printf("%+v\n", server)
	fmt.Printf("buffer size: %d", BUFFER_SIZE)
	return server, nil
}

func (s *TCPServer) Start() error {
	ln, err := net.Listen("tcp", s.TCPOpt.Connection.Address)
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
	defer fmt.Printf("remote closed: %d", conn.RemoteAddr())
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
		temp := make([]byte, read)
		copy(temp, buf)
		bufchan <- temp
	}
}
