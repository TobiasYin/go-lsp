package lsp

import (
	"fmt"
	"net"
	"reflect"

	"github.com/TobiasYin/go-lsp/jsonrpc"
)

type Server struct {
	Methods
	rpcServer *jsonrpc.Server
}

func NewServer(opt Options) *Server {
	s := &Server{}
	s.Opt = opt
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
	addr := "127.0.0.1:7999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		s.rpcServer.ConnComeIn(conn)
	}
}

func wrapErrorToRespError(err interface{}, code int) error{
	if isNil(err){
		return nil
	}
	if e, ok := err.(error); ok{
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