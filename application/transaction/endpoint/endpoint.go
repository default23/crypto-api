package endpoint

import (
	kitendpoint "github.com/go-kit/kit/endpoint"

	"github.com/default23/crypto-api/application/transaction"
)

// Endpoints is set that collects all of the endpoints that must be initialized
// with service.
type Endpoints struct {
	SignEndpoint kitendpoint.Endpoint
}

func NewEndpoints(service transaction.UseCase) Endpoints {
	return Endpoints{
		SignEndpoint: MakeSignEndpoint(service),
	}
}
