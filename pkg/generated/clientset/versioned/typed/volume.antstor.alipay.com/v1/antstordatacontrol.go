/*
Copyright 2021.

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
	"context"
	"time"

	v1 "lite.io/liteio/pkg/api/volume.antstor.alipay.com/v1"
	scheme "lite.io/liteio/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AntstorDataControlsGetter has a method to return a AntstorDataControlInterface.
// A group's client should implement this interface.
type AntstorDataControlsGetter interface {
	AntstorDataControls(namespace string) AntstorDataControlInterface
}

// AntstorDataControlInterface has methods to work with AntstorDataControl resources.
type AntstorDataControlInterface interface {
	Create(ctx context.Context, antstorDataControl *v1.AntstorDataControl, opts metav1.CreateOptions) (*v1.AntstorDataControl, error)
	Update(ctx context.Context, antstorDataControl *v1.AntstorDataControl, opts metav1.UpdateOptions) (*v1.AntstorDataControl, error)
	UpdateStatus(ctx context.Context, antstorDataControl *v1.AntstorDataControl, opts metav1.UpdateOptions) (*v1.AntstorDataControl, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AntstorDataControl, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AntstorDataControlList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AntstorDataControl, err error)
	AntstorDataControlExpansion
}

// antstorDataControls implements AntstorDataControlInterface
type antstorDataControls struct {
	client rest.Interface
	ns     string
}

// newAntstorDataControls returns a AntstorDataControls
func newAntstorDataControls(c *VolumeV1Client, namespace string) *antstorDataControls {
	return &antstorDataControls{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the antstorDataControl, and returns the corresponding antstorDataControl object, and an error if there is any.
func (c *antstorDataControls) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AntstorDataControl, err error) {
	result = &v1.AntstorDataControl{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("antstordatacontrols").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AntstorDataControls that match those selectors.
func (c *antstorDataControls) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AntstorDataControlList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AntstorDataControlList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("antstordatacontrols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested antstorDataControls.
func (c *antstorDataControls) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("antstordatacontrols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a antstorDataControl and creates it.  Returns the server's representation of the antstorDataControl, and an error, if there is any.
func (c *antstorDataControls) Create(ctx context.Context, antstorDataControl *v1.AntstorDataControl, opts metav1.CreateOptions) (result *v1.AntstorDataControl, err error) {
	result = &v1.AntstorDataControl{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("antstordatacontrols").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(antstorDataControl).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a antstorDataControl and updates it. Returns the server's representation of the antstorDataControl, and an error, if there is any.
func (c *antstorDataControls) Update(ctx context.Context, antstorDataControl *v1.AntstorDataControl, opts metav1.UpdateOptions) (result *v1.AntstorDataControl, err error) {
	result = &v1.AntstorDataControl{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("antstordatacontrols").
		Name(antstorDataControl.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(antstorDataControl).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *antstorDataControls) UpdateStatus(ctx context.Context, antstorDataControl *v1.AntstorDataControl, opts metav1.UpdateOptions) (result *v1.AntstorDataControl, err error) {
	result = &v1.AntstorDataControl{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("antstordatacontrols").
		Name(antstorDataControl.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(antstorDataControl).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the antstorDataControl and deletes it. Returns an error if one occurs.
func (c *antstorDataControls) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("antstordatacontrols").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *antstorDataControls) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("antstordatacontrols").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched antstorDataControl.
func (c *antstorDataControls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AntstorDataControl, err error) {
	result = &v1.AntstorDataControl{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("antstordatacontrols").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
