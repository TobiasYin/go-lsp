package protocol



type WorkDoneProgressClientCapabilities struct {

    // Window specific client capabilities.
    Window *struct {

    // Whether client supports server initiated progress using the
    // `windowworkDoneProgresscreate` request.
    // 
    // Since 3.15.0
    WorkDoneProgress *bool
}
}

type WorkDoneProgressBegin struct {

    Kind interface{} // 'begin'

    // Mandatory title of the progress operation. Used to briefly inform about
    // the kind of operation being performed.
    // 
    // Examples: "Indexing" or "Linking dependencies".
    Title string

    // Controls if a cancel button should show to allow the user to cancel the
    // long running operation. Clients that don't support cancellation are allowed
    // to ignore the setting.
    Cancellable *bool

    // Optional, more detailed associated progress message. Contains
    // complementary information to the `title`.
    // 
    // Examples: "325 files", "projectsrcmodule2", "node_modulessome_dep".
    // If unset, the previous progress message (if any) is still valid.
    Message *string

    // Optional progress percentage to display (value 100 is considered 100%).
    // If not provided infinite progress is assumed and clients are allowed
    // to ignore the `percentage` value in subsequent in report notifications.
    // 
    // The value should be steadily rising. Clients are free to ignore values
    // that are not following this rule. The value range is [0, 100].
    Percentage *uint
}

type WorkDoneProgressReport struct {

    Kind interface{} // 'report'

    // Controls enablement state of a cancel button.
    // 
    // Clients that don't support cancellation or don't support controlling the button's
    // enablement state are allowed to ignore the property.
    Cancellable *bool

    // Optional, more detailed associated progress message. Contains
    // complementary information to the `title`.
    // 
    // Examples: "325 files", "projectsrcmodule2", "node_modulessome_dep".
    // If unset, the previous progress message (if any) is still valid.
    Message *string

    // Optional progress percentage to display (value 100 is considered 100%).
    // If not provided infinite progress is assumed and clients are allowed
    // to ignore the `percentage` value in subsequent in report notifications.
    // 
    // The value should be steadily rising. Clients are free to ignore values
    // that are not following this rule. The value range is [0, 100]
    Percentage *uint
}

type WorkDoneProgressEnd struct {

    Kind interface{} // 'end'

    // Optional, a final message indicating to for example indicate the outcome
    // of the operation.
    Message *string
}

type WorkDoneProgressCreateParams struct {

    // The token to be used to report progress.
    Token ProgressToken
}

type WorkDoneProgressCancelParams struct {

    // The token to be used to report progress.
    Token ProgressToken
}

