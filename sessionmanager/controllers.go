package sessionmanager

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	ErrNoUserID  = errors.New("UserID Not Provided")
	ErrBadUserID = errors.New("UserID is not of UUID type")
)

type sessioncontrollers struct {
	HandleCreateSession     *httptransport.Server
	HandleGetSessionDetails *httptransport.Server
	HandleVerifySession     *httptransport.Server
}

// exported, http implementation of endpoints
func NewSessionController(svc SessionService, ctx context.Context) sessioncontrollers {
	return sessioncontrollers{
		HandleCreateSession: httptransport.NewServer(
			makeCreateSessionEndpoint(svc, ctx),
			decodeRequest,
			encodeResponse,
		),
		HandleGetSessionDetails: nil,
		HandleVerifySession: httptransport.NewServer(
			makeIsSessionActiveRequest(svc, ctx),
			decodeRequest,
			encodeResponse,
		),
	}
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request Session
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
