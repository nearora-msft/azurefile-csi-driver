package privatednsclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
	"k8s.io/klog/v2"
	azclients "sigs.k8s.io/cloud-provider-azure/pkg/azureclients"
)

var _ Interface = &Client{}

// Client implements privatednsclient Interface.
type Client struct {
	privatednsclient privatedns.PrivateZonesClient
}

// New creates a new privatedns client.
func New(config *azclients.ClientConfig) *Client {
	privateDnsClient := privatedns.NewPrivateZonesClient(config.SubscriptionID)
	privateDnsClient.Authorizer = config.Authorizer
	client := &Client{
		privatednsclient: privateDnsClient,
	}
	return client
}

// CreateOrUpdate creates or updates a private dns zone
func (c *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, parameters privatedns.PrivateZone, waitForCompletion bool) error {
	createOrUpdateFuture, err := c.privatednsclient.CreateOrUpdate(ctx, resourceGroupName, privateZoneName, parameters, "", "*")

	if err != nil {
		klog.V(5).Infof("Received error for %s, resourceGroup: %s, error: %s", "privatedns.create.request", resourceGroupName, err)
		return err
	}

	if waitForCompletion {
		err := createOrUpdateFuture.WaitForCompletionRef(ctx, c.privatednsclient.Client)
		if err != nil {
			klog.V(5).Infof("Received error while waiting for completion for %s, resourceGroup: %s, error: %s", "privatedns.create.request", resourceGroupName, err)
			return err
		}
	}
	return nil
}
