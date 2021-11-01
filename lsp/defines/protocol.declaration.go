package defines



/**
 * @since 3.14.0
 */
type DeclarationClientCapabilities struct {

    // Whether declaration supports dynamic registration. If this is set to `true`
    // the client supports the new `DeclarationRegistrationOptions` return value
    // for the corresponding server capability as well.
    DynamicRegistration *bool

    // The client supports additional metadata in the form of declaration links.
    LinkSupport *bool
}

type DeclarationOptions struct {
    WorkDoneProgressOptions

}

type DeclarationRegistrationOptions struct {
    DeclarationOptions
    TextDocumentRegistrationOptions
    StaticRegistrationOptions

}

type DeclarationParams struct {
    TextDocumentPositionParams
    WorkDoneProgressParams
    PartialResultParams

}

