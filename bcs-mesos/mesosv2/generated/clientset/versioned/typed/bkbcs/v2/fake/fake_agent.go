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

	v2 "github.com/Tencent/bk-bcs/bcs-mesos/mesosv2/apis/bkbcs/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAgents implements AgentInterface
type FakeAgents struct {
	Fake *FakeBkbcsV2
	ns   string
}

var agentsResource = schema.GroupVersionResource{Group: "bkbcs", Version: "v2", Resource: "agents"}

var agentsKind = schema.GroupVersionKind{Group: "bkbcs", Version: "v2", Kind: "Agent"}

// Get takes name of the agent, and returns the corresponding agent object, and an error if there is any.
func (c *FakeAgents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.Agent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(agentsResource, c.ns, name), &v2.Agent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Agent), err
}

// List takes label and field selectors, and returns the list of Agents that match those selectors.
func (c *FakeAgents) List(ctx context.Context, opts v1.ListOptions) (result *v2.AgentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(agentsResource, agentsKind, c.ns, opts), &v2.AgentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.AgentList{ListMeta: obj.(*v2.AgentList).ListMeta}
	for _, item := range obj.(*v2.AgentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested agents.
func (c *FakeAgents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(agentsResource, c.ns, opts))

}

// Create takes the representation of a agent and creates it.  Returns the server's representation of the agent, and an error, if there is any.
func (c *FakeAgents) Create(ctx context.Context, agent *v2.Agent, opts v1.CreateOptions) (result *v2.Agent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(agentsResource, c.ns, agent), &v2.Agent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Agent), err
}

// Update takes the representation of a agent and updates it. Returns the server's representation of the agent, and an error, if there is any.
func (c *FakeAgents) Update(ctx context.Context, agent *v2.Agent, opts v1.UpdateOptions) (result *v2.Agent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(agentsResource, c.ns, agent), &v2.Agent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Agent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAgents) UpdateStatus(ctx context.Context, agent *v2.Agent, opts v1.UpdateOptions) (*v2.Agent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(agentsResource, "status", c.ns, agent), &v2.Agent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Agent), err
}

// Delete takes name of the agent and deletes it. Returns an error if one occurs.
func (c *FakeAgents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(agentsResource, c.ns, name), &v2.Agent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAgents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(agentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v2.AgentList{})
	return err
}

// Patch applies the patch and returns the patched agent.
func (c *FakeAgents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.Agent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(agentsResource, c.ns, name, pt, data, subresources...), &v2.Agent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.Agent), err
}
