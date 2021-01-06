/*
Copyright 2021 The Kubernetes Authors.

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
	"context"

	v1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceClasses implements ServiceClassInterface
type FakeServiceClasses struct {
	Fake *FakeServicecatalogV1beta1
	ns   string
}

var serviceclassesResource = schema.GroupVersionResource{Group: "servicecatalog.k8s.io", Version: "v1beta1", Resource: "serviceclasses"}

var serviceclassesKind = schema.GroupVersionKind{Group: "servicecatalog.k8s.io", Version: "v1beta1", Kind: "ServiceClass"}

// Get takes name of the serviceClass, and returns the corresponding serviceClass object, and an error if there is any.
func (c *FakeServiceClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ServiceClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceclassesResource, c.ns, name), &v1beta1.ServiceClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}

// List takes label and field selectors, and returns the list of ServiceClasses that match those selectors.
func (c *FakeServiceClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ServiceClassList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceclassesResource, serviceclassesKind, c.ns, opts), &v1beta1.ServiceClassList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ServiceClassList{ListMeta: obj.(*v1beta1.ServiceClassList).ListMeta}
	for _, item := range obj.(*v1beta1.ServiceClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceClasses.
func (c *FakeServiceClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceclassesResource, c.ns, opts))

}

// Create takes the representation of a serviceClass and creates it.  Returns the server's representation of the serviceClass, and an error, if there is any.
func (c *FakeServiceClasses) Create(ctx context.Context, serviceClass *v1beta1.ServiceClass, opts v1.CreateOptions) (result *v1beta1.ServiceClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceclassesResource, c.ns, serviceClass), &v1beta1.ServiceClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}

// Update takes the representation of a serviceClass and updates it. Returns the server's representation of the serviceClass, and an error, if there is any.
func (c *FakeServiceClasses) Update(ctx context.Context, serviceClass *v1beta1.ServiceClass, opts v1.UpdateOptions) (result *v1beta1.ServiceClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceclassesResource, c.ns, serviceClass), &v1beta1.ServiceClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceClasses) UpdateStatus(ctx context.Context, serviceClass *v1beta1.ServiceClass, opts v1.UpdateOptions) (*v1beta1.ServiceClass, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceclassesResource, "status", c.ns, serviceClass), &v1beta1.ServiceClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}

// Delete takes name of the serviceClass and deletes it. Returns an error if one occurs.
func (c *FakeServiceClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceclassesResource, c.ns, name), &v1beta1.ServiceClass{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceclassesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ServiceClassList{})
	return err
}

// Patch applies the patch and returns the patched serviceClass.
func (c *FakeServiceClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ServiceClass, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceclassesResource, c.ns, name, pt, data, subresources...), &v1beta1.ServiceClass{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceClass), err
}
