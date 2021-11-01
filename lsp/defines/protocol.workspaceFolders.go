package defines



type WorkspaceFoldersInitializeParams struct {

    // The actual configured workspace folders.
    WorkspaceFolders interface{}  // []WorkspaceFolder, null, 
}

type WorkspaceFoldersClientCapabilities struct {

    // The workspace client capabilities
    Workspace *struct {

    // The client has support for workspace folders
    // 
    // @since 3.6.0
    WorkspaceFolders *bool
}
}

type WorkspaceFoldersServerCapabilities struct {

    // The workspace server capabilities
    Workspace *struct {

    WorkspaceFolders interface{}  // supported, changeNotifications, 
}
}

type WorkspaceFolder struct {

    // The associated URI for this workspace folder.
    Uri string

    // The name of the workspace folder. Used to refer to this
    // workspace folder in the user interface.
    Name string
}

/**
 * The parameters of a `workspace/didChangeWorkspaceFolders` notification.
 */
type DidChangeWorkspaceFoldersParams struct {

    // The actual workspace folder change event.
    Event WorkspaceFoldersChangeEvent
}

/**
 * The workspace folder change event.
 */
type WorkspaceFoldersChangeEvent struct {

    // The array of added workspace folders
    Added []WorkspaceFolder

    // The array of the removed workspace folders
    Removed []WorkspaceFolder
}

