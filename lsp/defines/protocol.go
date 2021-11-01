package defines

type ClientCapabilities struct {
	_ClientCapabilities
	WorkspaceFoldersClientCapabilities
	ConfigurationClientCapabilities
	WorkDoneProgressClientCapabilities
}
type ServerCapabilities struct {
	_ServerCapabilities
	WorkspaceFoldersServerCapabilities
}
type InitializeParams struct {
	_InitializeParams
	WorkspaceFoldersInitializeParams
}

type _ServerCapabilities struct {

	// Defines how text documents are synced. Is either a detailed structure defining each notification or
	// for backwards compatibility the TextDocumentSyncKind number.
	TextDocumentSync *interface{} // TextDocumentSyncOptions, TextDocumentSyncKind,

	// The server provides completion support.
	CompletionProvider *CompletionOptions

	// The server provides hover support.
	HoverProvider *interface{} // bool, HoverOptions,

	// The server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions

	// The server provides Goto Declaration support.
	DeclarationProvider *interface{} // bool, DeclarationOptions, DeclarationRegistrationOptions,

	// The server provides goto definition support.
	DefinitionProvider *interface{} // bool, DefinitionOptions,

	// The server provides Goto Type Definition support.
	TypeDefinitionProvider *interface{} // bool, TypeDefinitionOptions, TypeDefinitionRegistrationOptions,

	// The server provides Goto Implementation support.
	ImplementationProvider *interface{} // bool, ImplementationOptions, ImplementationRegistrationOptions,

	// The server provides find references support.
	ReferencesProvider *interface{} // bool, ReferenceOptions,

	// The server provides document highlight support.
	DocumentHighlightProvider *interface{} // bool, DocumentHighlightOptions,

	// The server provides document symbol support.
	DocumentSymbolProvider *interface{} // bool, DocumentSymbolOptions,

	// The server provides code actions. CodeActionOptions may only be
	// specified if the client states that it supports
	// `codeActionLiteralSupport` in its initial `initialize` request.
	CodeActionProvider *interface{} // bool, CodeActionOptions,

	// The server provides code lens.
	CodeLensProvider *CodeLensOptions

	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions

	// The server provides color provider support.
	ColorProvider *interface{} // bool, DocumentColorOptions, DocumentColorRegistrationOptions,

	// The server provides workspace symbol support.
	WorkspaceSymbolProvider *interface{} // bool, WorkspaceSymbolOptions,

	// The server provides document formatting.
	DocumentFormattingProvider *interface{} // bool, DocumentFormattingOptions,

	// The server provides document range formatting.
	DocumentRangeFormattingProvider *interface{} // bool, DocumentRangeFormattingOptions,

	// The server provides document formatting on typing.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions

	// The server provides rename support. RenameOptions may only be
	// specified if the client states that it supports
	// `prepareSupport` in its initial `initialize` request.
	RenameProvider *interface{} // bool, RenameOptions,

	// The server provides folding provider support.
	FoldingRangeProvider *interface{} // bool, FoldingRangeOptions, FoldingRangeRegistrationOptions,

	// The server provides selection range support.
	SelectionRangeProvider *interface{} // bool, SelectionRangeOptions, SelectionRangeRegistrationOptions,

	// The server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions

	// The server provides call hierarchy support.
	//
	// @since 3.16.0
	CallHierarchyProvider *interface{} // bool, CallHierarchyOptions, CallHierarchyRegistrationOptions,

	// The server provides linked editing range support.
	//
	// @since 3.16.0
	LinkedEditingRangeProvider *interface{} // bool, LinkedEditingRangeOptions, LinkedEditingRangeRegistrationOptions,

	// The server provides semantic tokens support.
	//
	// @since 3.16.0
	SemanticTokensProvider *interface{} // SemanticTokensOptions, SemanticTokensRegistrationOptions,

	// Window specific server capabilities.
	Workspace *struct {

		// The server is interested in notificationsrequests for operations on files.
		//
		// @since 3.16.0
		FileOperations *FileOperationOptions
	}

	// The server provides moniker support.
	//
	// @since 3.16.0
	MonikerProvider *interface{} // bool, MonikerOptions, MonikerRegistrationOptions,

	// Experimental server capabilities.
	Experimental *interface{}
}

/**
 * @deprecated Use ApplyWorkspaceEditResult instead.
 */
type ApplyWorkspaceEditResponse ApplyWorkspaceEditResult

/**
 * General parameters to to register for an notification or to register a provider.
 */
type Registration struct {

	// The id used to register the request. The id can be used to deregister
	// the request again.
	Id string

	// The method to register for.
	Method string

	// Options necessary for the registration.
	RegisterOptions *interface{}
}

type RegistrationParams struct {
	Registrations []Registration
}

/**
 * General parameters to unregister a request or notification.
 */
type Unregistration struct {

	// The id used to unregister the request or notification. Usually an id
	// provided during the register request.
	Id string

	// The method to unregister for.
	Method string
}

type ProgressToken interface{} // number | string

type WorkDoneProgressParams struct {

	// An optional token that a server can use to report work done progress.
	WorkDoneToken *ProgressToken
}

type PartialResultParams struct {

	// An optional token that a server can use to report partial results (e.g. streaming) to
	// the client.
	PartialResultToken *ProgressToken
}

/**
 * A parameter literal used in requests to pass a text document and a position inside that
 * document.
 */
type TextDocumentPositionParams struct {

	// The text document.
	TextDocument TextDocumentIdentifier

	// The position inside the text document.
	Position Position
}

/**
 * Workspace specific client capabilities.
 */
type WorkspaceClientCapabilities struct {

	// The client supports applying batch edits
	// to the workspace by supporting the request
	// 'workspaceapplyEdit'
	ApplyEdit *bool

	// Capabilities specific to `WorkspaceEdit`s
	WorkspaceEdit *WorkspaceEditClientCapabilities

	// Capabilities specific to the `workspacedidChangeConfiguration` notification.
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities

	// Capabilities specific to the `workspacedidChangeWatchedFiles` notification.
	DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities

	// Capabilities specific to the `workspacesymbol` request.
	Symbol *WorkspaceSymbolClientCapabilities

	// Capabilities specific to the `workspaceexecuteCommand` request.
	ExecuteCommand *ExecuteCommandClientCapabilities

	// Capabilities specific to the semantic token requests scoped to the
	// workspace.
	//
	// @since 3.16.0.
	SemanticTokens *SemanticTokensWorkspaceClientCapabilities

	// Capabilities specific to the code lens requests scoped to the
	// workspace.
	//
	// @since 3.16.0.
	CodeLens *CodeLensWorkspaceClientCapabilities

	// The client has support for file notificationsrequests for user operations on files.
	//
	// Since 3.16.0
	FileOperations *FileOperationClientCapabilities

	// Capabilities specific to the inline values requests scoped to the
	// workspace.
	//
	// @since 3.17.0.
	InlineValues *InlineValuesWorkspaceClientCapabilities
}

/**
 * Completion client capabilities
 */
type CompletionClientCapabilities struct {

	// Whether completion supports dynamic registration.
	DynamicRegistration *bool

	// The client supports the following `CompletionItem` specific
	// capabilities.
	CompletionItem *interface{} //i, n, t, e, r, f, a, c, e, {, },  ,  , /, /,

	CompletionItemKind *struct {

		// The completion item kind values the client supports. When this
		// property exists the client also guarantees that it will
		// handle values outside its set gracefully and falls back
		// to a default value when unknown.
		//
		// If this property is not present the client only supports
		// the completion items kinds from `Text` to `Reference` as defined in
		// the initial version of the protocol.
		ValueSet *[]CompletionItemKind
	}

	// Defines how the client handles whitespace and indentation
	// when accepting a completion item that uses multi line
	// text in either `insertText` or `textEdit`.
	//
	// @since 3.17.0
	InsertTextMode *InsertTextMode

	// The client supports to send additional context information for a
	// `textDocumentcompletion` request.
	ContextSupport *bool
}

/**
 * Text document specific client capabilities.
 */
type TextDocumentClientCapabilities struct {

	// Defines which synchronization capabilities the client supports.
	Synchronization *TextDocumentSyncClientCapabilities

	// Capabilities specific to the `textDocumentcompletion`
	Completion *CompletionClientCapabilities

	// Capabilities specific to the `textDocumenthover`
	Hover *HoverClientCapabilities

	// Capabilities specific to the `textDocumentsignatureHelp`
	SignatureHelp *SignatureHelpClientCapabilities

	// Capabilities specific to the `textDocumentdeclaration`
	//
	// @since 3.14.0
	Declaration *DeclarationClientCapabilities

	// Capabilities specific to the `textDocumentdefinition`
	Definition *DefinitionClientCapabilities

	// Capabilities specific to the `textDocumenttypeDefinition`
	//
	// @since 3.6.0
	TypeDefinition *TypeDefinitionClientCapabilities

	// Capabilities specific to the `textDocumentimplementation`
	//
	// @since 3.6.0
	Implementation *ImplementationClientCapabilities

	// Capabilities specific to the `textDocumentreferences`
	References *ReferenceClientCapabilities

	// Capabilities specific to the `textDocumentdocumentHighlight`
	DocumentHighlight *DocumentHighlightClientCapabilities

	// Capabilities specific to the `textDocumentdocumentSymbol`
	DocumentSymbol *DocumentSymbolClientCapabilities

	// Capabilities specific to the `textDocumentcodeAction`
	CodeAction *CodeActionClientCapabilities

	// Capabilities specific to the `textDocumentcodeLens`
	CodeLens *CodeLensClientCapabilities

	// Capabilities specific to the `textDocumentdocumentLink`
	DocumentLink *DocumentLinkClientCapabilities

	// Capabilities specific to the `textDocumentdocumentColor`
	ColorProvider *DocumentColorClientCapabilities

	// Capabilities specific to the `textDocumentformatting`
	Formatting *DocumentFormattingClientCapabilities

	// Capabilities specific to the `textDocumentrangeFormatting`
	RangeFormatting *DocumentRangeFormattingClientCapabilities

	// Capabilities specific to the `textDocumentonTypeFormatting`
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities

	// Capabilities specific to the `textDocumentrename`
	Rename *RenameClientCapabilities

	// Capabilities specific to `textDocumentfoldingRange` request.
	//
	// @since 3.10.0
	FoldingRange *FoldingRangeClientCapabilities

	// Capabilities specific to `textDocumentselectionRange` request.
	//
	// @since 3.15.0
	SelectionRange *SelectionRangeClientCapabilities

	// Capabilities specific to `textDocumentpublishDiagnostics` notification.
	PublishDiagnostics *PublishDiagnosticsClientCapabilities

	// Capabilities specific to the various call hierarchy request.
	//
	// @since 3.16.0
	CallHierarchy *CallHierarchyClientCapabilities

	// Capabilities specific to the various semantic token request.
	//
	// @since 3.16.0
	SemanticTokens *SemanticTokensClientCapabilities

	// Capabilities specific to the linked editing range request.
	//
	// @since 3.16.0
	LinkedEditingRange *LinkedEditingRangeClientCapabilities

	// Client capabilities specific to the moniker request.
	//
	// @since 3.16.0
	Moniker *MonikerClientCapabilities

	// Capabilities specific to the various type hierarchy requests.
	//
	// @since 3.17.0 - proposed state
	TypeHierarchy *TypeHierarchyClientCapabilities

	// Capabilities specific to the `textDocumentinlineValues` request.
	//
	// @since 3.17.0 - proposed state
	InlineValues *InlineValuesClientCapabilities
}

type WindowClientCapabilities struct {

	// Whether client supports handling progress notifications. If set
	// servers are allowed to report in `workDoneProgress` property in the
	// request specific server capabilities.
	//
	// @since 3.15.0
	WorkDoneProgress *bool

	// Capabilities specific to the showMessage request.
	//
	// @since 3.16.0
	ShowMessage *ShowMessageRequestClientCapabilities

	// Capabilities specific to the showDocument request.
	//
	// @since 3.16.0
	ShowDocument *ShowDocumentClientCapabilities
}

/**
 * Client capabilities specific to regular expressions.
 *
 * @since 3.16.0
 */
type RegularExpressionsClientCapabilities struct {

	// The engine's name.
	Engine string

	// The engine's version.
	Version *string
}

/**
 * Client capabilities specific to the used markdown parser.
 *
 * @since 3.16.0
 */
type MarkdownClientCapabilities struct {

	// The name of the parser.
	Parser string

	// The version of the parser.
	Version *string
}

/**
 * General client capabilities.
 *
 * @since 3.16.0
 */
type GeneralClientCapabilities struct {

	// Client capability that signals how the client
	// handles stale requests (e.g. a request
	// for which the client will not process the response
	// anymore since the information is outdated).
	//
	// @since 3.17.0
	StaleRequestSupport *interface{} // cancel, retryOnContentModified,

	// Client capabilities specific to regular expressions.
	//
	// @since 3.16.0
	RegularExpressions *RegularExpressionsClientCapabilities

	// Client capabilities specific to the client's markdown parser.
	//
	// @since 3.16.0
	Markdown *MarkdownClientCapabilities
}

/**
 * Defines the capabilities provided by the client.
 */
type _ClientCapabilities struct {

	// Workspace specific client capabilities.
	Workspace *WorkspaceClientCapabilities

	// Text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities

	// Window specific client capabilities.
	Window *WindowClientCapabilities

	// General client capabilities.
	//
	// @since 3.16.0
	General *GeneralClientCapabilities

	// Experimental client capabilities.
	Experimental *interface{}
}

/**
 * Static registration options to be returned in the initialize
 * request.
 */
type StaticRegistrationOptions struct {

	// The id used to register the request. The id can be used to deregister
	// the request again. See also Registration#id.
	Id *string
}

/**
 * General text document registration options.
 */
type TextDocumentRegistrationOptions struct {

	// A document selector to identify the scope of the registration. If set to null
	// the document selector provided on the client side will be used.
	DocumentSelector interface{} // DocumentSelector, null,
}

/**
 * Save options.
 */
type SaveOptions struct {

	// The client is supposed to include the content on save.
	IncludeText *bool
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress *bool
}

/**
 * The result returned from an initialize request.
 */
type InitializeResult struct {

	// The capabilities the language server provides.
	Capabilities ServerCapabilities

	// Information about the server.
	//
	// @since 3.15.0
	ServerInfo *struct {
		Name    string
		Version *string
	} // name, version,

	// Custom initialization results.
	Custom interface{}
}

/**
 * The initialize parameters
 */
type _InitializeParams struct {
	WorkDoneProgressParams

	// The process Id of the parent process that started
	// the server.
	ProcessId interface{} // int, null,

	// Information about the client
	//
	// @since 3.15.0
	ClientInfo *interface{} // name, version,

	// The locale the client is currently showing the user interface
	// in. This must not necessarily be the locale of the operating
	// system.
	//
	// Uses IETF language tags as the value's syntax
	// (See https:en.wikipedia.orgwikiIETF_language_tag)
	//
	// @since 3.16.0
	Locale *string

	// The rootPath of the workspace. Is null
	// if no folder is open.
	//
	// @deprecated in favour of rootUri.
	RootPath *interface{} // string, null,

	// The rootUri of the workspace. Is null if no
	// folder is open. If both `rootPath` and `rootUri` are set
	// `rootUri` wins.
	//
	// @deprecated in favour of workspaceFolders.
	RootUri interface{} // DocumentUri, null,

	// The capabilities provided by the client (editor or tool)
	Capabilities ClientCapabilities

	// User provided initialization options.
	InitializationOptions *interface{}

	// The initial trace setting. If omitted trace is disabled ('off').
	Trace *interface{} // interface{} // 'off', interface{} // 'messages', interface{} // 'compact', interface{} // 'verbose',
}

/**
 * The data type of the ResponseError if the
 * initialize request fails.
 */
type InitializeError struct {

	// Indicates whether the client execute the following retry logic:
	// (1) show the message provided by the ResponseError to the user
	// (2) user selects retry or cancel
	// (3) if user selected retry the initialize method is sent again.
	Retry bool
}

//---- Configuration notification ----
type DidChangeConfigurationClientCapabilities struct {

	// Did change configuration notification supports dynamic registration.
	DynamicRegistration *bool
}

type DidChangeConfigurationRegistrationOptions struct {
	Section *interface{} // string, []string,
}

/**
 * The parameters of a change configuration notification.
 */
type DidChangeConfigurationParams struct {

	// The actual changed settings
	Settings interface{}
}

/**
 * The parameters of a notification message.
 */
type ShowMessageParams struct {

	// The message type. See {@link MessageType}
	Type MessageType

	// The actual message
	Message string
}

/**
 * Show message request client capabilities
 */
type ShowMessageRequestClientCapabilities struct {

	// Capabilities specific to the `MessageActionItem` type.
	MessageActionItem *struct {

		// Whether the client supports additional attributes which
		// are preserved and send back to the server in the
		// request's response.
		AdditionalPropertiesSupport *bool
	}
}

type MessageActionItem struct {

	// A short title like 'Retry', 'Open Log' etc.
	Title string

	// Additional attributes that the client preserves and
	// sends back to the server. This depends on the client
	// capability window.messageActionItem.additionalPropertiesSupport
	Key interface{} // string, bool, int, interface{},
}

type ShowMessageRequestParams struct {

	// The message type. See {@link MessageType}
	Type MessageType

	// The actual message
	Message string

	// The message action items to present.
	Actions *[]MessageActionItem
}

/**
 * The log message parameters.
 */
type LogMessageParams struct {

	// The message type. See {@link MessageType}
	Type MessageType

	// The actual message
	Message string
}

//---- Text document notifications ----
type TextDocumentSyncClientCapabilities struct {

	// Whether text document synchronization supports dynamic registration.
	DynamicRegistration *bool

	// The client supports sending will save notifications.
	WillSave *bool

	// The client supports sending a will save request and
	// waits for a response providing text edits which will
	// be applied to the document before it is saved.
	WillSaveWaitUntil *bool

	// The client supports did save notifications.
	DidSave *bool
}

type TextDocumentSyncOptions struct {

	// Open and close notifications are sent to the server. If omitted open close notification should not
	// be sent.
	OpenClose *bool

	// Change notifications are sent to the server. See TextDocumentSyncKind.None, TextDocumentSyncKind.Full
	// and TextDocumentSyncKind.Incremental. If omitted it defaults to TextDocumentSyncKind.None.
	Change *TextDocumentSyncKind

	// If present will save notifications are sent to the server. If omitted the notification should not be
	// sent.
	WillSave *bool

	// If present will save wait until requests are sent to the server. If omitted the request should not be
	// sent.
	WillSaveWaitUntil *bool

	// If present save notifications are sent to the server. If omitted the notification should not be
	// sent.
	Save *interface{} // bool, SaveOptions,
}

/**
 * The parameters send in a open text document notification
 */
type DidOpenTextDocumentParams struct {

	// The document that was opened.
	TextDocument TextDocumentItem
}

/**
 * The change text document notification's parameters.
 */
type DidChangeTextDocumentParams struct {

	// The document that did change. The version number points
	// to the version after all provided content changes have
	// been applied.
	TextDocument VersionedTextDocumentIdentifier

	// The actual content changes. The content changes describe single state changes
	// to the document. So if there are two content changes c1 (at array index 0) and
	// c2 (at array index 1) for a document in state S then c1 moves the document from
	// S to S' and c2 from S' to S''. So c1 is computed on the state S and c2 is computed
	// on the state S'.
	//
	// To mirror the content of a document using change events use the following approach:
	// - start with the same initial content
	// - apply the 'textDocumentdidChange' notifications in the order you receive them.
	// - apply the `TextDocumentContentChangeEvent`s in a single notification in the order
	// you receive them.
	ContentChanges []TextDocumentContentChangeEvent
}

/**
 * An event describing a change to a text document. If range and rangeLength are omitted
 * the new text is considered to be the full content of the document.
 */
type TextDocumentContentChangeEvent struct {

	// The range of the document that changed.
	Range Range

	// The optional length of the range that got replaced.
	//
	// @deprecated use range instead.
	RangeLength *uint

	// The new text for the provided range.
	Text interface{} // string, {"text": string}
}

/**
 * Describe options to be used when registered for text document change events.
 */
type TextDocumentChangeRegistrationOptions struct {
	TextDocumentRegistrationOptions

	// How documents are synced to the server.
	SyncKind TextDocumentSyncKind
}

/**
 * The parameters send in a close text document notification
 */
type DidCloseTextDocumentParams struct {

	// The document that was closed.
	TextDocument TextDocumentIdentifier
}

/**
 * The parameters send in a save text document notification
 */
type DidSaveTextDocumentParams struct {

	// The document that was closed.
	TextDocument TextDocumentIdentifier

	// Optional the content when saved. Depends on the includeText value
	// when the save notification was requested.
	Text *string
}

/**
 * Save registration options.
 */
type TextDocumentSaveRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SaveOptions
}

/**
 * The parameters send in a will save text document notification.
 */
type WillSaveTextDocumentParams struct {

	// The document that will be saved.
	TextDocument TextDocumentIdentifier

	// The 'TextDocumentSaveReason'.
	Reason TextDocumentSaveReason
}

//---- File eventing ----
type DidChangeWatchedFilesClientCapabilities struct {

	// Did change watched files notification supports dynamic registration. Please note
	// that the current protocol doesn't support static configuration for file changes
	// from the server side.
	DynamicRegistration *bool
}

/**
 * The watched files change notification's parameters.
 */
type DidChangeWatchedFilesParams struct {

	// The actual file events.
	Changes []FileEvent
}

/**
 * An event describing a file change.
 */
type FileEvent struct {

	// The file's uri.
	Uri DocumentUri

	// The change type.
	Type FileChangeType
}

/**
 * Describe options to be used when registered for text document change events.
 */
type DidChangeWatchedFilesRegistrationOptions struct {

	// The watchers to register.
	Watchers []FileSystemWatcher
}

type FileSystemWatcher struct {

	// The  glob pattern to watch. Glob patterns can have the following syntax:
	// - `` to match one or more characters in a path segment
	// - `?` to match on one character in a path segment
	// - `` to match any number of path segments, including none
	// - `{}` to group conditions (e.g. `​.{ts,js}` matches all TypeScript and JavaScript files)
	// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
	// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
	GlobPattern string

	// The kind of events of interest. If omitted it defaults
	// to WatchKind.Create | WatchKind.Change | WatchKind.Delete
	// which is 7.
	Kind *uint
}

/**
 * The publish diagnostic client capabilities.
 */
type PublishDiagnosticsClientCapabilities struct {

	// Whether the clients accepts diagnostics with related information.
	RelatedInformation *bool

	// Client supports the tag property to provide meta data about a diagnostic.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.15.0
	TagSupport *struct {

		// The tags supported by the client.
		ValueSet []DiagnosticTag
	}

	// Whether the client interprets the version property of the
	// `textDocumentpublishDiagnostics` notification`s parameter.
	//
	// @since 3.15.0
	VersionSupport *bool

	// Client supports a codeDescription property
	//
	// @since 3.16.0
	CodeDescriptionSupport *bool

	// Whether code action supports the `data` property which is
	// preserved between a `textDocumentpublishDiagnostics` and
	// `textDocumentcodeAction` request.
	//
	// @since 3.16.0
	DataSupport *bool
}

/**
 * The publish diagnostic notification's parameters.
 */
type PublishDiagnosticsParams struct {

	// The URI for which diagnostic information is reported.
	Uri DocumentUri

	// Optional the version number of the document the diagnostics are published for.
	//
	// @since 3.15.0
	Version *int

	// An array of diagnostic information items.
	Diagnostics []Diagnostic
}

/**
 * Contains additional information about the context in which a completion request is triggered.
 */
type CompletionContext struct {

	// How the completion was triggered.
	TriggerKind CompletionTriggerKind

	// The trigger character (a single character) that has trigger code complete.
	// Is undefined if `triggerKind !== CompletionTriggerKind.TriggerCharacter`
	TriggerCharacter *string
}

/**
 * Completion parameters
 */
type CompletionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	// The completion context. This is only available it the client specifies
	// to send this using the client capability `textDocument.completion.contextSupport === true`
	Context *CompletionContext
}

/**
 * Completion options.
 */
type CompletionOptions struct {
	WorkDoneProgressOptions

	// Most tools trigger completion request automatically without explicitly requesting
	// it using a keyboard shortcut (e.g. Ctrl+Space). Typically they do so when the user
	// starts to type an identifier. For example if the user types `c` in a JavaScript file
	// code complete will automatically pop up present `console` besides others as a
	// completion item. Characters that make up identifiers don't need to be listed here.
	//
	// If code complete should automatically be trigger on characters not being valid inside
	// an identifier (for example `.` in JavaScript) list them in `triggerCharacters`.
	TriggerCharacters *[]string

	// The list of all possible characters that commit a completion. This field can be used
	// if clients don't support individual commit characters per completion item. See
	// `ClientCapabilities.textDocument.completion.completionItem.commitCharactersSupport`
	//
	// If a server provides both `allCommitCharacters` and commit characters on an individual
	// completion item the ones on the completion item win.
	//
	// @since 3.2.0
	AllCommitCharacters *[]string

	// The server provides support to resolve additional
	// information for a completion item.
	ResolveProvider *bool

	// The server supports the following `CompletionItem` specific
	// capabilities.
	//
	// @since 3.17.0 - proposed state
	CompletionItem *struct {

		// The server has support for completion item label
		// details (see also `CompletionItemLabelDetails`) when
		// receiving a completion item in a resolve call.
		//
		// @since 3.17.0 - proposed state
		LabelDetailsSupport *bool
	}
}

/**
 * Registration options for a [CompletionRequest](#CompletionRequest).
 */
type CompletionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CompletionOptions
}

//---- Hover Support -------------------------------
type HoverClientCapabilities struct {

	// Whether hover supports dynamic registration.
	DynamicRegistration *bool

	// Client supports the follow content formats for the content
	// property. The order describes the preferred format of the client.
	ContentFormat *[]MarkupKind
}

/**
 * Hover options.
 */
type HoverOptions struct {
	WorkDoneProgressOptions
}

/**
 * Parameters for a [HoverRequest](#HoverRequest).
 */
type HoverParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

/**
 * Registration options for a [HoverRequest](#HoverRequest).
 */
type HoverRegistrationOptions struct {
	TextDocumentRegistrationOptions
	HoverOptions
}

/**
 * Client Capabilities for a [SignatureHelpRequest](#SignatureHelpRequest).
 */
type SignatureHelpClientCapabilities struct {

	// Whether signature help supports dynamic registration.
	DynamicRegistration *bool

	// The client supports the following `SignatureInformation`
	// specific properties.
	SignatureInformation *interface{} // documentationFormat, parameterInformation, activeParameterSupport,

	// The client supports to send additional context information for a
	// `textDocumentsignatureHelp` request. A client that opts into
	// contextSupport will also support the `retriggerCharacters` on
	// `SignatureHelpOptions`.
	//
	// @since 3.15.0
	ContextSupport *bool
}

/**
 * Server Capabilities for a [SignatureHelpRequest](#SignatureHelpRequest).
 */
type SignatureHelpOptions struct {
	WorkDoneProgressOptions

	// List of characters that trigger signature help.
	TriggerCharacters *[]string

	// List of characters that re-trigger signature help.
	//
	// These trigger characters are only active when signature help is already showing. All trigger characters
	// are also counted as re-trigger characters.
	//
	// @since 3.15.0
	RetriggerCharacters *[]string
}

/**
 * Additional information about the context in which a signature help request was triggered.
 *
 * @since 3.15.0
 */
type SignatureHelpContext struct {

	// Action that caused signature help to be triggered.
	TriggerKind SignatureHelpTriggerKind

	// Character that caused signature help to be triggered.
	//
	// This is undefined when `triggerKind !== SignatureHelpTriggerKind.TriggerCharacter`
	TriggerCharacter *string

	// `true` if signature help was already showing when it was triggered.
	//
	// Retrigger occurs when the signature help is already active and can be caused by actions such as
	// typing a trigger character, a cursor move, or document content changes.
	IsRetrigger bool

	// The currently active `SignatureHelp`.
	//
	// The `activeSignatureHelp` has its `SignatureHelp.activeSignature` field updated based on
	// the user navigating through available signatures.
	ActiveSignatureHelp *SignatureHelp
}

/**
 * Parameters for a [SignatureHelpRequest](#SignatureHelpRequest).
 */
type SignatureHelpParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams

	// The signature help context. This is only available if the client specifies
	// to send this using the client capability `textDocument.signatureHelp.contextSupport === true`
	//
	// @since 3.15.0
	Context *SignatureHelpContext
}

/**
 * Registration options for a [SignatureHelpRequest](#SignatureHelpRequest).
 */
type SignatureHelpRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SignatureHelpOptions
}

/**
 * Client Capabilities for a [DefinitionRequest](#DefinitionRequest).
 */
type DefinitionClientCapabilities struct {

	// Whether definition supports dynamic registration.
	DynamicRegistration *bool

	// The client supports additional metadata in the form of definition links.
	//
	// @since 3.14.0
	LinkSupport *bool
}

/**
 * Server Capabilities for a [DefinitionRequest](#DefinitionRequest).
 */
type DefinitionOptions struct {
	WorkDoneProgressOptions
}

/**
 * Parameters for a [DefinitionRequest](#DefinitionRequest).
 */
type DefinitionParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

/**
 * Registration options for a [DefinitionRequest](#DefinitionRequest).
 */
type DefinitionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DefinitionOptions
}

/**
 * Client Capabilities for a [ReferencesRequest](#ReferencesRequest).
 */
type ReferenceClientCapabilities struct {

	// Whether references supports dynamic registration.
	DynamicRegistration *bool
}

/**
 * Parameters for a [ReferencesRequest](#ReferencesRequest).
 */
type ReferenceParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams

	Context ReferenceContext
}

/**
 * Reference options.
 */
type ReferenceOptions struct {
	WorkDoneProgressOptions
}

/**
 * Registration options for a [ReferencesRequest](#ReferencesRequest).
 */
type ReferenceRegistrationOptions struct {
	TextDocumentRegistrationOptions
	ReferenceOptions
}

/**
 * Client Capabilities for a [DocumentHighlightRequest](#DocumentHighlightRequest).
 */
type DocumentHighlightClientCapabilities struct {

	// Whether document highlight supports dynamic registration.
	DynamicRegistration *bool
}

/**
 * Parameters for a [DocumentHighlightRequest](#DocumentHighlightRequest).
 */
type DocumentHighlightParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
	PartialResultParams
}

/**
 * Provider options for a [DocumentHighlightRequest](#DocumentHighlightRequest).
 */
type DocumentHighlightOptions struct {
	WorkDoneProgressOptions
}

/**
 * Registration options for a [DocumentHighlightRequest](#DocumentHighlightRequest).
 */
type DocumentHighlightRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentHighlightOptions
}

/**
 * Client Capabilities for a [DocumentSymbolRequest](#DocumentSymbolRequest).
 */
type DocumentSymbolClientCapabilities struct {

	// Whether document symbol supports dynamic registration.
	DynamicRegistration *bool

	// Specific capabilities for the `SymbolKind`.
	SymbolKind *struct {

		// The symbol kind values the client supports. When this
		// property exists the client also guarantees that it will
		// handle values outside its set gracefully and falls back
		// to a default value when unknown.
		//
		// If this property is not present the client only supports
		// the symbol kinds from `File` to `Array` as defined in
		// the initial version of the protocol.
		ValueSet *[]SymbolKind
	}

	// The client support hierarchical document symbols.
	HierarchicalDocumentSymbolSupport *bool

	// The client supports tags on `SymbolInformation`. Tags are supported on
	// `DocumentSymbol` if `hierarchicalDocumentSymbolSupport` is set to true.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0
	TagSupport *struct {

		// The tags supported by the client.
		ValueSet []SymbolTag
	}

	// The client supports an additional label presented in the UI when
	// registering a document symbol provider.
	//
	// @since 3.16.0
	LabelSupport *bool
}

/**
 * Parameters for a [DocumentSymbolRequest](#DocumentSymbolRequest).
 */
type DocumentSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier
}

/**
 * Provider options for a [DocumentSymbolRequest](#DocumentSymbolRequest).
 */
type DocumentSymbolOptions struct {
	WorkDoneProgressOptions

	// A human-readable string that is shown when multiple outlines trees
	// are shown for the same document.
	//
	// @since 3.16.0
	Label *string
}

/**
 * Registration options for a [DocumentSymbolRequest](#DocumentSymbolRequest).
 */
type DocumentSymbolRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentSymbolOptions
}

/**
 * The Client Capabilities of a [CodeActionRequest](#CodeActionRequest).
 */
type CodeActionClientCapabilities struct {

	// Whether code action supports dynamic registration.
	DynamicRegistration *bool

	// The client support code action literals of type `CodeAction` as a valid
	// response of the `textDocumentcodeAction` request. If the property is not
	// set the request can only return `Command` literals.
	//
	// @since 3.8.0
	CodeActionLiteralSupport *struct {

		// The code action kind is support with the following value
		// set.
		CodeActionKind struct {

			// The code action kind values the client supports. When this
			// property exists the client also guarantees that it will
			// handle values outside its set gracefully and falls back
			// to a default value when unknown.
			ValueSet []CodeActionKind
		}
	}

	// Whether code action supports the `isPreferred` property.
	//
	// @since 3.15.0
	IsPreferredSupport *bool

	// Whether code action supports the `disabled` property.
	//
	// @since 3.16.0
	DisabledSupport *bool

	// Whether code action supports the `data` property which is
	// preserved between a `textDocumentcodeAction` and a
	// `codeActionresolve` request.
	//
	// @since 3.16.0
	DataSupport *bool

	// Whether the client support resolving additional code action
	// properties via a separate `codeActionresolve` request.
	//
	// @since 3.16.0
	ResolveSupport *struct {

		// The properties that a client can resolve lazily.
		Properties []string
	}

	// Whether th client honors the change annotations in
	// text edits and resource operations returned via the
	// `CodeAction#edit` property by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	//
	// @since 3.16.0
	HonorsChangeAnnotations *bool
}

/**
 * The parameters of a [CodeActionRequest](#CodeActionRequest).
 */
type CodeActionParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The document in which the command was invoked.
	TextDocument TextDocumentIdentifier

	// The range for which the command was invoked.
	Range Range

	// Context carrying additional information.
	Context CodeActionContext
}

/**
 * Provider options for a [CodeActionRequest](#CodeActionRequest).
 */
type CodeActionOptions struct {
	WorkDoneProgressOptions

	// CodeActionKinds that this server may return.
	//
	// The list of kinds may be generic, such as `CodeActionKind.Refactor`, or the server
	// may list out every specific kind they provide.
	CodeActionKinds *[]CodeActionKind

	// The server provides support to resolve additional
	// information for a code action.
	//
	// @since 3.16.0
	ResolveProvider *bool
}

/**
 * Registration options for a [CodeActionRequest](#CodeActionRequest).
 */
type CodeActionRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeActionOptions
}

/**
 * Client capabilities for a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
 */
type WorkspaceSymbolClientCapabilities struct {

	// Symbol request supports dynamic registration.
	DynamicRegistration *bool

	// Specific capabilities for the `SymbolKind` in the `workspacesymbol` request.
	SymbolKind *struct {

		// The symbol kind values the client supports. When this
		// property exists the client also guarantees that it will
		// handle values outside its set gracefully and falls back
		// to a default value when unknown.
		//
		// If this property is not present the client only supports
		// the symbol kinds from `File` to `Array` as defined in
		// the initial version of the protocol.
		ValueSet *[]SymbolKind
	}

	// The client supports tags on `SymbolInformation`.
	// Clients supporting tags have to handle unknown tags gracefully.
	//
	// @since 3.16.0
	TagSupport *struct {

		// The tags supported by the client.
		ValueSet []SymbolTag
	}
}

/**
 * The parameters of a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
 */
type WorkspaceSymbolParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// A query string to filter symbols by. Clients may send an empty
	// string here to request all symbols.
	Query string
}

/**
 * Server capabilities for a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
 */
type WorkspaceSymbolOptions struct {
	WorkDoneProgressOptions
}

/**
 * Registration options for a [WorkspaceSymbolRequest](#WorkspaceSymbolRequest).
 */
type WorkspaceSymbolRegistrationOptions struct {
	WorkspaceSymbolOptions
}

/**
 * The client capabilities  of a [CodeLensRequest](#CodeLensRequest).
 */
type CodeLensClientCapabilities struct {

	// Whether code lens supports dynamic registration.
	DynamicRegistration *bool
}

/**
 * @since 3.16.0
 */
type CodeLensWorkspaceClientCapabilities struct {

	// Whether the client implementation supports a refresh request sent from the
	// server to the client.
	//
	// Note that this event is global and will force the client to refresh all
	// code lenses currently shown. It should be used with absolute care and is
	// useful for situation where a server for example detect a project wide
	// change that requires such a calculation.
	RefreshSupport *bool
}

/**
 * The parameters of a [CodeLensRequest](#CodeLensRequest).
 */
type CodeLensParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The document to request code lens for.
	TextDocument TextDocumentIdentifier
}

/**
 * Code Lens provider options of a [CodeLensRequest](#CodeLensRequest).
 */
type CodeLensOptions struct {
	WorkDoneProgressOptions

	// Code lens has a resolve provider as well.
	ResolveProvider *bool
}

/**
 * Registration options for a [CodeLensRequest](#CodeLensRequest).
 */
type CodeLensRegistrationOptions struct {
	TextDocumentRegistrationOptions
	CodeLensOptions
}

/**
 * The client capabilities of a [DocumentLinkRequest](#DocumentLinkRequest).
 */
type DocumentLinkClientCapabilities struct {

	// Whether document link supports dynamic registration.
	DynamicRegistration *bool

	// Whether the client support the `tooltip` property on `DocumentLink`.
	//
	// @since 3.15.0
	TooltipSupport *bool
}

/**
 * The parameters of a [DocumentLinkRequest](#DocumentLinkRequest).
 */
type DocumentLinkParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The document to provide document links for.
	TextDocument TextDocumentIdentifier
}

/**
 * Provider options for a [DocumentLinkRequest](#DocumentLinkRequest).
 */
type DocumentLinkOptions struct {
	WorkDoneProgressOptions

	// Document links have a resolve provider as well.
	ResolveProvider *bool
}

/**
 * Registration options for a [DocumentLinkRequest](#DocumentLinkRequest).
 */
type DocumentLinkRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentLinkOptions
}

/**
 * Client capabilities of a [DocumentFormattingRequest](#DocumentFormattingRequest).
 */
type DocumentFormattingClientCapabilities struct {

	// Whether formatting supports dynamic registration.
	DynamicRegistration *bool
}

/**
 * The parameters of a [DocumentFormattingRequest](#DocumentFormattingRequest).
 */
type DocumentFormattingParams struct {
	WorkDoneProgressParams

	// The document to format.
	TextDocument TextDocumentIdentifier

	// The format options
	Options FormattingOptions
}

/**
 * Provider options for a [DocumentFormattingRequest](#DocumentFormattingRequest).
 */
type DocumentFormattingOptions struct {
	WorkDoneProgressOptions
}

/**
 * Registration options for a [DocumentFormattingRequest](#DocumentFormattingRequest).
 */
type DocumentFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentFormattingOptions
}

/**
 * Client capabilities of a [DocumentRangeFormattingRequest](#DocumentRangeFormattingRequest).
 */
type DocumentRangeFormattingClientCapabilities struct {

	// Whether range formatting supports dynamic registration.
	DynamicRegistration *bool
}

/**
 * The parameters of a [DocumentRangeFormattingRequest](#DocumentRangeFormattingRequest).
 */
type DocumentRangeFormattingParams struct {
	WorkDoneProgressParams

	// The document to format.
	TextDocument TextDocumentIdentifier

	// The range to format
	Range Range

	// The format options
	Options FormattingOptions
}

/**
 * Provider options for a [DocumentRangeFormattingRequest](#DocumentRangeFormattingRequest).
 */
type DocumentRangeFormattingOptions struct {
	WorkDoneProgressOptions
}

/**
 * Registration options for a [DocumentRangeFormattingRequest](#DocumentRangeFormattingRequest).
 */
type DocumentRangeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentRangeFormattingOptions
}

/**
 * Client capabilities of a [DocumentOnTypeFormattingRequest](#DocumentOnTypeFormattingRequest).
 */
type DocumentOnTypeFormattingClientCapabilities struct {

	// Whether on type formatting supports dynamic registration.
	DynamicRegistration *bool
}

/**
 * The parameters of a [DocumentOnTypeFormattingRequest](#DocumentOnTypeFormattingRequest).
 */
type DocumentOnTypeFormattingParams struct {

	// The document to format.
	TextDocument TextDocumentIdentifier

	// The position at which this request was send.
	Position Position

	// The character that has been typed.
	Ch string

	// The format options.
	Options FormattingOptions
}

/**
 * Provider options for a [DocumentOnTypeFormattingRequest](#DocumentOnTypeFormattingRequest).
 */
type DocumentOnTypeFormattingOptions struct {

	// A character on which formatting should be triggered, like `}`.
	FirstTriggerCharacter string

	// More trigger characters.
	MoreTriggerCharacter *[]string
}

/**
 * Registration options for a [DocumentOnTypeFormattingRequest](#DocumentOnTypeFormattingRequest).
 */
type DocumentOnTypeFormattingRegistrationOptions struct {
	TextDocumentRegistrationOptions
	DocumentOnTypeFormattingOptions
}

type RenameClientCapabilities struct {

	// Whether rename supports dynamic registration.
	DynamicRegistration *bool

	// Client supports testing for validity of rename operations
	// before execution.
	//
	// @since 3.12.0
	PrepareSupport *bool

	// Client supports the default behavior result.
	//
	// The value indicates the default behavior used by the
	// client.
	//
	// @since 3.16.0
	PrepareSupportDefaultBehavior *PrepareSupportDefaultBehavior

	// Whether th client honors the change annotations in
	// text edits and resource operations returned via the
	// rename request's workspace edit by for example presenting
	// the workspace edit in the user interface and asking
	// for confirmation.
	//
	// @since 3.16.0
	HonorsChangeAnnotations *bool
}

/**
 * The parameters of a [RenameRequest](#RenameRequest).
 */
type RenameParams struct {
	WorkDoneProgressParams

	// The document to rename.
	TextDocument TextDocumentIdentifier

	// The position at which this request was sent.
	Position Position

	// The new name of the symbol. If the given name is not valid the
	// request must return a [ResponseError](#ResponseError) with an
	// appropriate message set.
	NewName string
}

/**
 * Provider options for a [RenameRequest](#RenameRequest).
 */
type RenameOptions struct {
	WorkDoneProgressOptions

	// Renames should be checked and tested before being executed.
	//
	// @since version 3.12.0
	PrepareProvider *bool
}

/**
 * Registration options for a [RenameRequest](#RenameRequest).
 */
type RenameRegistrationOptions struct {
	TextDocumentRegistrationOptions
	RenameOptions
}

type PrepareRenameParams struct {
	TextDocumentPositionParams
	WorkDoneProgressParams
}

/**
 * The client capabilities of a [ExecuteCommandRequest](#ExecuteCommandRequest).
 */
type ExecuteCommandClientCapabilities struct {

	// Execute command supports dynamic registration.
	DynamicRegistration *bool
}

/**
 * The parameters of a [ExecuteCommandRequest](#ExecuteCommandRequest).
 */
type ExecuteCommandParams struct {
	WorkDoneProgressParams

	// The identifier of the actual command handler.
	Command string

	// Arguments that the command should be invoked with.
	Arguments *[]interface{}
}

/**
 * The server capabilities of a [ExecuteCommandRequest](#ExecuteCommandRequest).
 */
type ExecuteCommandOptions struct {
	WorkDoneProgressOptions

	// The commands to be executed on the server
	Commands []string
}

/**
 * Registration options for a [ExecuteCommandRequest](#ExecuteCommandRequest).
 */
type ExecuteCommandRegistrationOptions struct {
	ExecuteCommandOptions
}

//---- Apply Edit request ----------------------------------------
type WorkspaceEditClientCapabilities struct {

	// The client supports versioned document changes in `WorkspaceEdit`s
	DocumentChanges *bool

	// The resource operations the client supports. Clients should at least
	// support 'create', 'rename' and 'delete' files and folders.
	//
	// @since 3.13.0
	ResourceOperations *[]ResourceOperationKind

	// The failure handling strategy of a client if applying the workspace edit
	// fails.
	//
	// @since 3.13.0
	FailureHandling *FailureHandlingKind

	// Whether the client normalizes line endings to the client specific
	// setting.
	// If set to `true` the client will normalize line ending characters
	// in a workspace edit containing to the client specific new line
	// character.
	//
	// @since 3.16.0
	NormalizesLineEndings *bool

	// Whether the client in general supports change annotations on text edits,
	// create file, rename file and delete file changes.
	//
	// @since 3.16.0
	ChangeAnnotationSupport *struct {

		// Whether the client groups edits with equal labels into tree nodes,
		// for instance all edits labelled with "Changes in Strings" would
		// be a tree node.
		GroupsOnLabel *bool
	}
}

/**
 * The parameters passed via a apply workspace edit request.
 */
type ApplyWorkspaceEditParams struct {

	// An optional label of the workspace edit. This label is
	// presented in the user interface for example on an undo
	// stack to undo the workspace edit.
	Label *string

	// The edits to apply.
	Edit WorkspaceEdit
}

/**
 * The result returned from the apply workspace edit request.
 *
 * @since 3.17 renamed from ApplyWorkspaceEditResponse
 */
type ApplyWorkspaceEditResult struct {

	// Indicates whether the edit was applied or not.
	Applied bool

	// An optional textual description for why the edit was not applied.
	// This may be used by the server for diagnostic logging or to provide
	// a suitable error for a request that triggered the edit.
	FailureReason *string

	// Depending on the client's failure handling strategy `failedChange` might
	// contain the index of the change that failed. This property is only available
	// if the client signals a `failureHandlingStrategy` in its client capabilities.
	FailedChange *uint
}

/**
 * The `client/registerCapability` request is sent from the server to the client to register a new capability
 * handler on the client side.
 */
type RegistrationRequest string

const (
	RegistrationRequestType RegistrationRequest = "new ProtocolRequestType<RegistrationParams, void, never, void, void>('client/registerCapability')"
)

/**
 * The `client/unregisterCapability` request is sent from the server to the client to unregister a previously registered capability
 * handler on the client side.
 */
type UnregistrationRequest string

const (
	UnregistrationRequestType UnregistrationRequest = "new ProtocolRequestType<UnregistrationParams, void, never, void, void>('client/unregisterCapability')"
)

type ResourceOperationKind string

const (
	/**
	 * Supports creating new files and folders.
	 */
	ResourceOperationKindCreate ResourceOperationKind = "create"
	/**
	 * Supports renaming existing files and folders.
	 */
	ResourceOperationKindRename ResourceOperationKind = "rename"
	/**
	 * Supports deleting existing files and folders.
	 */
	ResourceOperationKindDelete ResourceOperationKind = "delete"
)

type FailureHandlingKind string

const (
	/**
	 * Applying the workspace change is simply aborted if one of the changes provided
	 * fails. All operations executed before the failing operation stay executed.
	 */
	FailureHandlingKindAbort FailureHandlingKind = "abort"
	/**
	 * All operations are executed transactional. That means they either all
	 * succeed or no changes at all are applied to the workspace.
	 */
	FailureHandlingKindTransactional FailureHandlingKind = "transactional"
	/**
	 * If the workspace edit contains only textual file changes they are executed transactional.
	 * If resource changes (create, rename or delete file) are part of the change the failure
	 * handling strategy is abort.
	 */
	FailureHandlingKindTextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"
	/**
	 * The client tries to undo the operations already executed. But there is no
	 * guarantee that this is succeeding.
	 */
	FailureHandlingKindUndo FailureHandlingKind = "undo"
)

/**
 * The initialize request is sent from the client to the server.
 * It is sent once as the request after starting up the server.
 * The requests parameter is of type [InitializeParams](#InitializeParams)
 * the response if of type [InitializeResult](#InitializeResult) of a Thenable that
 * resolves to such.
 */
type InitializeRequest string

const (
	InitializeRequestType InitializeRequest = "new ProtocolRequestType<InitializeParams & WorkDoneProgressParams, InitializeResult, never, InitializeError, void>('initialize')"
)

/**
 * Known error codes for an `InitializeError`;
 */
type InitializeErrorCode int

var initializeErrorStringMap = map[InitializeErrorCode]string{
	InitializeErrorUnknownProtocolVersion: "unknownProtocolVersion",
}

func (i InitializeErrorCode) String() string {
	if s, ok := initializeErrorStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * If the protocol version provided by the client can't be handled by the server.
	 * @deprecated This initialize error got replaced by client capabilities. There is
	 * no version handshake in version 3.0x
	 */
	InitializeErrorUnknownProtocolVersion InitializeErrorCode = 1
)

/**
 * The initialized notification is sent from the client to the
 * server after the client is fully initialized and the server
 * is allowed to send requests from the server to the client.
 */
type InitializedNotification string

const (
	InitializedNotificationType InitializedNotification = "new ProtocolNotificationType<InitializedParams, void>('initialized')"
)

/**
 * A shutdown request is sent from the client to the server.
 * It is sent once when the client decides to shutdown the
 * server. The only notification that is sent after a shutdown request
 * is the exit event.
 */
type ShutdownRequest string

const (
	ShutdownRequestType ShutdownRequest = "new ProtocolRequestType0<void, never, void, void>('shutdown')"
)

/**
 * The exit event is sent from the client to the server to
 * ask the server to exit its process.
 */
type ExitNotification string

const (
	ExitNotificationType ExitNotification = "new ProtocolNotificationType0<void>('exit')"
)

/**
 * The configuration change notification is sent from the client to the server
 * when the client's configuration has changed. The notification contains
 * the changed configuration as defined by the language client.
 */
type DidChangeConfigurationNotification string

const (
	DidChangeConfigurationNotificationType DidChangeConfigurationNotification = "new ProtocolNotificationType<DidChangeConfigurationParams, DidChangeConfigurationRegistrationOptions>('workspace/didChangeConfiguration')"
)

/**
 * The message type
 */
type MessageType int

var messageTypeStringMap = map[MessageType]string{
	MessageTypeError:   "Error",
	MessageTypeWarning: "Warning",
	MessageTypeInfo:    "Info",
	MessageTypeLog:     "Log",
}

func (i MessageType) String() string {
	if s, ok := messageTypeStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * An error message.
	 */
	MessageTypeError MessageType = 1
	/**
	 * A warning message.
	 */
	MessageTypeWarning MessageType = 2
	/**
	 * An information message.
	 */
	MessageTypeInfo MessageType = 3
	/**
	 * A log message.
	 */
	MessageTypeLog MessageType = 4
)

/**
 * The show message notification is sent from a server to a client to ask
 * the client to display a particular message in the user interface.
 */
type ShowMessageNotification string

const (
	ShowMessageNotificationType ShowMessageNotification = "new ProtocolNotificationType<ShowMessageParams, void>('window/showMessage')"
)

/**
 * The show message request is sent from the server to the client to show a message
 * and a set of options actions to the user.
 */
type ShowMessageRequest string

const (
	ShowMessageRequestType ShowMessageRequest = "new ProtocolRequestType<ShowMessageRequestParams, MessageActionItem | null, never, void, void>('window/showMessageRequest')"
)

/**
 * The log message notification is sent from the server to the client to ask
 * the client to log a particular message.
 */
type LogMessageNotification string

const (
	LogMessageNotificationType LogMessageNotification = "new ProtocolNotificationType<LogMessageParams, void>('window/logMessage')"
)

/**
 * The telemetry event notification is sent from the server to the client to ask
 * the client to log telemetry data.
 */
type TelemetryEventNotification string

const (
	TelemetryEventNotificationType TelemetryEventNotification = "new ProtocolNotificationType<any, void>('telemetry/event')"
)

/**
 * Defines how the host (editor) should sync
 * document changes to the language server.
 */
type TextDocumentSyncKind int

var textDocumentSyncKindStringMap = map[TextDocumentSyncKind]string{
	TextDocumentSyncKindNone:        "None",
	TextDocumentSyncKindFull:        "Full",
	TextDocumentSyncKindIncremental: "Incremental",
}

func (i TextDocumentSyncKind) String() string {
	if s, ok := textDocumentSyncKindStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * Documents should not be synced at all.
	 */
	TextDocumentSyncKindNone TextDocumentSyncKind = 0
	/**
	 * Documents are synced by always sending the full content
	 * of the document.
	 */
	TextDocumentSyncKindFull TextDocumentSyncKind = 1
	/**
	 * Documents are synced by sending the full content on open.
	 * After that only incremental updates to the document are
	 * send.
	 */
	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
)

/**
 * The document open notification is sent from the client to the server to signal
 * newly opened text documents. The document's truth is now managed by the client
 * and the server must not try to read the document's truth using the document's
 * uri. Open in this sense means it is managed by the client. It doesn't necessarily
 * mean that its content is presented in an editor. An open notification must not
 * be sent more than once without a corresponding close notification send before.
 * This means open and close notification must be balanced and the max open count
 * is one.
 */
type DidOpenTextDocumentNotification string

const (
	DidOpenTextDocumentNotificationMethod DidOpenTextDocumentNotification = "textDocument/didOpen"

	DidOpenTextDocumentNotificationType DidOpenTextDocumentNotification = "new ProtocolNotificationType<DidOpenTextDocumentParams, TextDocumentRegistrationOptions>(method)"
)

/**
 * The document change notification is sent from the client to the server to signal
 * changes to a text document.
 */
type DidChangeTextDocumentNotification string

const (
	DidChangeTextDocumentNotificationMethod DidChangeTextDocumentNotification = "textDocument/didChange"

	DidChangeTextDocumentNotificationType DidChangeTextDocumentNotification = "new ProtocolNotificationType<DidChangeTextDocumentParams, TextDocumentChangeRegistrationOptions>(method)"
)

/**
 * The document close notification is sent from the client to the server when
 * the document got closed in the client. The document's truth now exists where
 * the document's uri points to (e.g. if the document's uri is a file uri the
 * truth now exists on disk). As with the open notification the close notification
 * is about managing the document's content. Receiving a close notification
 * doesn't mean that the document was open in an editor before. A close
 * notification requires a previous open notification to be sent.
 */
type DidCloseTextDocumentNotification string

const (
	DidCloseTextDocumentNotificationMethod DidCloseTextDocumentNotification = "textDocument/didClose"

	DidCloseTextDocumentNotificationType DidCloseTextDocumentNotification = "new ProtocolNotificationType<DidCloseTextDocumentParams, TextDocumentRegistrationOptions>(method)"
)

/**
 * The document save notification is sent from the client to the server when
 * the document got saved in the client.
 */
type DidSaveTextDocumentNotification string

const (
	DidSaveTextDocumentNotificationMethod DidSaveTextDocumentNotification = "textDocument/didSave"

	DidSaveTextDocumentNotificationType DidSaveTextDocumentNotification = "new ProtocolNotificationType<DidSaveTextDocumentParams, TextDocumentSaveRegistrationOptions>(method)"
)

/**
 * Represents reasons why a text document is saved.
 */
type TextDocumentSaveReason int

var textDocumentSaveReasonStringMap = map[TextDocumentSaveReason]string{
	TextDocumentSaveReasonManual:     "Manual",
	TextDocumentSaveReasonAfterDelay: "AfterDelay",
	TextDocumentSaveReasonFocusOut:   "FocusOut",
}

func (i TextDocumentSaveReason) String() string {
	if s, ok := textDocumentSaveReasonStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * Manually triggered, e.g. by the user pressing save, by starting debugging,
	 * or by an API call.
	 */
	TextDocumentSaveReasonManual TextDocumentSaveReason = 1
	/**
	 * Automatic after a delay.
	 */
	TextDocumentSaveReasonAfterDelay TextDocumentSaveReason = 2
	/**
	 * When the editor lost focus.
	 */
	TextDocumentSaveReasonFocusOut TextDocumentSaveReason = 3
)

/**
 * A document will save notification is sent from the client to the server before
 * the document is actually saved.
 */
type WillSaveTextDocumentNotification string

const (
	WillSaveTextDocumentNotificationMethod WillSaveTextDocumentNotification = "textDocument/willSave"

	WillSaveTextDocumentNotificationType WillSaveTextDocumentNotification = "new ProtocolNotificationType<WillSaveTextDocumentParams, TextDocumentRegistrationOptions>(method)"
)

/**
 * A document will save request is sent from the client to the server before
 * the document is actually saved. The request can return an array of TextEdits
 * which will be applied to the text document before it is saved. Please note that
 * clients might drop results if computing the text edits took too long or if a
 * server constantly fails on this request. This is done to keep the save fast and
 * reliable.
 */
type WillSaveTextDocumentWaitUntilRequest string

const (
	WillSaveTextDocumentWaitUntilRequestMethod WillSaveTextDocumentWaitUntilRequest = "textDocument/willSaveWaitUntil"

	WillSaveTextDocumentWaitUntilRequestType WillSaveTextDocumentWaitUntilRequest = "new ProtocolRequestType<WillSaveTextDocumentParams, TextEdit[] | null, never, void, TextDocumentRegistrationOptions>(method)"
)

/**
 * The watched files notification is sent from the client to the server when
 * the client detects changes to file watched by the language client.
 */
type DidChangeWatchedFilesNotification string

const (
	DidChangeWatchedFilesNotificationType DidChangeWatchedFilesNotification = "new ProtocolNotificationType<DidChangeWatchedFilesParams, DidChangeWatchedFilesRegistrationOptions>('workspace/didChangeWatchedFiles')"
)

/**
 * The file event type
 */
type FileChangeType int

var fileChangeTypeStringMap = map[FileChangeType]string{
	FileChangeTypeCreated: "Created",
	FileChangeTypeChanged: "Changed",
	FileChangeTypeDeleted: "Deleted",
}

func (i FileChangeType) String() string {
	if s, ok := fileChangeTypeStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * The file got created.
	 */
	FileChangeTypeCreated FileChangeType = 1
	/**
	 * The file got changed.
	 */
	FileChangeTypeChanged FileChangeType = 2
	/**
	 * The file got deleted.
	 */
	FileChangeTypeDeleted FileChangeType = 3
)

type WatchKind int

var watchKindStringMap = map[WatchKind]string{
	WatchKindCreate: "Create",
	WatchKindChange: "Change",
	WatchKindDelete: "Delete",
}

func (i WatchKind) String() string {
	if s, ok := watchKindStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * Interested in create events.
	 */
	WatchKindCreate WatchKind = 1
	/**
	 * Interested in change events
	 */
	WatchKindChange WatchKind = 2
	/**
	 * Interested in delete events
	 */
	WatchKindDelete WatchKind = 4
)

/**
 * Diagnostics notification are sent from the server to the client to signal
 * results of validation runs.
 */
type PublishDiagnosticsNotification string

const (
	PublishDiagnosticsNotificationType PublishDiagnosticsNotification = "new ProtocolNotificationType<PublishDiagnosticsParams, void>('textDocument/publishDiagnostics')"
)

/**
 * How a completion was triggered
 */
type CompletionTriggerKind int

var completionTriggerKindStringMap = map[CompletionTriggerKind]string{
	CompletionTriggerKindInvoked:                         "Invoked",
	CompletionTriggerKindTriggerCharacter:                "TriggerCharacter",
	CompletionTriggerKindTriggerForIncompleteCompletions: "TriggerForIncompleteCompletions",
}

func (i CompletionTriggerKind) String() string {
	if s, ok := completionTriggerKindStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * Completion was triggered by typing an identifier (24x7 code
	 * complete), manual invocation (e.g Ctrl+Space) or via API.
	 */
	CompletionTriggerKindInvoked CompletionTriggerKind = 1
	/**
	 * Completion was triggered by a trigger character specified by
	 * the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
	 */
	CompletionTriggerKindTriggerCharacter CompletionTriggerKind = 2
	/**
	 * Completion was re-triggered as current completion list is incomplete
	 */
	CompletionTriggerKindTriggerForIncompleteCompletions CompletionTriggerKind = 3
)

/**
 * Request to request completion at a given text document position. The request's
 * parameter is of type [TextDocumentPosition](#TextDocumentPosition) the response
 * is of type [CompletionItem[]](#CompletionItem) or [CompletionList](#CompletionList)
 * or a Thenable that resolves to such.
 *
 * The request can delay the computation of the [`detail`](#CompletionItem.detail)
 * and [`documentation`](#CompletionItem.documentation) properties to the `completionItem/resolve`
 * request. However, properties that are needed for the initial sorting and filtering, like `sortText`,
 * `filterText`, `insertText`, and `textEdit`, must not be changed during resolve.
 */
type CompletionRequest string

const (
	CompletionRequestMethod CompletionRequest = "textDocument/completion"

	CompletionRequestType CompletionRequest = "new ProtocolRequestType<CompletionParams, CompletionItem[] | CompletionList | null, CompletionItem[], void, CompletionRegistrationOptions>(method)"
)

/**
 * Request to resolve additional information for a given completion item.The request's
 * parameter is of type [CompletionItem](#CompletionItem) the response
 * is of type [CompletionItem](#CompletionItem) or a Thenable that resolves to such.
 */
type CompletionResolveRequest string

const (
	CompletionResolveRequestMethod CompletionResolveRequest = "completionItem/resolve"

	CompletionResolveRequestType CompletionResolveRequest = "new ProtocolRequestType<CompletionItem, CompletionItem, never, void, void>(method)"
)

/**
 * Request to request hover information at a given text document position. The request's
 * parameter is of type [TextDocumentPosition](#TextDocumentPosition) the response is of
 * type [Hover](#Hover) or a Thenable that resolves to such.
 */
type HoverRequest string

const (
	HoverRequestMethod HoverRequest = "textDocument/hover"

	HoverRequestType HoverRequest = "new ProtocolRequestType<HoverParams, Hover | null, never, void, HoverRegistrationOptions>(method)"
)

/**
 * How a signature help was triggered.
 *
 * @since 3.15.0
 */
type SignatureHelpTriggerKind int

var signatureHelpTriggerKindStringMap = map[SignatureHelpTriggerKind]string{
	SignatureHelpTriggerKindInvoked:          "Invoked",
	SignatureHelpTriggerKindTriggerCharacter: "TriggerCharacter",
	SignatureHelpTriggerKindContentChange:    "ContentChange",
}

func (i SignatureHelpTriggerKind) String() string {
	if s, ok := signatureHelpTriggerKindStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * Signature help was invoked manually by the user or by a command.
	 */
	SignatureHelpTriggerKindInvoked SignatureHelpTriggerKind = 1
	/**
	 * Signature help was triggered by a trigger character.
	 */
	SignatureHelpTriggerKindTriggerCharacter SignatureHelpTriggerKind = 2
	/**
	 * Signature help was triggered by the cursor moving or by the document content changing.
	 */
	SignatureHelpTriggerKindContentChange SignatureHelpTriggerKind = 3
)

type SignatureHelpRequest string

const (
	SignatureHelpRequestMethod SignatureHelpRequest = "textDocument/signatureHelp"

	SignatureHelpRequestType SignatureHelpRequest = "new ProtocolRequestType<SignatureHelpParams, SignatureHelp | null, never, void, SignatureHelpRegistrationOptions>(method)"
)

/**
 * A request to resolve the definition location of a symbol at a given text
 * document position. The request's parameter is of type [TextDocumentPosition]
 * (#TextDocumentPosition) the response is of either type [Definition](#Definition)
 * or a typed array of [DefinitionLink](#DefinitionLink) or a Thenable that resolves
 * to such.
 */
type DefinitionRequest string

const (
	DefinitionRequestMethod DefinitionRequest = "textDocument/definition"

	DefinitionRequestType DefinitionRequest = "new ProtocolRequestType<DefinitionParams, Definition | DefinitionLink[] | null, Location[] | DefinitionLink[], void, DefinitionRegistrationOptions>(method)"
)

/**
 * A request to resolve project-wide references for the symbol denoted
 * by the given text document position. The request's parameter is of
 * type [ReferenceParams](#ReferenceParams) the response is of type
 * [Location[]](#Location) or a Thenable that resolves to such.
 */
type ReferencesRequest string

const (
	ReferencesRequestMethod ReferencesRequest = "textDocument/references"

	ReferencesRequestType ReferencesRequest = "new ProtocolRequestType<ReferenceParams, Location[] | null, Location[], void, ReferenceRegistrationOptions>(method)"
)

/**
 * Request to resolve a [DocumentHighlight](#DocumentHighlight) for a given
 * text document position. The request's parameter is of type [TextDocumentPosition]
 * (#TextDocumentPosition) the request response is of type [DocumentHighlight[]]
 * (#DocumentHighlight) or a Thenable that resolves to such.
 */
type DocumentHighlightRequest string

const (
	DocumentHighlightRequestMethod DocumentHighlightRequest = "textDocument/documentHighlight"

	DocumentHighlightRequestType DocumentHighlightRequest = "new ProtocolRequestType<DocumentHighlightParams, DocumentHighlight[] | null, DocumentHighlight[], void, DocumentHighlightRegistrationOptions>(method)"
)

/**
 * A request to list all symbols found in a given text document. The request's
 * parameter is of type [TextDocumentIdentifier](#TextDocumentIdentifier) the
 * response is of type [SymbolInformation[]](#SymbolInformation) or a Thenable
 * that resolves to such.
 */
type DocumentSymbolRequest string

const (
	DocumentSymbolRequestMethod DocumentSymbolRequest = "textDocument/documentSymbol"

	DocumentSymbolRequestType DocumentSymbolRequest = "new ProtocolRequestType<DocumentSymbolParams, SymbolInformation[] | DocumentSymbol[] | null, SymbolInformation[] | DocumentSymbol[], void, DocumentSymbolRegistrationOptions>(method)"
)

/**
 * A request to provide commands for the given text document and range.
 */
type CodeActionRequest string

const (
	CodeActionRequestMethod CodeActionRequest = "textDocument/codeAction"

	CodeActionRequestType CodeActionRequest = "new ProtocolRequestType<CodeActionParams, (Command | CodeAction)[] | null, (Command | CodeAction)[], void, CodeActionRegistrationOptions>(method)"
)

/**
 * Request to resolve additional information for a given code action.The request's
 * parameter is of type [CodeAction](#CodeAction) the response
 * is of type [CodeAction](#CodeAction) or a Thenable that resolves to such.
 */
type CodeActionResolveRequest string

const (
	CodeActionResolveRequestMethod CodeActionResolveRequest = "codeAction/resolve"

	CodeActionResolveRequestType CodeActionResolveRequest = "new ProtocolRequestType<CodeAction, CodeAction, never, void, void>(method)"
)

/**
 * A request to list project-wide symbols matching the query string given
 * by the [WorkspaceSymbolParams](#WorkspaceSymbolParams). The response is
 * of type [SymbolInformation[]](#SymbolInformation) or a Thenable that
 * resolves to such.
 */
type WorkspaceSymbolRequest string

const (
	WorkspaceSymbolRequestMethod WorkspaceSymbolRequest = "workspace/symbol"

	WorkspaceSymbolRequestType WorkspaceSymbolRequest = "new ProtocolRequestType<WorkspaceSymbolParams, SymbolInformation[] | null, SymbolInformation[], void, WorkspaceSymbolRegistrationOptions>(method)"
)

/**
 * A request to provide code lens for the given text document.
 */
type CodeLensRequest string

const (
	CodeLensRequestMethod CodeLensRequest = "textDocument/codeLens"

	CodeLensRequestType CodeLensRequest = "new ProtocolRequestType<CodeLensParams, CodeLens[] | null, CodeLens[], void, CodeLensRegistrationOptions>(method)"
)

/**
 * A request to resolve a command for a given code lens.
 */
type CodeLensResolveRequest string

const (
	CodeLensResolveRequestMethod CodeLensResolveRequest = "codeLens/resolve"

	CodeLensResolveRequestType CodeLensResolveRequest = "new ProtocolRequestType<CodeLens, CodeLens, never, void, void>(method)"
)

/**
 * A request to refresh all code actions
 *
 * @since 3.16.0
 */
type CodeLensRefreshRequest string

const (
	CodeLensRefreshRequestMethod CodeLensRefreshRequest = "`workspace/codeLens/refresh`"

	CodeLensRefreshRequestType CodeLensRefreshRequest = "new ProtocolRequestType0<void, void, void, void>(method)"
)

/**
 * A request to provide document links
 */
type DocumentLinkRequest string

const (
	DocumentLinkRequestMethod DocumentLinkRequest = "textDocument/documentLink"

	DocumentLinkRequestType DocumentLinkRequest = "new ProtocolRequestType<DocumentLinkParams, DocumentLink[] | null, DocumentLink[], void, DocumentLinkRegistrationOptions>(method)"
)

/**
 * Request to resolve additional information for a given document link. The request's
 * parameter is of type [DocumentLink](#DocumentLink) the response
 * is of type [DocumentLink](#DocumentLink) or a Thenable that resolves to such.
 */
type DocumentLinkResolveRequest string

const (
	DocumentLinkResolveRequestMethod DocumentLinkResolveRequest = "documentLink/resolve"

	DocumentLinkResolveRequestType DocumentLinkResolveRequest = "new ProtocolRequestType<DocumentLink, DocumentLink, never, void, void>(method)"
)

/**
 * A request to to format a whole document.
 */
type DocumentFormattingRequest string

const (
	DocumentFormattingRequestMethod DocumentFormattingRequest = "textDocument/formatting"

	DocumentFormattingRequestType DocumentFormattingRequest = "new ProtocolRequestType<DocumentFormattingParams, TextEdit[] | null, never, void, DocumentFormattingRegistrationOptions>(method)"
)

/**
 * A request to to format a range in a document.
 */
type DocumentRangeFormattingRequest string

const (
	DocumentRangeFormattingRequestMethod DocumentRangeFormattingRequest = "textDocument/rangeFormatting"

	DocumentRangeFormattingRequestType DocumentRangeFormattingRequest = "new ProtocolRequestType<DocumentRangeFormattingParams, TextEdit[] | null, never, void, DocumentRangeFormattingRegistrationOptions>(method)"
)

/**
 * A request to format a document on type.
 */
type DocumentOnTypeFormattingRequest string

const (
	DocumentOnTypeFormattingRequestMethod DocumentOnTypeFormattingRequest = "textDocument/onTypeFormatting"

	DocumentOnTypeFormattingRequestType DocumentOnTypeFormattingRequest = "new ProtocolRequestType<DocumentOnTypeFormattingParams, TextEdit[] | null, never, void, DocumentOnTypeFormattingRegistrationOptions>(method)"
)

//---- Rename ----------------------------------------------
type PrepareSupportDefaultBehavior int

var prepareSupportDefaultBehaviorStringMap = map[PrepareSupportDefaultBehavior]string{
	PrepareSupportDefaultBehaviorIdentifier: "Identifier",
}

func (i PrepareSupportDefaultBehavior) String() string {
	if s, ok := prepareSupportDefaultBehaviorStringMap[i]; ok {
		return s
	}
	return "unknown"
}

const (
	/**
	 * The client's default behavior is to select the identifier
	 * according the to language's syntax rule.
	 */
	PrepareSupportDefaultBehaviorIdentifier PrepareSupportDefaultBehavior = 1
)

/**
 * A request to rename a symbol.
 */
type RenameRequest string

const (
	RenameRequestMethod RenameRequest = "textDocument/rename"

	RenameRequestType RenameRequest = "new ProtocolRequestType<RenameParams, WorkspaceEdit | null, never, void, RenameRegistrationOptions>(method)"
)

/**
 * A request to test and perform the setup necessary for a rename.
 *
 * @since 3.16 - support for default behavior
 */
type PrepareRenameRequest string

const (
	PrepareRenameRequestMethod PrepareRenameRequest = "textDocument/prepareRename"

	PrepareRenameRequestType PrepareRenameRequest = "new ProtocolRequestType<PrepareRenameParams, Range | { range: Range, placeholder: string } | { defaultBehavior: boolean } | null, never, void, void>(method)"
)

/**
 * A request send from the client to the server to execute a command. The request might return
 * a workspace edit which the client will apply to the workspace.
 */
type ExecuteCommandRequest string

const (
	ExecuteCommandRequestType ExecuteCommandRequest = "new ProtocolRequestType<ExecuteCommandParams, any | null, never, void, ExecuteCommandRegistrationOptions>('workspace/executeCommand')"
)

/**
 * A request sent from the server to the client to modified certain resources.
 */
type ApplyWorkspaceEditRequest string

const (
	ApplyWorkspaceEditRequestType ApplyWorkspaceEditRequest = "new ProtocolRequestType<ApplyWorkspaceEditParams, ApplyWorkspaceEditResult, never, void, void>('workspace/applyEdit')"
)
