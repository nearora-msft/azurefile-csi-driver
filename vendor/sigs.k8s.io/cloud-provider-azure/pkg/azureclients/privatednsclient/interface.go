package privatednsclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
)

// Interface is the client interface for Private DNS Zones
type Interface interface {

	// CreateOrUpdate creates or updates a private dns zone.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, parameters privatedns.PrivateZone, waitForCompletion bool) error
}
