/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2 "github.com/Tencent/bk-bcs/bcs-mesos/kubebkbcsv2/apis/bkbcs/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBcsEndpoints implements BcsEndpointInterface
type FakeBcsEndpoints struct {
	Fake *FakeBkbcsV2
	ns   string
}

var bcsendpointsResource = schema.GroupVersionResource{Group: "bkbcs", Version: "v2", Resource: "bcsendpoints"}

var bcsendpointsKind = schema.GroupVersionKind{Group: "bkbcs", Version: "v2", Kind: "BcsEndpoint"}

// Get takes name of the bcsEndpoint, and returns the corresponding bcsEndpoint object, and an error if there is any.
func (c *FakeBcsEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.BcsEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bcsendpointsResource, c.ns, name), &v2.BcsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsEndpoint), err
}

// List takes label and field selectors, and returns the list of BcsEndpoints that match those selectors.
func (c *FakeBcsEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v2.BcsEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bcsendpointsResource, bcsendpointsKind, c.ns, opts), &v2.BcsEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.BcsEndpointList{ListMeta: obj.(*v2.BcsEndpointList).ListMeta}
	for _, item := range obj.(*v2.BcsEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bcsEndpoints.
func (c *FakeBcsEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bcsendpointsResource, c.ns, opts))

}

// Create takes the representation of a bcsEndpoint and creates it.  Returns the server's representation of the bcsEndpoint, and an error, if there is any.
func (c *FakeBcsEndpoints) Create(ctx context.Context, bcsEndpoint *v2.BcsEndpoint, opts v1.CreateOptions) (result *v2.BcsEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bcsendpointsResource, c.ns, bcsEndpoint), &v2.BcsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsEndpoint), err
}

// Update takes the representation of a bcsEndpoint and updates it. Returns the server's representation of the bcsEndpoint, and an error, if there is any.
func (c *FakeBcsEndpoints) Update(ctx context.Context, bcsEndpoint *v2.BcsEndpoint, opts v1.UpdateOptions) (result *v2.BcsEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bcsendpointsResource, c.ns, bcsEndpoint), &v2.BcsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBcsEndpoints) UpdateStatus(ctx context.Context, bcsEndpoint *v2.BcsEndpoint, opts v1.UpdateOptions) (*v2.BcsEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bcsendpointsResource, "status", c.ns, bcsEndpoint), &v2.BcsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsEndpoint), err
}

// Delete takes name of the bcsEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeBcsEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bcsendpointsResource, c.ns, name), &v2.BcsEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBcsEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bcsendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2.BcsEndpointList{})
	return err
}

// Patch applies the patch and returns the patched bcsEndpoint.
func (c *FakeBcsEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.BcsEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bcsendpointsResource, c.ns, name, pt, data, subresources...), &v2.BcsEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.BcsEndpoint), err
}
