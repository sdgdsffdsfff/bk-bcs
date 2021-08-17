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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	bkbcsv1 "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/apis/bkbcs/v1"
	versioned "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/generated/clientset/versioned"
	internalinterfaces "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/generated/listers/bkbcs/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BcsDbPrivConfigInformer provides access to a shared informer and lister for
// BcsDbPrivConfigs.
type BcsDbPrivConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.BcsDbPrivConfigLister
}

type bcsDbPrivConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBcsDbPrivConfigInformer constructs a new informer for BcsDbPrivConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBcsDbPrivConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBcsDbPrivConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBcsDbPrivConfigInformer constructs a new informer for BcsDbPrivConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBcsDbPrivConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BkbcsV1().BcsDbPrivConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BkbcsV1().BcsDbPrivConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&bkbcsv1.BcsDbPrivConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *bcsDbPrivConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBcsDbPrivConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bcsDbPrivConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bkbcsv1.BcsDbPrivConfig{}, f.defaultInformer)
}

func (f *bcsDbPrivConfigInformer) Lister() v1.BcsDbPrivConfigLister {
	return v1.NewBcsDbPrivConfigLister(f.Informer().GetIndexer())
}
