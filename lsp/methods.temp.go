package lsp

import "github.com/TobiasYin/go-lsp/lsp/defines"

type method struct {
	Name          string
	Args          interface{}
	Result        interface{}
	Error         interface{}
	Code          interface{}
	ProgressToken interface{}
	WithBuiltin   bool
}

type or []interface{}

var methods = []method{
	{
		Name:        "Initialize",
		Args:        defines.InitializeParams{},
		Result:      defines.InitializeResult{},
		Error:       defines.InitializeError{},
		Code:        defines.InitializeErrorUnknownProtocolVersion,
		WithBuiltin: true,
	},
	{
		Name: "Initialized",
		Args: defines.InitializeParams{},
	},
	{
		Name: "Shutdown",
	},
	{
		Name: "Exit",
	},
	{
		Name: "DidChangeConfiguration",
		Args: defines.DidChangeConfigurationParams{},
	},
	{
		Name: "DidChangeWatchedFiles",
		Args: defines.DidChangeWatchedFilesParams{},
	},
	{
		Name: "DidOpenTextDocument",
		Args: defines.DidOpenTextDocumentParams{},
	},
	{
		Name: "DidChangeTextDocument",
		Args: defines.DidChangeTextDocumentParams{},
	},
	{
		Name: "DidCloseTextDocument",
		Args: defines.DidCloseTextDocumentParams{},
	},
	{
		Name: "WillSaveTextDocument",
		Args: defines.WillSaveTextDocumentParams{},
	},
	{
		Name: "DidSaveTextDocument",
		Args: defines.DidSaveTextDocumentParams{},
	},
	{
		Name:   "CodeActionResolve",
		Args:   defines.CodeAction{},
		Result: defines.CodeAction{},
		Error:  nil,
	},
	{
		Name:   "CodeLensResolve",
		Args:   defines.CodeLens{},
		Result: defines.CodeLens{},
		Error:  nil,
	},
	{
		Name:   "DocumentOnTypeFormatting",
		Args:   defines.DocumentOnTypeFormattingParams{},
		Result: []defines.TextEdit{},
		Error:  nil,
	},
	{
		Name:   "PrepareRename",
		Args:   defines.PrepareRenameParams{},
		Result: defines.Range{},
		Error:  nil,
	},
	{
		Name:   "DocumentLinkResolve",
		Args:   defines.DocumentLink{},
		Result: defines.DocumentLink{},
		Error:  nil,
	},
	{
		Name:          "Declaration",
		Args:          defines.DeclarationParams{},
		Result:        []defines.DeclarationLink{},
		Error:         nil,
		ProgressToken: or{[]defines.Location{}, []defines.DeclarationLink{}},
	},
	{
		Name:          "Definition",
		Args:          defines.DefinitionParams{},
		Result:        []defines.DefinitionLink{},
		Error:         nil,
		ProgressToken: or{[]defines.Location{}, []defines.DefinitionLink{}},
	},
	{
		Name:          "TypeDefinition",
		Args:          defines.TypeDefinitionParams{},
		Result:        []defines.DefinitionLink{},
		Error:         nil,
		ProgressToken: or{[]defines.Location{}, []defines.DefinitionLink{}},
	},
	{
		Name:          "Implementation",
		Args:          defines.ImplementationParams{},
		Result:        []defines.DefinitionLink{},
		Error:         nil,
		ProgressToken: or{[]defines.Location{}, []defines.DefinitionLink{}},
	},
	{
		Name:          "References",
		Args:          defines.ReferenceParams{},
		Result:        []defines.Location{},
		Error:         nil,
		ProgressToken: []defines.Location{},
	},
	{
		Name:          "DocumentHighlight",
		Args:          defines.DocumentHighlightParams{},
		Result:        []defines.DocumentHighlight{},
		Error:         nil,
		ProgressToken: []defines.DocumentHighlight{},
	},
	{
		Name:          "DocumentSymbol",
		Args:          defines.DocumentSymbolParams{},
		Result:        or{[]defines.SymbolInformation{}, []defines.DocumentSymbol{}},
		Error:         nil,
		ProgressToken: or{[]defines.SymbolInformation{}, []defines.DocumentSymbol{}},
	},
	{
		Name:          "WorkspaceSymbol",
		Args:          defines.WorkspaceSymbolParams{},
		Result:        []defines.SymbolInformation{},
		Error:         nil,
		ProgressToken: []defines.SymbolInformation{},
	},
	{
		Name:          "CodeAction",
		Args:          defines.CodeActionParams{},
		Result:        or{defines.Command{}, []defines.CodeAction{}},
		Error:         nil,
		ProgressToken: or{defines.Command{}, []defines.CodeAction{}},
	},
	{
		Name:          "CodeLens",
		Args:          defines.CodeLensParams{},
		Result:        []defines.CodeLens{},
		Error:         nil,
		ProgressToken: []defines.CodeLens{},
	},
	{
		Name:          "DocumentFormatting",
		Args:          defines.DocumentFormattingParams{},
		Result:        []defines.TextEdit{},
		Error:         nil,
		ProgressToken: nil,
	},
	{
		Name:          "DocumentRangeFormatting",
		Args:          defines.DocumentRangeFormattingParams{},
		Result:        []defines.TextEdit{},
		Error:         nil,
		ProgressToken: nil,
	},
	{
		Name:          "RenameRequest",
		Args:          defines.RenameParams{},
		Result:        defines.WorkspaceEdit{},
		Error:         nil,
		ProgressToken: nil,
	},
	{
		Name:          "DocumentLinks",
		Args:          defines.DocumentLinkParams{},
		Result:        []defines.DocumentLink{},
		Error:         nil,
		ProgressToken: []defines.DocumentLink{},
	},
	{
		Name:          "DocumentColor",
		Args:          defines.DocumentColorParams{},
		Result:        []defines.ColorInformation{},
		Error:         nil,
		ProgressToken: []defines.ColorInformation{},
	},
	{
		Name:          "ColorPresentation",
		Args:          defines.ColorPresentationParams{},
		Result:        []defines.ColorPresentation{},
		Error:         nil,
		ProgressToken: []defines.ColorPresentation{},
	},
	{
		Name:          "FoldingRanges",
		Args:          defines.FoldingRangeParams{},
		Result:        []defines.FoldingRange{},
		Error:         nil,
		ProgressToken: []defines.FoldingRange{},
	},
	{
		Name:          "SelectionRanges",
		Args:          defines.SelectionRangeParams{},
		Result:        []defines.SelectionRange{},
		Error:         nil,
		ProgressToken: []defines.SelectionRange{},
	},
	{
		Name:          "ExecuteCommand",
		Args:          defines.ExecuteCommandParams{},
		Result:        interface{}(nil),
		Error:         nil,
		ProgressToken: nil,
	},
}
