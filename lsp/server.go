package lsp

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"

	"github.com/TobiasYin/go-lsp/jsonrpc"
	"github.com/TobiasYin/go-lsp/logs"
)

type Server struct {
	Methods
	rpcServer *jsonrpc.Server
}

func NewServer(opt *Options) *Server {
	s := &Server{}
	s.Opt = *opt
	s.rpcServer = jsonrpc.NewServer()
	return s
}

func (s *Server) Run() {
	mtds := s.GetMethods()
	for _, m := range mtds {
		if m != nil {
			s.rpcServer.RegisterMethod(*m)
		}
	}

	s.run()
}
func (s *Server) run() {
	addr := s.Opt.Address
	netType := s.Opt.Network
	if netType != "" {
		if addr == "" {
			addr = "127.0.0.1:7998"
		}
		logs.Printf("use socket mode: net: %s, addr: %s\n", netType, addr)
		listener, err := net.Listen(netType, addr)
		if err != nil {
			panic(err)
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			go s.rpcServer.ConnComeIn(conn)
		}
	} else {
		logs.Println("use stdio mode.")
		// use stdio mode
		s.rpcServer.ConnComeIn(NewStdio())
	}
}

func wrapErrorToRespError(err interface{}, code int) error {
	if isNil(err) {
		return nil
	}
	if e, ok := err.(error); ok {
		return e
	}
	return jsonrpc.ResponseError{
		Code:    code,
		Message: fmt.Sprintf("%v", err),
		Data:    err,
	}
}
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return true
	}
	return false
}

func (s *Server) SendNotification(method string, params json.RawMessage) error {
	notification := jsonrpc.NotificationMessage{
		BaseMessage: jsonrpc.BaseMessage{
			Jsonrpc: "2.0",
		},
		Method: method,
		Params: params,
	}

	sessions := s.rpcServer.GetAllSession()

	for _, session := range sessions {
		err := session.Notify(notification)
		if err != nil {
			return err
		}
	}
	return nil
}
