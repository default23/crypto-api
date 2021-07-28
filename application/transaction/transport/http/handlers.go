package http

import (
	"context"
	"encoding/json"
	"net/http"

	kitendpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/default23/crypto-api/application/transaction/endpoint"
)

type Handlers struct {
	SignHandler http.Handler
}

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoint.Endpoints) Handlers {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(httptransport.DefaultErrorEncoder),
	}

	return Handlers{
		SignHandler: makeSignHandler(endpoints, options),
	}
}

// Register add handlers to mux.Router with theirs paths.
func (h *Handlers) Register(router *mux.Router) {
	router.Methods("POST").Path("/sign_transaction").Handler(h.SignHandler)
}

// encodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitendpoint.Failer); ok && f.Failed() != nil {
		httptransport.DefaultErrorEncoder(ctx, f.Failed(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
