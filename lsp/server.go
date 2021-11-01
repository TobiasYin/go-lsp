package lsp

import "github.com/TobiasYin/go-lsp/jsonrpc"

func NewServer() *jsonrpc.Server {
	s := jsonrpc.NewServer()
	s.RegisterMethod(Initialize())
	return s
}
