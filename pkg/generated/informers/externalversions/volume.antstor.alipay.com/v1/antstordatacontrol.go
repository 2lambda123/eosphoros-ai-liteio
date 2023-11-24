/*
Copyright 2021.

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
// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	volumeantstoralipaycomv1 "code.alipay.com/dbplatform/node-disk-controller/pkg/api/volume.antstor.alipay.com/v1"
	versioned "code.alipay.com/dbplatform/node-disk-controller/pkg/generated/clientset/versioned"
	internalinterfaces "code.alipay.com/dbplatform/node-disk-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "code.alipay.com/dbplatform/node-disk-controller/pkg/generated/listers/volume.antstor.alipay.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AntstorDataControlInformer provides access to a shared informer and lister for
// AntstorDataControls.
type AntstorDataControlInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AntstorDataControlLister
}

type antstorDataControlInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAntstorDataControlInformer constructs a new informer for AntstorDataControl type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAntstorDataControlInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAntstorDataControlInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAntstorDataControlInformer constructs a new informer for AntstorDataControl type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAntstorDataControlInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VolumeV1().AntstorDataControls(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VolumeV1().AntstorDataControls(namespace).Watch(context.TODO(), options)
			},
		},
		&volumeantstoralipaycomv1.AntstorDataControl{},
		resyncPeriod,
		indexers,
	)
}

func (f *antstorDataControlInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAntstorDataControlInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *antstorDataControlInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&volumeantstoralipaycomv1.AntstorDataControl{}, f.defaultInformer)
}

func (f *antstorDataControlInformer) Lister() v1.AntstorDataControlLister {
	return v1.NewAntstorDataControlLister(f.Informer().GetIndexer())
}