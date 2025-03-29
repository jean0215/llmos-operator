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

	v1 "github.com/llmos-ai/llmos-operator/pkg/apis/ml.llmos.ai/v1"
	scheme "github.com/llmos-ai/llmos-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ModelsGetter has a method to return a ModelInterface.
// A group's client should implement this interface.
type ModelsGetter interface {
	Models(namespace string) ModelInterface
}

// ModelInterface has methods to work with Model resources.
type ModelInterface interface {
	Create(ctx context.Context, model *v1.Model, opts metav1.CreateOptions) (*v1.Model, error)
	Update(ctx context.Context, model *v1.Model, opts metav1.UpdateOptions) (*v1.Model, error)
	UpdateStatus(ctx context.Context, model *v1.Model, opts metav1.UpdateOptions) (*v1.Model, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Model, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ModelList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Model, err error)
	ModelExpansion
}

// models implements ModelInterface
type models struct {
	client rest.Interface
	ns     string
}

// newModels returns a Models
func newModels(c *MlV1Client, namespace string) *models {
	return &models{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the model, and returns the corresponding model object, and an error if there is any.
func (c *models) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Model, err error) {
	result = &v1.Model{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("models").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Models that match those selectors.
func (c *models) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ModelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ModelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("models").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested models.
func (c *models) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("models").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a model and creates it.  Returns the server's representation of the model, and an error, if there is any.
func (c *models) Create(ctx context.Context, model *v1.Model, opts metav1.CreateOptions) (result *v1.Model, err error) {
	result = &v1.Model{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("models").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(model).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a model and updates it. Returns the server's representation of the model, and an error, if there is any.
func (c *models) Update(ctx context.Context, model *v1.Model, opts metav1.UpdateOptions) (result *v1.Model, err error) {
	result = &v1.Model{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("models").
		Name(model.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(model).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *models) UpdateStatus(ctx context.Context, model *v1.Model, opts metav1.UpdateOptions) (result *v1.Model, err error) {
	result = &v1.Model{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("models").
		Name(model.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(model).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the model and deletes it. Returns an error if one occurs.
func (c *models) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("models").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *models) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("models").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched model.
func (c *models) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Model, err error) {
	result = &v1.Model{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("models").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
