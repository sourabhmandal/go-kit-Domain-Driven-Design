package strconv

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
)

type StringControllers struct {
	UppercaseHandler 				*httptransport.Server
	CountHandler 						*httptransport.Server
}



// referenced in api package to access
// these HTTP implements
func StrconvController(svc StringService) StringControllers {
	logger := log.NewLogfmtLogger(os.Stderr)

	var uppercase endpoint.Endpoint
	uppercase = makeUppercaseEndpoint(svc)
	uppercase = loggingMiddleware(
		log.With(logger, "method", "uppercase"),
		)(uppercase)

	var count endpoint.Endpoint
	count = makeCountEndpoint(svc)
	count = loggingMiddleware(
		log.With(logger, "method", "count"),
		)(count)


	uppercaseHandler := httptransport.NewServer(
		uppercase,
		decodeUppercaseRequest,
		encodeResponse,
	)
	
	countHandler := httptransport.NewServer(
		count,
		decodeCountRequest,
		encodeResponse,
	)
	
	return StringControllers{
		UppercaseHandler: uppercaseHandler,
		CountHandler : countHandler,
	}
}



// encoder decoder functions for
// the controller to parse json
func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}