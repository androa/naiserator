// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	naiseratorv1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/nais.io/v1alpha1"
	networkingv1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/networking.istio.io/v1alpha3"
	rbacv1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/rbac.istio.io/v1alpha1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	NaiseratorV1alpha1() naiseratorv1alpha1.NaiseratorV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Naiserator() naiseratorv1alpha1.NaiseratorV1alpha1Interface
	NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface
	// Deprecated: please explicitly pick a version if possible.
	Networking() networkingv1alpha3.NetworkingV1alpha3Interface
	RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Rbac() rbacv1alpha1.RbacV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	naiseratorV1alpha1 *naiseratorv1alpha1.NaiseratorV1alpha1Client
	networkingV1alpha3 *networkingv1alpha3.NetworkingV1alpha3Client
	rbacV1alpha1       *rbacv1alpha1.RbacV1alpha1Client
}

// NaiseratorV1alpha1 retrieves the NaiseratorV1alpha1Client
func (c *Clientset) NaiseratorV1alpha1() naiseratorv1alpha1.NaiseratorV1alpha1Interface {
	return c.naiseratorV1alpha1
}

// Deprecated: Naiserator retrieves the default version of NaiseratorClient.
// Please explicitly pick a version.
func (c *Clientset) Naiserator() naiseratorv1alpha1.NaiseratorV1alpha1Interface {
	return c.naiseratorV1alpha1
}

// NetworkingV1alpha3 retrieves the NetworkingV1alpha3Client
func (c *Clientset) NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface {
	return c.networkingV1alpha3
}

// Deprecated: Networking retrieves the default version of NetworkingClient.
// Please explicitly pick a version.
func (c *Clientset) Networking() networkingv1alpha3.NetworkingV1alpha3Interface {
	return c.networkingV1alpha3
}

// RbacV1alpha1 retrieves the RbacV1alpha1Client
func (c *Clientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	return c.rbacV1alpha1
}

// Deprecated: Rbac retrieves the default version of RbacClient.
// Please explicitly pick a version.
func (c *Clientset) Rbac() rbacv1alpha1.RbacV1alpha1Interface {
	return c.rbacV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.naiseratorV1alpha1, err = naiseratorv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkingV1alpha3, err = networkingv1alpha3.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.rbacV1alpha1, err = rbacv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.naiseratorV1alpha1 = naiseratorv1alpha1.NewForConfigOrDie(c)
	cs.networkingV1alpha3 = networkingv1alpha3.NewForConfigOrDie(c)
	cs.rbacV1alpha1 = rbacv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.naiseratorV1alpha1 = naiseratorv1alpha1.New(c)
	cs.networkingV1alpha3 = networkingv1alpha3.New(c)
	cs.rbacV1alpha1 = rbacv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
