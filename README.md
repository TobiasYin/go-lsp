# LSP(language server protocol) defines for golang

[lsp](https://microsoft.github.io/language-server-protocol) types is from [vscode-languageserver-node](https://github.com/microsoft/vscode-languageserver-node).

Project is working in progress.
## Example

```go
func main() {
	server := lsp.NewServer(&lsp.Options{CompletionProvider: &defines.CompletionOptions{
		TriggerCharacters: &[]string{"."},
	}})
	server.OnHover(func(ctx context.Context, req *defines.HoverParams) (result *defines.Hover, err error) {
		logs.Println(req)
		return &defines.Hover{Contents: defines.MarkupContent{Kind: defines.MarkupKindPlainText, Value: "hello world"}}, nil
	})

	server.OnCompletion(func(ctx context.Context, req *defines.CompletionParams) (result *[]defines.CompletionItem, err error) {
		logs.Println(req)
		d := defines.CompletionItemKindText
		return &[]defines.CompletionItem{defines.CompletionItem{
			Label:               "code",
			Kind:                &d,
			InsertText:          strPtr("Hello"),
		}}, nil
	})

	server.Run()
}
```
