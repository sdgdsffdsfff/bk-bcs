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

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/apis/bkbcs/v1"
	"github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type BkbcsV1Interface interface {
	RESTClient() rest.Interface
	BcsDbPrivConfigsGetter
	BcsLogConfigsGetter
}

// BkbcsV1Client is used to interact with features provided by the bkbcs group.
type BkbcsV1Client struct {
	restClient rest.Interface
}

func (c *BkbcsV1Client) BcsDbPrivConfigs(namespace string) BcsDbPrivConfigInterface {
	return newBcsDbPrivConfigs(c, namespace)
}

func (c *BkbcsV1Client) BcsLogConfigs(namespace string) BcsLogConfigInterface {
	return newBcsLogConfigs(c, namespace)
}

// NewForConfig creates a new BkbcsV1Client for the given config.
func NewForConfig(c *rest.Config) (*BkbcsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &BkbcsV1Client{client}, nil
}

// NewForConfigOrDie creates a new BkbcsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *BkbcsV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new BkbcsV1Client for the given RESTClient.
func New(c rest.Interface) *BkbcsV1Client {
	return &BkbcsV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *BkbcsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
