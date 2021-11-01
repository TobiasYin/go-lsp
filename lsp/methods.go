package lsp

import (
	"context"
	"fmt"

	"github.com/TobiasYin/go-lsp/jsonrpc"
	"github.com/TobiasYin/go-lsp/lsp/defines"
)

// initialize
func initialize(ctx context.Context, req interface{}, resp interface{}) error {
	init := req.(*defines.InitializeParams)
	res := resp.(*defines.InitializeResult)
	fmt.Println(init)
	res.Capabilities.TextDocumentSync = defines.TextDocumentSyncKindFull
	res.Capabilities.HoverProvider = true
	return nil
}
func Initialize() jsonrpc.MethodInfo {
	return jsonrpc.MethodInfo{
		Name: "initialize",
		NewRequest: func() interface{} {
			return &defines.InitializeParams{}
		},
		NewResponse: func() interface{} {
			return &defines.InitializeResult{}
		},
		Handler: initialize,
	}
}
