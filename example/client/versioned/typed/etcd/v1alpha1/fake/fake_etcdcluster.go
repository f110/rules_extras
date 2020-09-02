/*

MIT License

Copyright (c) 2019 Fumihiro Ito

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.f110.dev/rules_extras/example/api/etcd/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEtcdClusters implements EtcdClusterInterface
type FakeEtcdClusters struct {
	Fake *FakeEtcdV1alpha1
	ns   string
}

var etcdclustersResource = schema.GroupVersionResource{Group: "etcd.f110.dev", Version: "v1alpha1", Resource: "etcdclusters"}

var etcdclustersKind = schema.GroupVersionKind{Group: "etcd.f110.dev", Version: "v1alpha1", Kind: "EtcdCluster"}

// Get takes name of the etcdCluster, and returns the corresponding etcdCluster object, and an error if there is any.
func (c *FakeEtcdClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EtcdCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(etcdclustersResource, c.ns, name), &v1alpha1.EtcdCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdCluster), err
}

// List takes label and field selectors, and returns the list of EtcdClusters that match those selectors.
func (c *FakeEtcdClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EtcdClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(etcdclustersResource, etcdclustersKind, c.ns, opts), &v1alpha1.EtcdClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EtcdClusterList{ListMeta: obj.(*v1alpha1.EtcdClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.EtcdClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested etcdClusters.
func (c *FakeEtcdClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(etcdclustersResource, c.ns, opts))

}

// Create takes the representation of a etcdCluster and creates it.  Returns the server's representation of the etcdCluster, and an error, if there is any.
func (c *FakeEtcdClusters) Create(ctx context.Context, etcdCluster *v1alpha1.EtcdCluster, opts v1.CreateOptions) (result *v1alpha1.EtcdCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(etcdclustersResource, c.ns, etcdCluster), &v1alpha1.EtcdCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdCluster), err
}

// Update takes the representation of a etcdCluster and updates it. Returns the server's representation of the etcdCluster, and an error, if there is any.
func (c *FakeEtcdClusters) Update(ctx context.Context, etcdCluster *v1alpha1.EtcdCluster, opts v1.UpdateOptions) (result *v1alpha1.EtcdCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(etcdclustersResource, c.ns, etcdCluster), &v1alpha1.EtcdCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEtcdClusters) UpdateStatus(ctx context.Context, etcdCluster *v1alpha1.EtcdCluster, opts v1.UpdateOptions) (*v1alpha1.EtcdCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(etcdclustersResource, "status", c.ns, etcdCluster), &v1alpha1.EtcdCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdCluster), err
}

// Delete takes name of the etcdCluster and deletes it. Returns an error if one occurs.
func (c *FakeEtcdClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(etcdclustersResource, c.ns, name), &v1alpha1.EtcdCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEtcdClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(etcdclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EtcdClusterList{})
	return err
}

// Patch applies the patch and returns the patched etcdCluster.
func (c *FakeEtcdClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EtcdCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(etcdclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.EtcdCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdCluster), err
}
