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

	volumeantstoralipaycomv1 "lite.io/liteio/pkg/api/volume.antstor.alipay.com/v1"
	versioned "lite.io/liteio/pkg/generated/clientset/versioned"
	internalinterfaces "lite.io/liteio/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "lite.io/liteio/pkg/generated/listers/volume.antstor.alipay.com/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VolumeMigrationInformer provides access to a shared informer and lister for
// VolumeMigrations.
type VolumeMigrationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.VolumeMigrationLister
}

type volumeMigrationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVolumeMigrationInformer constructs a new informer for VolumeMigration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVolumeMigrationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVolumeMigrationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVolumeMigrationInformer constructs a new informer for VolumeMigration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVolumeMigrationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VolumeV1().VolumeMigrations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VolumeV1().VolumeMigrations(namespace).Watch(context.TODO(), options)
			},
		},
		&volumeantstoralipaycomv1.VolumeMigration{},
		resyncPeriod,
		indexers,
	)
}

func (f *volumeMigrationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVolumeMigrationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *volumeMigrationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&volumeantstoralipaycomv1.VolumeMigration{}, f.defaultInformer)
}

func (f *volumeMigrationInformer) Lister() v1.VolumeMigrationLister {
	return v1.NewVolumeMigrationLister(f.Informer().GetIndexer())
}
