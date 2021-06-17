// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/networksecurity/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkSecurityServerTLSPolicies implements NetworkSecurityServerTLSPolicyInterface
type FakeNetworkSecurityServerTLSPolicies struct {
	Fake *FakeNetworksecurityV1beta1
	ns   string
}

var networksecurityservertlspoliciesResource = schema.GroupVersionResource{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Resource: "networksecurityservertlspolicies"}

var networksecurityservertlspoliciesKind = schema.GroupVersionKind{Group: "networksecurity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "NetworkSecurityServerTLSPolicy"}

// Get takes name of the networkSecurityServerTLSPolicy, and returns the corresponding networkSecurityServerTLSPolicy object, and an error if there is any.
func (c *FakeNetworkSecurityServerTLSPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkSecurityServerTLSPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networksecurityservertlspoliciesResource, c.ns, name), &v1beta1.NetworkSecurityServerTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityServerTLSPolicy), err
}

// List takes label and field selectors, and returns the list of NetworkSecurityServerTLSPolicies that match those selectors.
func (c *FakeNetworkSecurityServerTLSPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkSecurityServerTLSPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networksecurityservertlspoliciesResource, networksecurityservertlspoliciesKind, c.ns, opts), &v1beta1.NetworkSecurityServerTLSPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkSecurityServerTLSPolicyList{ListMeta: obj.(*v1beta1.NetworkSecurityServerTLSPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkSecurityServerTLSPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkSecurityServerTLSPolicies.
func (c *FakeNetworkSecurityServerTLSPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networksecurityservertlspoliciesResource, c.ns, opts))

}

// Create takes the representation of a networkSecurityServerTLSPolicy and creates it.  Returns the server's representation of the networkSecurityServerTLSPolicy, and an error, if there is any.
func (c *FakeNetworkSecurityServerTLSPolicies) Create(ctx context.Context, networkSecurityServerTLSPolicy *v1beta1.NetworkSecurityServerTLSPolicy, opts v1.CreateOptions) (result *v1beta1.NetworkSecurityServerTLSPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networksecurityservertlspoliciesResource, c.ns, networkSecurityServerTLSPolicy), &v1beta1.NetworkSecurityServerTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityServerTLSPolicy), err
}

// Update takes the representation of a networkSecurityServerTLSPolicy and updates it. Returns the server's representation of the networkSecurityServerTLSPolicy, and an error, if there is any.
func (c *FakeNetworkSecurityServerTLSPolicies) Update(ctx context.Context, networkSecurityServerTLSPolicy *v1beta1.NetworkSecurityServerTLSPolicy, opts v1.UpdateOptions) (result *v1beta1.NetworkSecurityServerTLSPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networksecurityservertlspoliciesResource, c.ns, networkSecurityServerTLSPolicy), &v1beta1.NetworkSecurityServerTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityServerTLSPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkSecurityServerTLSPolicies) UpdateStatus(ctx context.Context, networkSecurityServerTLSPolicy *v1beta1.NetworkSecurityServerTLSPolicy, opts v1.UpdateOptions) (*v1beta1.NetworkSecurityServerTLSPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networksecurityservertlspoliciesResource, "status", c.ns, networkSecurityServerTLSPolicy), &v1beta1.NetworkSecurityServerTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityServerTLSPolicy), err
}

// Delete takes name of the networkSecurityServerTLSPolicy and deletes it. Returns an error if one occurs.
func (c *FakeNetworkSecurityServerTLSPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networksecurityservertlspoliciesResource, c.ns, name), &v1beta1.NetworkSecurityServerTLSPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkSecurityServerTLSPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networksecurityservertlspoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkSecurityServerTLSPolicyList{})
	return err
}

// Patch applies the patch and returns the patched networkSecurityServerTLSPolicy.
func (c *FakeNetworkSecurityServerTLSPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkSecurityServerTLSPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networksecurityservertlspoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkSecurityServerTLSPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkSecurityServerTLSPolicy), err
}
