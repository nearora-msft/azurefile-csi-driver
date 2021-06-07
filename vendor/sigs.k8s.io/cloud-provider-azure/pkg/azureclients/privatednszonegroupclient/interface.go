package privatednszonegroupclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
)

type Interface interface {

	// CreateOrUpdate creates or updates a private dns zone group endpoint.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters network.PrivateDNSZoneGroup, waitForCompletion bool) error
}
