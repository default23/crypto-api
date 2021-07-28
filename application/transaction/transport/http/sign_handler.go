package http

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/default23/crypto-api/domain"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/default23/crypto-api/application/transaction/endpoint"
)

func makeSignHandler(endpoints endpoint.Endpoints, options []httptransport.ServerOption) http.Handler {
	return httptransport.NewServer(
		endpoints.SignEndpoint,
		decodeSignRequest,
		encodeHTTPGenericResponse,
		options...,
	)
}

// decodeHTTPSumRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded sum request from the HTTP request body. Primarily useful in a
// server.
func decodeSignRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.SignRequest

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &req)
	if err != nil {
		return nil, err
	}
	req.Request = data

	_, err = domain.NewGate(req.Gate)
	if err != nil {
		return nil, err
	}

	return req, err
}
