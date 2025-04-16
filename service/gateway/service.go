package main

import (
	"context"
	"encoding/json"
	"s0611-23/internal/demobase"
	"s0611-23/internal/proxyproto"
	"strconv"
)

// Service ...
type Service struct {
	proxyproto.UnimplementedCentrifugoProxyServer
	query demobase.Querier
}

// NewService ...
func NewService(q demobase.Querier) *Service {
	return &Service{
		query: q,
	}
}

// Connect ...
func (s *Service) Connect(ctx context.Context, req *proxyproto.ConnectRequest) (*proxyproto.ConnectResponse, error) {
	type AuthRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	authRequest := &AuthRequest{}

	if err := json.Unmarshal(req.Data, authRequest); err != nil {
		return respondError(107, "bad request"), nil
	}

	account, err := s.query.UserLogin(context.Background(), demobase.UserLoginParams{
		Username: authRequest.Username,
		Password: authRequest.Password,
	})

	if err != nil {
		return respondError(101, "unauthorized"), nil
	}

	if !account.Enabled {
		return respondError(101, "unauthorized"), nil
	}

	return respondResult(strconv.Itoa(int(account.ID))), nil
}
