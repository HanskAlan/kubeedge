/*
Copyright 2020 The KubeEdge Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/kubeedge/kubeedge/cloud/pkg/apis/reliablesyncs/v1alpha1"
	scheme "github.com/kubeedge/kubeedge/cloud/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ObjectSyncsGetter has a method to return a ObjectSyncInterface.
// A group's client should implement this interface.
type ObjectSyncsGetter interface {
	ObjectSyncs(namespace string) ObjectSyncInterface
}

// ObjectSyncInterface has methods to work with ObjectSync resources.
type ObjectSyncInterface interface {
	Create(*v1alpha1.ObjectSync) (*v1alpha1.ObjectSync, error)
	Update(*v1alpha1.ObjectSync) (*v1alpha1.ObjectSync, error)
	UpdateStatus(*v1alpha1.ObjectSync) (*v1alpha1.ObjectSync, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ObjectSync, error)
	List(opts v1.ListOptions) (*v1alpha1.ObjectSyncList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ObjectSync, err error)
	ObjectSyncExpansion
}

// objectSyncs implements ObjectSyncInterface
type objectSyncs struct {
	client rest.Interface
	ns     string
}

// newObjectSyncs returns a ObjectSyncs
func newObjectSyncs(c *ReliablesyncsV1alpha1Client, namespace string) *objectSyncs {
	return &objectSyncs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the objectSync, and returns the corresponding objectSync object, and an error if there is any.
func (c *objectSyncs) Get(name string, options v1.GetOptions) (result *v1alpha1.ObjectSync, err error) {
	result = &v1alpha1.ObjectSync{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("objectsyncs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ObjectSyncs that match those selectors.
func (c *objectSyncs) List(opts v1.ListOptions) (result *v1alpha1.ObjectSyncList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ObjectSyncList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("objectsyncs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested objectSyncs.
func (c *objectSyncs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("objectsyncs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a objectSync and creates it.  Returns the server's representation of the objectSync, and an error, if there is any.
func (c *objectSyncs) Create(objectSync *v1alpha1.ObjectSync) (result *v1alpha1.ObjectSync, err error) {
	result = &v1alpha1.ObjectSync{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("objectsyncs").
		Body(objectSync).
		Do().
		Into(result)
	return
}

// Update takes the representation of a objectSync and updates it. Returns the server's representation of the objectSync, and an error, if there is any.
func (c *objectSyncs) Update(objectSync *v1alpha1.ObjectSync) (result *v1alpha1.ObjectSync, err error) {
	result = &v1alpha1.ObjectSync{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("objectsyncs").
		Name(objectSync.Name).
		Body(objectSync).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *objectSyncs) UpdateStatus(objectSync *v1alpha1.ObjectSync) (result *v1alpha1.ObjectSync, err error) {
	result = &v1alpha1.ObjectSync{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("objectsyncs").
		Name(objectSync.Name).
		SubResource("status").
		Body(objectSync).
		Do().
		Into(result)
	return
}

// Delete takes name of the objectSync and deletes it. Returns an error if one occurs.
func (c *objectSyncs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("objectsyncs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *objectSyncs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("objectsyncs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched objectSync.
func (c *objectSyncs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ObjectSync, err error) {
	result = &v1alpha1.ObjectSync{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("objectsyncs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
