// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	rbacistioiov1alpha1 "github.com/nais/naiserator/pkg/apis/rbac.istio.io/v1alpha1"
	versioned "github.com/nais/naiserator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nais/naiserator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nais/naiserator/pkg/client/listers/rbac.istio.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceRoleInformer provides access to a shared informer and lister for
// ServiceRoles.
type ServiceRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ServiceRoleLister
}

type serviceRoleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceRoleInformer constructs a new informer for ServiceRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceRoleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceRoleInformer constructs a new informer for ServiceRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().ServiceRoles(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().ServiceRoles(namespace).Watch(options)
			},
		},
		&rbacistioiov1alpha1.ServiceRole{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceRoleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceRoleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceRoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbacistioiov1alpha1.ServiceRole{}, f.defaultInformer)
}

func (f *serviceRoleInformer) Lister() v1alpha1.ServiceRoleLister {
	return v1alpha1.NewServiceRoleLister(f.Informer().GetIndexer())
}
