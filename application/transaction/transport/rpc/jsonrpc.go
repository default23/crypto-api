package rpc

import (
	"context"
	"net/http"

	"github.com/default23/crypto-api/application/transaction/endpoint"
	"github.com/default23/crypto-api/infrastructure/logging"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/kit/transport/http/jsonrpc"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	SignHandler http.Handler
}

// NewJSONRPCHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewJSONRPCHandler(endpoints endpoint.Endpoints, logger *logrus.Entry) *jsonrpc.Server {
	handler := jsonrpc.NewServer(
		makeEndpointCodecMap(endpoints),
		jsonrpc.ServerBefore(withLogger(logger)),
	)

	return handler
}

func withLogger(logger *logrus.Entry) httptransport.RequestFunc {
	return func(ctx context.Context, request *http.Request) context.Context {
		return logging.WithLogger(ctx, logger)
	}
}

// makeEndpointCodecMap returns a codec map configured for the transaction.
func makeEndpointCodecMap(endpoints endpoint.Endpoints) jsonrpc.EndpointCodecMap {
	return jsonrpc.EndpointCodecMap{
		"sign_transaction": jsonrpc.EndpointCodec{
			Endpoint: endpoints.SignEndpoint,
			Decode:   decodeSignRequest,
			Encode:   encodeSignResponse,
		},
	}
}
