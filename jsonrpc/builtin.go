package jsonrpc

import "context"

// $/cancelRequest
type cancelParams struct {
	ID interface{} `json:"id"`
}

func cancelRequest(ctx context.Context, req interface{}, resp interface{}) error {
	params := req.(*cancelParams)
	if params.ID == nil {
		return nil
	}
	session := getSession(ctx)
	session.cancelJob(params.ID)
	return nil
}
func CancelRequest() MethodInfo {
	return MethodInfo{
		Name: "$/cancelRequest",
		NewRequest: func() interface{} {
			return &cancelParams{}
		},
		NewResponse: func() interface{} {
			return nil
		},
		Handler: cancelRequest,
	}
}

// $/progress
// TODO not implemented yet
