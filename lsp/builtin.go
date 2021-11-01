package lsp

import (
	"context"

	"github.com/TobiasYin/go-lsp/lsp/defines"
)

func (m *Methods) builtinInitialize(ctx context.Context, req *defines.InitializeParams) (defines.InitializeResult, error) {
	resp := defines.InitializeResult{}
	resp.Capabilities.TextDocumentSync = defines.TextDocumentSyncKindFull
	if m.onCompletion != nil {
		// TODO
	}
	if m.onHover != nil {
		resp.Capabilities.HoverProvider = true
	}
	if m.onSignatureHelp != nil {
		// TODO
	}
	if m.onDeclaration != nil {
		resp.Capabilities.DeclarationProvider = true
	}
	if m.onDefinition != nil {
		resp.Capabilities.DefinitionProvider = true
	}
	if m.onTypeDefinition != nil {
		resp.Capabilities.TypeDefinitionProvider = true
	}
	if m.onImplementation != nil {
		resp.Capabilities.ImplementationProvider = true
	}
	if m.onReferences != nil {
		resp.Capabilities.ReferencesProvider = true
	}
	if m.onDocumentHighlight != nil {
		resp.Capabilities.DocumentHighlightProvider = true
	}
	if m.onDocumentSymbolWithSliceDocumentSymbol != nil {
		resp.Capabilities.DocumentSymbolProvider = true
	}
	if m.onDocumentSymbolWithSliceSymbolInformation != nil {
		resp.Capabilities.DocumentSymbolProvider = true
	}
	if m.onCodeActionWithSliceCodeAction != nil {
		resp.Capabilities.CodeActionProvider = true
	}
	if m.onCodeActionWithSliceCommand != nil {
		resp.Capabilities.CodeActionProvider = true
	}
	if m.onCodeLens != nil {
		//resp.Capabilities.CodeLensProvider = true
		// TODO
	}
	if m.onDocumentLinks != nil {
		//resp.Capabilities.DocumentLinkProvider = true
		// TODO
	}
	if m.onDocumentColor != nil {
		resp.Capabilities.ColorProvider = true
	}
	if m.onWorkspaceSymbol != nil {
		resp.Capabilities.WorkspaceSymbolProvider = true
	}
	if m.onDocumentFormatting != nil {
		resp.Capabilities.DocumentFormattingProvider = true
	}
	if m.onDocumentRangeFormatting != nil {
		resp.Capabilities.DocumentRangeFormattingProvider = true
	}
	if m.onDocumentOnTypeFormatting != nil {
		//resp.Capabilities.DocumentOnTypeFormattingProvider = true
		// TODO
	}
	if m.onPrepareRename != nil {
		resp.Capabilities.RenameProvider = true
	}
	if m.onFoldingRanges != nil {
		resp.Capabilities.FoldingRangeProvider = true
	}
	if m.onSelectionRanges != nil {
		resp.Capabilities.SelectionRangeProvider = true
	}
	if m.onExecuteCommand != nil {
		//resp.Capabilities.ExecuteCommandProvider = true
		// TODO
	}
	//if m.onDocumentLinks != nil{
	//	resp.Capabilities.LinkedEditingRangeProvider = true
	//}
	//if m.onToke != nil{
	//	resp.Capabilities.SemanticTokensProvider = true
	//}
	//if m.onMon != nil{
	//	resp.Capabilities.MonikerProvider = true
	//}
	//if m.onTypeHierarchy != nil{
	//	resp.Capabilities.TypeHierarchyProvider = true
	//}

	return resp, nil
}
