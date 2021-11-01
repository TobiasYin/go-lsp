package lsp

import (
	"context"

	"github.com/TobiasYin/go-lsp/lsp/defines"
)

func(m *Methods) builtinInitialize(ctx context.Context, req *defines.InitializeParams) (defines.InitializeResult, error){
	mtds := m.GetMethods()
	for _, m := range mtds {
		if m != nil {
			// TODO
		}
	}
	return defines.InitializeResult{}, nil
}