package protocol



//---- Get Configuration request ----
type ConfigurationClientCapabilities struct {

    // The workspace client capabilities
    Workspace *struct {

    // The client supports `workspaceconfiguration` requests.
    // 
    // @since 3.6.0
    Configuration *bool
}
}

type ConfigurationItem struct {

    // The scope to get the configuration section for.
    ScopeUri *string

    // The configuration section asked for.
    Section *string
}

/**
 * The parameters of a configuration request.
 */
type ConfigurationParams struct {

    Items []ConfigurationItem
}

