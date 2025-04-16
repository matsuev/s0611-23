package main

import "s0611-23/internal/proxyproto"

func respondError(code uint32, message string) *proxyproto.ConnectResponse {
	return &proxyproto.ConnectResponse{
		Error: &proxyproto.Error{
			Code:    code,
			Message: message,
		},
	}
}

func respondResult(user string) *proxyproto.ConnectResponse {
	return &proxyproto.ConnectResponse{
		Result: &proxyproto.ConnectResult{
			User: user,
		},
	}
}
