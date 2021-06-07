package privateendpointclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
)

// Interface is the client interface for Private Endpoints.
type Interface interface {

	// CreateOrUpdate creates or updates a private endpoint.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, endpointName string, privateEndpoint network.PrivateEndpoint, waitForCompletion bool) error
}
