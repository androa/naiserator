// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/nais/naiserator/pkg/apis/rbac.istio.io/v1alpha1"
	scheme "github.com/nais/naiserator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceRolesGetter has a method to return a ServiceRoleInterface.
// A group's client should implement this interface.
type ServiceRolesGetter interface {
	ServiceRoles(namespace string) ServiceRoleInterface
}

// ServiceRoleInterface has methods to work with ServiceRole resources.
type ServiceRoleInterface interface {
	Create(*v1alpha1.ServiceRole) (*v1alpha1.ServiceRole, error)
	Update(*v1alpha1.ServiceRole) (*v1alpha1.ServiceRole, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ServiceRole, error)
	List(opts v1.ListOptions) (*v1alpha1.ServiceRoleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceRole, err error)
	ServiceRoleExpansion
}

// serviceRoles implements ServiceRoleInterface
type serviceRoles struct {
	client rest.Interface
	ns     string
}

// newServiceRoles returns a ServiceRoles
func newServiceRoles(c *RbacV1alpha1Client, namespace string) *serviceRoles {
	return &serviceRoles{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceRole, and returns the corresponding serviceRole object, and an error if there is any.
func (c *serviceRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.ServiceRole, err error) {
	result = &v1alpha1.ServiceRole{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceRoles that match those selectors.
func (c *serviceRoles) List(opts v1.ListOptions) (result *v1alpha1.ServiceRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ServiceRoleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("serviceroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceRoles.
func (c *serviceRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("serviceroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a serviceRole and creates it.  Returns the server's representation of the serviceRole, and an error, if there is any.
func (c *serviceRoles) Create(serviceRole *v1alpha1.ServiceRole) (result *v1alpha1.ServiceRole, err error) {
	result = &v1alpha1.ServiceRole{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("serviceroles").
		Body(serviceRole).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceRole and updates it. Returns the server's representation of the serviceRole, and an error, if there is any.
func (c *serviceRoles) Update(serviceRole *v1alpha1.ServiceRole) (result *v1alpha1.ServiceRole, err error) {
	result = &v1alpha1.ServiceRole{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("serviceroles").
		Name(serviceRole.Name).
		Body(serviceRole).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceRole and deletes it. Returns an error if one occurs.
func (c *serviceRoles) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceroles").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *serviceRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("serviceroles").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceRole.
func (c *serviceRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ServiceRole, err error) {
	result = &v1alpha1.ServiceRole{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("serviceroles").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
