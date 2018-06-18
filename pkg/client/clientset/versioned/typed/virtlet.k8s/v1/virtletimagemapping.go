/*
Copyright 2018 Mirantis

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

package v1

import (
	v1 "github.com/Mirantis/virtlet/pkg/api/virtlet.k8s/v1"
	scheme "github.com/Mirantis/virtlet/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtletImageMappingsGetter has a method to return a VirtletImageMappingInterface.
// A group's client should implement this interface.
type VirtletImageMappingsGetter interface {
	VirtletImageMappings(namespace string) VirtletImageMappingInterface
}

// VirtletImageMappingInterface has methods to work with VirtletImageMapping resources.
type VirtletImageMappingInterface interface {
	Create(*v1.VirtletImageMapping) (*v1.VirtletImageMapping, error)
	Update(*v1.VirtletImageMapping) (*v1.VirtletImageMapping, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.VirtletImageMapping, error)
	List(opts meta_v1.ListOptions) (*v1.VirtletImageMappingList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VirtletImageMapping, err error)
	VirtletImageMappingExpansion
}

// virtletImageMappings implements VirtletImageMappingInterface
type virtletImageMappings struct {
	client rest.Interface
	ns     string
}

// newVirtletImageMappings returns a VirtletImageMappings
func newVirtletImageMappings(c *VirtletV1Client, namespace string) *virtletImageMappings {
	return &virtletImageMappings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtletImageMapping, and returns the corresponding virtletImageMapping object, and an error if there is any.
func (c *virtletImageMappings) Get(name string, options meta_v1.GetOptions) (result *v1.VirtletImageMapping, err error) {
	result = &v1.VirtletImageMapping{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtletimagemappings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtletImageMappings that match those selectors.
func (c *virtletImageMappings) List(opts meta_v1.ListOptions) (result *v1.VirtletImageMappingList, err error) {
	result = &v1.VirtletImageMappingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtletimagemappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtletImageMappings.
func (c *virtletImageMappings) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtletimagemappings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a virtletImageMapping and creates it.  Returns the server's representation of the virtletImageMapping, and an error, if there is any.
func (c *virtletImageMappings) Create(virtletImageMapping *v1.VirtletImageMapping) (result *v1.VirtletImageMapping, err error) {
	result = &v1.VirtletImageMapping{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtletimagemappings").
		Body(virtletImageMapping).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtletImageMapping and updates it. Returns the server's representation of the virtletImageMapping, and an error, if there is any.
func (c *virtletImageMappings) Update(virtletImageMapping *v1.VirtletImageMapping) (result *v1.VirtletImageMapping, err error) {
	result = &v1.VirtletImageMapping{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtletimagemappings").
		Name(virtletImageMapping.Name).
		Body(virtletImageMapping).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtletImageMapping and deletes it. Returns an error if one occurs.
func (c *virtletImageMappings) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtletimagemappings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtletImageMappings) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtletimagemappings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtletImageMapping.
func (c *virtletImageMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.VirtletImageMapping, err error) {
	result = &v1.VirtletImageMapping{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtletimagemappings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}