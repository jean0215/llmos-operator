/*
Copyright 2024 llmos.ai.

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

	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCephObjectZones implements CephObjectZoneInterface
type FakeCephObjectZones struct {
	Fake *FakeCephV1
	ns   string
}

var cephobjectzonesResource = v1.SchemeGroupVersion.WithResource("cephobjectzones")

var cephobjectzonesKind = v1.SchemeGroupVersion.WithKind("CephObjectZone")

// Get takes name of the cephObjectZone, and returns the corresponding cephObjectZone object, and an error if there is any.
func (c *FakeCephObjectZones) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CephObjectZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cephobjectzonesResource, c.ns, name), &v1.CephObjectZone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephObjectZone), err
}

// List takes label and field selectors, and returns the list of CephObjectZones that match those selectors.
func (c *FakeCephObjectZones) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CephObjectZoneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cephobjectzonesResource, cephobjectzonesKind, c.ns, opts), &v1.CephObjectZoneList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.CephObjectZoneList{ListMeta: obj.(*v1.CephObjectZoneList).ListMeta}
	for _, item := range obj.(*v1.CephObjectZoneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cephObjectZones.
func (c *FakeCephObjectZones) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cephobjectzonesResource, c.ns, opts))

}

// Create takes the representation of a cephObjectZone and creates it.  Returns the server's representation of the cephObjectZone, and an error, if there is any.
func (c *FakeCephObjectZones) Create(ctx context.Context, cephObjectZone *v1.CephObjectZone, opts metav1.CreateOptions) (result *v1.CephObjectZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cephobjectzonesResource, c.ns, cephObjectZone), &v1.CephObjectZone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephObjectZone), err
}

// Update takes the representation of a cephObjectZone and updates it. Returns the server's representation of the cephObjectZone, and an error, if there is any.
func (c *FakeCephObjectZones) Update(ctx context.Context, cephObjectZone *v1.CephObjectZone, opts metav1.UpdateOptions) (result *v1.CephObjectZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cephobjectzonesResource, c.ns, cephObjectZone), &v1.CephObjectZone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephObjectZone), err
}

// Delete takes name of the cephObjectZone and deletes it. Returns an error if one occurs.
func (c *FakeCephObjectZones) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cephobjectzonesResource, c.ns, name, opts), &v1.CephObjectZone{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCephObjectZones) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cephobjectzonesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.CephObjectZoneList{})
	return err
}

// Patch applies the patch and returns the patched cephObjectZone.
func (c *FakeCephObjectZones) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephObjectZone, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cephobjectzonesResource, c.ns, name, pt, data, subresources...), &v1.CephObjectZone{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.CephObjectZone), err
}
