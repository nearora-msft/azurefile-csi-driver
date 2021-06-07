package virtualnetworklinksclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
	"k8s.io/klog/v2"
	azclients "sigs.k8s.io/cloud-provider-azure/pkg/azureclients"
)

var _ Interface = &Client{}

// Client implements virtualnetworklinksclient Interface.
type Client struct {
	virtualNetworkLinksClient privatedns.VirtualNetworkLinksClient
}

// New creates a new virtualnetworklinks client.
func New(config *azclients.ClientConfig) *Client {
	virtualNetworkLinksClient := privatedns.NewVirtualNetworkLinksClient(config.SubscriptionID)
	virtualNetworkLinksClient.Authorizer = config.Authorizer

	client := &Client{
		virtualNetworkLinksClient: virtualNetworkLinksClient,
	}
	return client
}

// CreateOrUpdate creates or updates a virtual network
func (c *Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, privateZoneName string, virtualNetworkLinkName string, parameters privatedns.VirtualNetworkLink, waitForCompletion bool) error {
	vNetLinksFuture, err := c.virtualNetworkLinksClient.CreateOrUpdate(ctx, resourceGroupName, privateZoneName, virtualNetworkLinkName, parameters, "", "*")
	if err != nil {
		klog.V(5).Infof("Received error for %s, resourceGroup: %s, privateZoneName: %s, error: %s", "virtualnetworklinks.create.request", resourceGroupName, privateZoneName, err)
		return err
	}
	if waitForCompletion {
		err := vNetLinksFuture.WaitForCompletionRef(ctx, c.virtualNetworkLinksClient.Client)
		if err != nil {
			klog.V(5).Infof("Received error while waiting for completion for %s, resourceGroup: %s, privateZoneName: %s, error: %s", "virtualnetworklinks.create.request", resourceGroupName, privateZoneName, err)
			return err
		}
	}
	return nil
}
