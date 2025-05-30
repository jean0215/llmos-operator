/*
Copyright 2025 llmos.ai.

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
// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/llmos-ai/llmos-operator/pkg/apis/ml.llmos.ai/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRegistries implements RegistryInterface
type FakeRegistries struct {
	Fake *FakeMlV1
}

var registriesResource = v1.SchemeGroupVersion.WithResource("registries")

var registriesKind = v1.SchemeGroupVersion.WithKind("Registry")

// Get takes name of the registry, and returns the corresponding registry object, and an error if there is any.
func (c *FakeRegistries) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(registriesResource, name), &v1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Registry), err
}

// List takes label and field selectors, and returns the list of Registries that match those selectors.
func (c *FakeRegistries) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(registriesResource, registriesKind, opts), &v1.RegistryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RegistryList{ListMeta: obj.(*v1.RegistryList).ListMeta}
	for _, item := range obj.(*v1.RegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested registries.
func (c *FakeRegistries) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(registriesResource, opts))
}

// Create takes the representation of a registry and creates it.  Returns the server's representation of the registry, and an error, if there is any.
func (c *FakeRegistries) Create(ctx context.Context, registry *v1.Registry, opts metav1.CreateOptions) (result *v1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(registriesResource, registry), &v1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Registry), err
}

// Update takes the representation of a registry and updates it. Returns the server's representation of the registry, and an error, if there is any.
func (c *FakeRegistries) Update(ctx context.Context, registry *v1.Registry, opts metav1.UpdateOptions) (result *v1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(registriesResource, registry), &v1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Registry), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRegistries) UpdateStatus(ctx context.Context, registry *v1.Registry, opts metav1.UpdateOptions) (*v1.Registry, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(registriesResource, "status", registry), &v1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Registry), err
}

// Delete takes name of the registry and deletes it. Returns an error if one occurs.
func (c *FakeRegistries) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(registriesResource, name, opts), &v1.Registry{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRegistries) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(registriesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RegistryList{})
	return err
}

// Patch applies the patch and returns the patched registry.
func (c *FakeRegistries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Registry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(registriesResource, name, pt, data, subresources...), &v1.Registry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Registry), err
}
