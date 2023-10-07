package tcp_test

import (
	"bufio"
	"fmt"
	"net"
	"runtime"
	"time"

	"gosocket/data"
	"gosocket/util"

	"github.com/ogios/sutils"
	"golang.org/x/exp/slog"
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

	err = s.loopAccept()
	if err != nil {
		return err
	}
	return nil
}

func (s *TCPServer) loopAccept() error {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			slog.Error(fmt.Sprintf("!!!Connection Accept Error: %s", err))
			return err
		}
		fmt.Println(conn.RemoteAddr())
		go readBuf(conn)
	}
}

func readBuf(conn net.Conn) (err error) {
	defer runtime.GC()
	defer slog.Info(fmt.Sprintf("remote closed: %d", conn.RemoteAddr()))
	defer conn.Close()
	defer func() {
		if err != nil {
			slog.Error(err.Error())
		}
		if e := recover(); e != nil {
			err = e.(error)
			slog.Error(err.Error())
		}
	}()
	conn.SetDeadline(time.Now().Add(time.Second * 10))
	n := &data.NoIn{
		Si: sutils.NewSBodyIn(bufio.NewReader(conn)),
	}
	err = data.ParseSocketData(n)
	if err == nil {
		err = data.Notify(n)
		if err == nil {
			err = n.Item.Clear()
		}
	}
	n = nil
	return err
}
