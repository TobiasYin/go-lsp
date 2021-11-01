package defines

type WorkspaceDocumentDiagnosticReport interface{} // WorkspaceFullDocumentDiagnosticReport | WorkspaceUnchangedDocumentDiagnosticReport;

/**
 * @since 3.17.0 - proposed state
 */
type DiagnosticClientCapabilities struct {

    // Whether implementation supports dynamic registration. If this is set to `true`
    // the client supports the new `(TextDocumentRegistrationOptions & StaticRegistrationOptions)`
    // return value for the corresponding server capability as well.
    DynamicRegistration *bool

    // Whether the clients supports related documents for document diagnostic pulls.
    RelatedDocumentSupport *bool
}

/**
 * Diagnostic options.
 *
 * @since 3.17.0 - proposed state
 */
type DiagnosticOptions struct {
    WorkDoneProgressOptions


    // An optional identifier under which the diagnostics are
    // managed by the client.
    Identifier *string

    // Whether the language has inter file dependencies meaning that
    // editing code in one file can result in a different diagnostic
    // set in another file. Inter file dependencies are common for
    // most programming languages and typically uncommon for linters.
    InterFileDependencies bool

    // The server provides support for workspace diagnostics as well.
    WorkspaceDiagnostics bool
}

/**
 * Diagnostic registration options.
 *
 * @since 3.17.0 - proposed state
 */
type DiagnosticRegistrationOptions struct {
    TextDocumentRegistrationOptions
    DiagnosticOptions
    StaticRegistrationOptions

}

/**
 * Cancellation data returned from a diagnostic request.
 *
 * @since 3.17.0 - proposed state
 */
type DiagnosticServerCancellationData struct {

    RetriggerRequest bool
}

/**
 * Parameters of the document diagnostic request.
 *
 * @since 3.17.0 - proposed state
 */
type DocumentDiagnosticParams struct {
    WorkDoneProgressParams
    PartialResultParams


    // The text document.
    TextDocument TextDocumentIdentifier

    // The additional identifier  provided during registration.
    Identifier *string

    // The result id of a previous response if provided.
    PreviousResultId *string
}

/**
 * Parameters of the workspace diagnostic request.
 *
 * @since 3.17.0 - proposed state
 */
type WorkspaceDiagnosticParams struct {
    WorkDoneProgressParams
    PartialResultParams


    // The additional identifier provided during registration.
    Identifier *string

    // The currently known diagnostic reports with their
    // previous result ids.
    PreviousResultIds []PreviousResultId
}

type PreviousResultId struct {
    Uri   URI
    Value string
}

type FullDocumentDiagnosticReport struct {
    Kind     interface{} // DocumentDiagnosticReportKind.full
    ResultId *string
    Items    []interface{}
}

/**
 * A full document diagnostic report for a workspace diagnostic result.
 *
 * @since 3.17.0 - proposed state
 */
type WorkspaceFullDocumentDiagnosticReport struct {
    FullDocumentDiagnosticReport


    // The URI for which diagnostic information is reported.
    Uri DocumentUri

    // The version number for which the diagnostics are reported.
    // If the document is not marked as open `null` can be provided.
    Version interface{}  // int, null, 
}

type UnchangedDocumentDiagnosticReport struct {
    Kind     interface{} //DocumentDiagnosticReportKind.unChanged
    ResultId string
}


/**
 * An unchanged document diagnostic report for a workspace diagnostic result.
 *
 * @since 3.17.0 - proposed state
 */
type WorkspaceUnchangedDocumentDiagnosticReport struct {
    UnchangedDocumentDiagnosticReport


    // The URI for which diagnostic information is reported.
    Uri DocumentUri

    // The version number for which the diagnostics are reported.
    // If the document is not marked as open `null` can be provided.
    Version interface{}  // int, null, 
}

/**
 * A workspace diagnostic report.
 *
 * @since 3.17.0 - proposed state
 */
type WorkspaceDiagnosticReport struct {

    Items []WorkspaceDocumentDiagnosticReport
}

/**
 * A partial result for a workspace diagnostic report.
 *
 * @since 3.17.0 - proposed state
 */
type WorkspaceDiagnosticReportPartialResult struct {

    Items []WorkspaceDocumentDiagnosticReport
}

/**
 * The document diagnostic report kinds.
 *
 * @since 3.17.0 - proposed state
 */
type DocumentDiagnosticReportKind string
const (
    /**
     * A diagnostic report with a full
     * set of problems.
     */
    DocumentDiagnosticReportKindFull DocumentDiagnosticReportKind = "full"
    /**
     * A report indicating that the last
     * returned report is still accurate.
     */
    DocumentDiagnosticReportKindUnChanged DocumentDiagnosticReportKind = "unChanged"
)