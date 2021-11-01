package protocol



/**
 * Client capabilities for the show document request.
 *
 * @since 3.16.0
 */
type ShowDocumentClientCapabilities struct {

    // The client has support for the show document
    // request.
    Support bool
}

/**
 * Params to show a document.
 *
 * @since 3.16.0
 */
type ShowDocumentParams struct {

    // The document uri to show.
    Uri URI

    // Indicates to show the resource in an external program.
    // To show for example `https:code.visualstudio.com`
    // in the default WEB browser set `external` to `true`.
    External *bool

    // An optional property to indicate whether the editor
    // showing the document should take focus or not.
    // Clients might ignore this property if an external
    // program in started.
    TakeFocus *bool

    // An optional selection range if the document is a text
    // document. Clients might ignore the property if an
    // external program is started or the file is not a text
    // file.
    Selection *Range
}

/**
 * The result of an show document request.
 *
 * @since 3.16.0
 */
type ShowDocumentResult struct {

    // A boolean indicating if the show was successful.
    Success bool
}

