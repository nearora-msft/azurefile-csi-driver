package virtualnetworklinksclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
)

type Interface interface {

	// CreateOrUpdate creates or updates a private dns zone.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters privatedns.VirtualNetworkLink, waitForCompletion bool) error
}
