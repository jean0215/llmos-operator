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

package v1

import (
	"context"
	"time"

	scheme "github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/scheme"
	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CephObjectZoneGroupsGetter has a method to return a CephObjectZoneGroupInterface.
// A group's client should implement this interface.
type CephObjectZoneGroupsGetter interface {
	CephObjectZoneGroups(namespace string) CephObjectZoneGroupInterface
}

// CephObjectZoneGroupInterface has methods to work with CephObjectZoneGroup resources.
type CephObjectZoneGroupInterface interface {
	Create(ctx context.Context, cephObjectZoneGroup *v1.CephObjectZoneGroup, opts metav1.CreateOptions) (*v1.CephObjectZoneGroup, error)
	Update(ctx context.Context, cephObjectZoneGroup *v1.CephObjectZoneGroup, opts metav1.UpdateOptions) (*v1.CephObjectZoneGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CephObjectZoneGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CephObjectZoneGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephObjectZoneGroup, err error)
	CephObjectZoneGroupExpansion
}

// cephObjectZoneGroups implements CephObjectZoneGroupInterface
type cephObjectZoneGroups struct {
	client rest.Interface
	ns     string
}

// newCephObjectZoneGroups returns a CephObjectZoneGroups
func newCephObjectZoneGroups(c *CephV1Client, namespace string) *cephObjectZoneGroups {
	return &cephObjectZoneGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cephObjectZoneGroup, and returns the corresponding cephObjectZoneGroup object, and an error if there is any.
func (c *cephObjectZoneGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.CephObjectZoneGroup, err error) {
	result = &v1.CephObjectZoneGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephobjectzonegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CephObjectZoneGroups that match those selectors.
func (c *cephObjectZoneGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.CephObjectZoneGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.CephObjectZoneGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cephobjectzonegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cephObjectZoneGroups.
func (c *cephObjectZoneGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cephobjectzonegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cephObjectZoneGroup and creates it.  Returns the server's representation of the cephObjectZoneGroup, and an error, if there is any.
func (c *cephObjectZoneGroups) Create(ctx context.Context, cephObjectZoneGroup *v1.CephObjectZoneGroup, opts metav1.CreateOptions) (result *v1.CephObjectZoneGroup, err error) {
	result = &v1.CephObjectZoneGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cephobjectzonegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephObjectZoneGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cephObjectZoneGroup and updates it. Returns the server's representation of the cephObjectZoneGroup, and an error, if there is any.
func (c *cephObjectZoneGroups) Update(ctx context.Context, cephObjectZoneGroup *v1.CephObjectZoneGroup, opts metav1.UpdateOptions) (result *v1.CephObjectZoneGroup, err error) {
	result = &v1.CephObjectZoneGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cephobjectzonegroups").
		Name(cephObjectZoneGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cephObjectZoneGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cephObjectZoneGroup and deletes it. Returns an error if one occurs.
func (c *cephObjectZoneGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephobjectzonegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cephObjectZoneGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cephobjectzonegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cephObjectZoneGroup.
func (c *cephObjectZoneGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephObjectZoneGroup, err error) {
	result = &v1.CephObjectZoneGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cephobjectzonegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
