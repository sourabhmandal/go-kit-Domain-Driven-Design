package sessionmanager

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

/*
 * These are the message types that
 * is used by transport layer (HTTP, gRPC)
 * i.e
 * These are type definitions for
 * Request and Response data over transport/network layer
 */
type (
	createSessionRequest struct {
		UserID string `json:"userId"`
	}

	createSessionResponse struct {
		Ok string `json:"ok"`
	}

	isSessionActiveRequest struct {
		UserID string `json:"userId"`
	}

	isSessionActiveResponse struct {
		IsActive bool `json:"isActive"`
	}
)

func makeCreateSessionEndpoint(svc SessionService, ctx context.Context) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createSessionRequest)
		ok, err := svc.CreateSession(ctx, req.UserID)
		return createSessionResponse{Ok: ok}, err

	}
}

func makeIsSessionActiveRequest(svc SessionService, ctx context.Context) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(isSessionActiveRequest)
		ok, err := svc.IsActiveSession(ctx, req.UserID)
		return isSessionActiveResponse{IsActive: ok}, err
	}
}