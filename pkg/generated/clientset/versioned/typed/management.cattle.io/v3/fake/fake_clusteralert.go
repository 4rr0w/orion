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

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterAlerts implements ClusterAlertInterface
type FakeClusterAlerts struct {
	Fake *FakeManagementV3
	ns   string
}

var clusteralertsResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "clusteralerts"}

var clusteralertsKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterAlert"}

// Get takes name of the clusterAlert, and returns the corresponding clusterAlert object, and an error if there is any.
func (c *FakeClusterAlerts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusteralertsResource, c.ns, name), &v3.ClusterAlert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlert), err
}

// List takes label and field selectors, and returns the list of ClusterAlerts that match those selectors.
func (c *FakeClusterAlerts) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterAlertList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusteralertsResource, clusteralertsKind, c.ns, opts), &v3.ClusterAlertList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ClusterAlertList{ListMeta: obj.(*v3.ClusterAlertList).ListMeta}
	for _, item := range obj.(*v3.ClusterAlertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterAlerts.
func (c *FakeClusterAlerts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusteralertsResource, c.ns, opts))

}

// Create takes the representation of a clusterAlert and creates it.  Returns the server's representation of the clusterAlert, and an error, if there is any.
func (c *FakeClusterAlerts) Create(ctx context.Context, clusterAlert *v3.ClusterAlert, opts v1.CreateOptions) (result *v3.ClusterAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusteralertsResource, c.ns, clusterAlert), &v3.ClusterAlert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlert), err
}

// Update takes the representation of a clusterAlert and updates it. Returns the server's representation of the clusterAlert, and an error, if there is any.
func (c *FakeClusterAlerts) Update(ctx context.Context, clusterAlert *v3.ClusterAlert, opts v1.UpdateOptions) (result *v3.ClusterAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusteralertsResource, c.ns, clusterAlert), &v3.ClusterAlert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlert), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterAlerts) UpdateStatus(ctx context.Context, clusterAlert *v3.ClusterAlert, opts v1.UpdateOptions) (*v3.ClusterAlert, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusteralertsResource, "status", c.ns, clusterAlert), &v3.ClusterAlert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlert), err
}

// Delete takes name of the clusterAlert and deletes it. Returns an error if one occurs.
func (c *FakeClusterAlerts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusteralertsResource, c.ns, name, opts), &v3.ClusterAlert{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterAlerts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusteralertsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ClusterAlertList{})
	return err
}

// Patch applies the patch and returns the patched clusterAlert.
func (c *FakeClusterAlerts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterAlert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusteralertsResource, c.ns, name, pt, data, subresources...), &v3.ClusterAlert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterAlert), err
}
