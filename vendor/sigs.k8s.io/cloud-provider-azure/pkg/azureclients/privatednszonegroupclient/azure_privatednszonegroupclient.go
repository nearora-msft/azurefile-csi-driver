package privatednszonegroupclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
	"k8s.io/klog/v2"
	azclients "sigs.k8s.io/cloud-provider-azure/pkg/azureclients"
)

var _ Interface = &Client{}

// Client implements privatednszonegroupclient client Interface.
type Client struct {
	privateDnsZoneGroupClient network.PrivateDNSZoneGroupsClient
}

// New creates a new private dns zone group client.
func New(config *azclients.ClientConfig) *Client {
	privateDnsZoneGroupClient := network.NewPrivateDNSZoneGroupsClient(config.SubscriptionID)
	privateDnsZoneGroupClient.Authorizer = config.Authorizer
	client := &Client{
		privateDnsZoneGroupClient: privateDnsZoneGroupClient,
	}
	return client
}

// CreateOrUpdate creates or updates a private dns zone group
func (c *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateEndpointName string, privateDNSZoneGroupName string, parameters network.PrivateDNSZoneGroup, waitForCompletion bool) error {
	privateDNSZoneGroupFuture, err := c.privateDnsZoneGroupClient.CreateOrUpdate(ctx, resourceGroupName, privateEndpointName, privateDNSZoneGroupName, parameters)
	if err != nil {
		klog.V(5).Infof("Received error for %s, resourceGroup: %s, privateEndpointName: %s, error: %s", "privatednszonegroup.create.request", resourceGroupName, privateEndpointName, err)
		return err
	}
	if waitForCompletion {
		err = privateDNSZoneGroupFuture.WaitForCompletionRef(ctx, c.privateDnsZoneGroupClient.Client)
		if err != nil {
			klog.V(5).Infof("Received error while waiting for completion for %s, resourceGroup: %s, privateEndpointName: %s, error: %s", "privatednszonegroup.create.request", resourceGroupName, privateEndpointName, err)
			return err
		}
	}
	return nil
}
