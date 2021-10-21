package api

import (
	"context"
	"zauth/sessionmanager"
	"zauth/strconv"

	"github.com/gorilla/mux"
)

func RegisterStrconvRoutes(svc strconv.StringService, r *mux.Router) {
	strconvHandlers := strconv.StrconvController(svc)
	r.Methods("POST").Path("/uppercase").Handler(strconvHandlers.UppercaseHandler)
	r.Methods("POST").Path("/count").Handler(strconvHandlers.CountHandler)
}

func RegisterSessionRoutes(svc sessionmanager.SessionService,
	ctx context.Context, r *mux.Router) {
	sessionmanagerHandlers := sessionmanager.NewSessionController(svc, ctx)

	r.Methods("POST").Path("/session").Handler(sessionmanagerHandlers.HandleCreateSession)
	r.Methods("GET").Path("/session/{userId}/verify").Handler(sessionmanagerHandlers.HandleVerifySession)
}
