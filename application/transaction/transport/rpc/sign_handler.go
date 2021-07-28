package rpc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/default23/crypto-api/application/transaction/endpoint"
	"github.com/default23/crypto-api/domain"
	"github.com/go-kit/kit/transport/http/jsonrpc"
)

func decodeSignRequest(_ context.Context, msg json.RawMessage) (interface{}, error) {
	var req endpoint.SignRequest
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return nil, &jsonrpc.Error{
			Code:    -32000,
			Message: fmt.Sprintf("couldn't unmarshal body to sign request: %s", err),
		}
	}
	req.Request = msg

	_, err = domain.NewGate(req.Gate)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func encodeSignResponse(_ context.Context, obj interface{}) (json.RawMessage, error) {
	res, ok := obj.(endpoint.SignResponse)
	if !ok {
		return nil, &jsonrpc.Error{
			Code:    -32000,
			Message: fmt.Sprintf("Asserting result to SignResponse failed. Got %T, %+v", obj, obj),
		}
	}

	if err := res.Failed(); err != nil {
		return nil, &jsonrpc.Error{
			Code:    -32000,
			Message: err.Error(),
		}
	}

	b, err := json.Marshal(res)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal response: %s", err)
	}

	return b, nil
}
