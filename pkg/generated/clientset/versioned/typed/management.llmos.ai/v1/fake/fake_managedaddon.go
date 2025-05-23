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

	v1 "github.com/llmos-ai/llmos-operator/pkg/apis/management.llmos.ai/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagedAddons implements ManagedAddonInterface
type FakeManagedAddons struct {
	Fake *FakeManagementV1
	ns   string
}

var managedaddonsResource = v1.SchemeGroupVersion.WithResource("managedaddons")

var managedaddonsKind = v1.SchemeGroupVersion.WithKind("ManagedAddon")

// Get takes name of the managedAddon, and returns the corresponding managedAddon object, and an error if there is any.
func (c *FakeManagedAddons) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ManagedAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedaddonsResource, c.ns, name), &v1.ManagedAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagedAddon), err
}

// List takes label and field selectors, and returns the list of ManagedAddons that match those selectors.
func (c *FakeManagedAddons) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ManagedAddonList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedaddonsResource, managedaddonsKind, c.ns, opts), &v1.ManagedAddonList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ManagedAddonList{ListMeta: obj.(*v1.ManagedAddonList).ListMeta}
	for _, item := range obj.(*v1.ManagedAddonList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedAddons.
func (c *FakeManagedAddons) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedaddonsResource, c.ns, opts))

}

// Create takes the representation of a managedAddon and creates it.  Returns the server's representation of the managedAddon, and an error, if there is any.
func (c *FakeManagedAddons) Create(ctx context.Context, managedAddon *v1.ManagedAddon, opts metav1.CreateOptions) (result *v1.ManagedAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedaddonsResource, c.ns, managedAddon), &v1.ManagedAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagedAddon), err
}

// Update takes the representation of a managedAddon and updates it. Returns the server's representation of the managedAddon, and an error, if there is any.
func (c *FakeManagedAddons) Update(ctx context.Context, managedAddon *v1.ManagedAddon, opts metav1.UpdateOptions) (result *v1.ManagedAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedaddonsResource, c.ns, managedAddon), &v1.ManagedAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagedAddon), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedAddons) UpdateStatus(ctx context.Context, managedAddon *v1.ManagedAddon, opts metav1.UpdateOptions) (*v1.ManagedAddon, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedaddonsResource, "status", c.ns, managedAddon), &v1.ManagedAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagedAddon), err
}

// Delete takes name of the managedAddon and deletes it. Returns an error if one occurs.
func (c *FakeManagedAddons) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(managedaddonsResource, c.ns, name, opts), &v1.ManagedAddon{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedAddons) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedaddonsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ManagedAddonList{})
	return err
}

// Patch applies the patch and returns the patched managedAddon.
func (c *FakeManagedAddons) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ManagedAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedaddonsResource, c.ns, name, pt, data, subresources...), &v1.ManagedAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ManagedAddon), err
}
