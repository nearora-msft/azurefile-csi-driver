package privateendpointclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
	"k8s.io/klog/v2"
	azclients "sigs.k8s.io/cloud-provider-azure/pkg/azureclients"
)

var _ Interface = &Client{}

// Client implements privateendpointclient Interface.
type Client struct {
	privateEndpointClient network.PrivateEndpointsClient
	location              string
}

// New creates a new private endpoint client.
func New(config *azclients.ClientConfig) *Client {
	privateEndpointClient := network.NewPrivateEndpointsClient(config.SubscriptionID)
	privateEndpointClient.Authorizer = config.Authorizer

	client := &Client{
		privateEndpointClient: privateEndpointClient,
	}
	return client
}

// CreateOrUpdate creates or updates a private endpoint.
func (c *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, endpointName string, privateEndpoint network.PrivateEndpoint, waitForCompletion bool) error {
	privateEndpointFuture, err := c.privateEndpointClient.CreateOrUpdate(ctx, resourceGroupName, endpointName, privateEndpoint)
	if err != nil {
		klog.V(5).Infof("Received error for %s, resourceGroup: %s, error: %s", "privateendpoint.create.request", resourceGroupName, err)
		return err
	}
	if waitForCompletion {
		err = privateEndpointFuture.WaitForCompletionRef(ctx, c.privateEndpointClient.Client)
		if err != nil {
			klog.V(5).Infof("Received error while waiting for completion for %s, resourceGroup: %s, error: %s", "privateendpoint.create.request", resourceGroupName, err)
			return err
		}

	}
	return nil
}
