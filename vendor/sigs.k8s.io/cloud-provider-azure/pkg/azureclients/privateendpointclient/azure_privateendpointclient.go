/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package privateendpointclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-02-01/network"
	"k8s.io/klog/v2"

	azclients "sigs.k8s.io/cloud-provider-azure/pkg/azureclients"
)

var _ Interface = &Client{}

// Client implements privateendpoint client Interface.
type Client struct {
	privateEndpointClient network.PrivateEndpointsClient
	resourceGroup         string
	location              string
}

// New creates a new privateendpoint client.
func New(config *azclients.ClientConfig) *Client {
	klog.V(2).Infof("Create a private endpoint client")
	privateEndpointClient := network.NewPrivateEndpointsClient(config.SubscriptionID)
	privateEndpointClient.Authorizer = config.Authorizer

	client := &Client{
		privateEndpointClient: privateEndpointClient,
	}
	klog.V(2).Infof("Created a private endpoint client successfully")
	return client
}

// CreateOrUpdate creates or updates a Subnet.
func (c *Client) CreateOrUpdate(ctx context.Context, resourceId *string, endpointName string, properties *network.PrivateEndpointProperties) error {
	klog.V(2).Infof("Create a private endpoint")
	privateEndpoint := network.PrivateEndpoint{
		ID:                        resourceId,
		Location:                  &c.location,
		PrivateEndpointProperties: properties,
	}
	_, err := c.privateEndpointClient.CreateOrUpdate(ctx, c.resourceGroup, endpointName, privateEndpoint)
	if err != nil {
		return err
	}
	return nil
}
