package defines



//---- Client capability ----
type DocumentColorClientCapabilities struct {

    // Whether implementation supports dynamic registration. If this is set to `true`
    // the client supports the new `DocumentColorRegistrationOptions` return value
    // for the corresponding server capability as well.
    DynamicRegistration *bool
}

type DocumentColorOptions struct {
    WorkDoneProgressOptions

}

type DocumentColorRegistrationOptions struct {
    TextDocumentRegistrationOptions
    StaticRegistrationOptions
    DocumentColorOptions

}

/**
 * Parameters for a [DocumentColorRequest](#DocumentColorRequest).
 */
type DocumentColorParams struct {
    WorkDoneProgressParams
    PartialResultParams


    // The text document.
    TextDocument TextDocumentIdentifier
}

/**
 * Parameters for a [ColorPresentationRequest](#ColorPresentationRequest).
 */
type ColorPresentationParams struct {
    WorkDoneProgressParams
    PartialResultParams


    // The text document.
    TextDocument TextDocumentIdentifier

    // The color to request presentations for.
    Color Color

    // The range where the color would be inserted. Serves as a context.
    Range Range
}

