// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HTTPAPISpecBindingsGetter has a method to return a HTTPAPISpecBindingInterface.
// A group's client should implement this interface.
type HTTPAPISpecBindingsGetter interface {
	HTTPAPISpecBindings(namespace string) HTTPAPISpecBindingInterface
}

// HTTPAPISpecBindingInterface has methods to work with HTTPAPISpecBinding resources.
type HTTPAPISpecBindingInterface interface {
	Create(ctx context.Context, hTTPAPISpecBinding *v1alpha2.HTTPAPISpecBinding, opts v1.CreateOptions) (*v1alpha2.HTTPAPISpecBinding, error)
	Update(ctx context.Context, hTTPAPISpecBinding *v1alpha2.HTTPAPISpecBinding, opts v1.UpdateOptions) (*v1alpha2.HTTPAPISpecBinding, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.HTTPAPISpecBinding, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.HTTPAPISpecBindingList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.HTTPAPISpecBinding, err error)
	HTTPAPISpecBindingExpansion
}

// hTTPAPISpecBindings implements HTTPAPISpecBindingInterface
type hTTPAPISpecBindings struct {
	client rest.Interface
	ns     string
}

// newHTTPAPISpecBindings returns a HTTPAPISpecBindings
func newHTTPAPISpecBindings(c *ConfigV1alpha2Client, namespace string) *hTTPAPISpecBindings {
	return &hTTPAPISpecBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hTTPAPISpecBinding, and returns the corresponding hTTPAPISpecBinding object, and an error if there is any.
func (c *hTTPAPISpecBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	result = &v1alpha2.HTTPAPISpecBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HTTPAPISpecBindings that match those selectors.
func (c *hTTPAPISpecBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.HTTPAPISpecBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.HTTPAPISpecBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hTTPAPISpecBindings.
func (c *hTTPAPISpecBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hTTPAPISpecBinding and creates it.  Returns the server's representation of the hTTPAPISpecBinding, and an error, if there is any.
func (c *hTTPAPISpecBindings) Create(ctx context.Context, hTTPAPISpecBinding *v1alpha2.HTTPAPISpecBinding, opts v1.CreateOptions) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	result = &v1alpha2.HTTPAPISpecBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPAPISpecBinding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hTTPAPISpecBinding and updates it. Returns the server's representation of the hTTPAPISpecBinding, and an error, if there is any.
func (c *hTTPAPISpecBindings) Update(ctx context.Context, hTTPAPISpecBinding *v1alpha2.HTTPAPISpecBinding, opts v1.UpdateOptions) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	result = &v1alpha2.HTTPAPISpecBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		Name(hTTPAPISpecBinding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hTTPAPISpecBinding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hTTPAPISpecBinding and deletes it. Returns an error if one occurs.
func (c *hTTPAPISpecBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hTTPAPISpecBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpapispecbindings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hTTPAPISpecBinding.
func (c *hTTPAPISpecBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.HTTPAPISpecBinding, err error) {
	result = &v1alpha2.HTTPAPISpecBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("httpapispecbindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
