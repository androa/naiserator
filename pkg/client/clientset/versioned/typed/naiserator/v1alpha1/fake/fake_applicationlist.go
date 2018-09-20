/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/nais/naiserator/pkg/apis/naiserator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationLists implements ApplicationListInterface
type FakeApplicationLists struct {
	Fake *FakeNaiseratorV1alpha1
	ns   string
}

var applicationlistsResource = schema.GroupVersionResource{Group: "naiserator.nais.io", Version: "v1alpha1", Resource: "applicationlists"}

var applicationlistsKind = schema.GroupVersionKind{Group: "naiserator.nais.io", Version: "v1alpha1", Kind: "ApplicationList"}

// Get takes name of the applicationList, and returns the corresponding applicationList object, and an error if there is any.
func (c *FakeApplicationLists) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationlistsResource, c.ns, name), &v1alpha1.ApplicationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationList), err
}

// List takes label and field selectors, and returns the list of ApplicationLists that match those selectors.
func (c *FakeApplicationLists) List(opts v1.ListOptions) (result *v1alpha1.ApplicationListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationlistsResource, applicationlistsKind, c.ns, opts), &v1alpha1.ApplicationListList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationListList), err
}

// Watch returns a watch.Interface that watches the requested applicationLists.
func (c *FakeApplicationLists) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationlistsResource, c.ns, opts))

}

// Create takes the representation of a applicationList and creates it.  Returns the server's representation of the applicationList, and an error, if there is any.
func (c *FakeApplicationLists) Create(applicationList *v1alpha1.ApplicationList) (result *v1alpha1.ApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationlistsResource, c.ns, applicationList), &v1alpha1.ApplicationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationList), err
}

// Update takes the representation of a applicationList and updates it. Returns the server's representation of the applicationList, and an error, if there is any.
func (c *FakeApplicationLists) Update(applicationList *v1alpha1.ApplicationList) (result *v1alpha1.ApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationlistsResource, c.ns, applicationList), &v1alpha1.ApplicationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationList), err
}

// Delete takes name of the applicationList and deletes it. Returns an error if one occurs.
func (c *FakeApplicationLists) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationlistsResource, c.ns, name), &v1alpha1.ApplicationList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationLists) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationlistsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationListList{})
	return err
}

// Patch applies the patch and returns the patched applicationList.
func (c *FakeApplicationLists) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationlistsResource, c.ns, name, data, subresources...), &v1alpha1.ApplicationList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationList), err
}