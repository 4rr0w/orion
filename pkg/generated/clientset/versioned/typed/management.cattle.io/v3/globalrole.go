/*
Copyright 2024 Rancher Labs, Inc.

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

package v3

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GlobalRolesGetter has a method to return a GlobalRoleInterface.
// A group's client should implement this interface.
type GlobalRolesGetter interface {
	GlobalRoles() GlobalRoleInterface
}

// GlobalRoleInterface has methods to work with GlobalRole resources.
type GlobalRoleInterface interface {
	Create(ctx context.Context, globalRole *v3.GlobalRole, opts v1.CreateOptions) (*v3.GlobalRole, error)
	Update(ctx context.Context, globalRole *v3.GlobalRole, opts v1.UpdateOptions) (*v3.GlobalRole, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.GlobalRole, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.GlobalRoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalRole, err error)
	GlobalRoleExpansion
}

// globalRoles implements GlobalRoleInterface
type globalRoles struct {
	client rest.Interface
}

// newGlobalRoles returns a GlobalRoles
func newGlobalRoles(c *ManagementV3Client) *globalRoles {
	return &globalRoles{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalRole, and returns the corresponding globalRole object, and an error if there is any.
func (c *globalRoles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.GlobalRole, err error) {
	result = &v3.GlobalRole{}
	err = c.client.Get().
		Resource("globalroles").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalRoles that match those selectors.
func (c *globalRoles) List(ctx context.Context, opts v1.ListOptions) (result *v3.GlobalRoleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.GlobalRoleList{}
	err = c.client.Get().
		Resource("globalroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalRoles.
func (c *globalRoles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a globalRole and creates it.  Returns the server's representation of the globalRole, and an error, if there is any.
func (c *globalRoles) Create(ctx context.Context, globalRole *v3.GlobalRole, opts v1.CreateOptions) (result *v3.GlobalRole, err error) {
	result = &v3.GlobalRole{}
	err = c.client.Post().
		Resource("globalroles").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalRole).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a globalRole and updates it. Returns the server's representation of the globalRole, and an error, if there is any.
func (c *globalRoles) Update(ctx context.Context, globalRole *v3.GlobalRole, opts v1.UpdateOptions) (result *v3.GlobalRole, err error) {
	result = &v3.GlobalRole{}
	err = c.client.Put().
		Resource("globalroles").
		Name(globalRole.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(globalRole).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the globalRole and deletes it. Returns an error if one occurs.
func (c *globalRoles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalroles").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalRoles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalroles").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched globalRole.
func (c *globalRoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.GlobalRole, err error) {
	result = &v3.GlobalRole{}
	err = c.client.Patch(pt).
		Resource("globalroles").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
