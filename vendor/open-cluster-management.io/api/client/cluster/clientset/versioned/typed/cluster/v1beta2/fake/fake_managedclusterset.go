// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta2 "open-cluster-management.io/api/cluster/v1beta2"
)

// FakeManagedClusterSets implements ManagedClusterSetInterface
type FakeManagedClusterSets struct {
	Fake *FakeClusterV1beta2
}

var managedclustersetsResource = v1beta2.SchemeGroupVersion.WithResource("managedclustersets")

var managedclustersetsKind = v1beta2.SchemeGroupVersion.WithKind("ManagedClusterSet")

// Get takes name of the managedClusterSet, and returns the corresponding managedClusterSet object, and an error if there is any.
func (c *FakeManagedClusterSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.ManagedClusterSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(managedclustersetsResource, name), &v1beta2.ManagedClusterSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ManagedClusterSet), err
}

// List takes label and field selectors, and returns the list of ManagedClusterSets that match those selectors.
func (c *FakeManagedClusterSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.ManagedClusterSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(managedclustersetsResource, managedclustersetsKind, opts), &v1beta2.ManagedClusterSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.ManagedClusterSetList{ListMeta: obj.(*v1beta2.ManagedClusterSetList).ListMeta}
	for _, item := range obj.(*v1beta2.ManagedClusterSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedClusterSets.
func (c *FakeManagedClusterSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(managedclustersetsResource, opts))
}

// Create takes the representation of a managedClusterSet and creates it.  Returns the server's representation of the managedClusterSet, and an error, if there is any.
func (c *FakeManagedClusterSets) Create(ctx context.Context, managedClusterSet *v1beta2.ManagedClusterSet, opts v1.CreateOptions) (result *v1beta2.ManagedClusterSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(managedclustersetsResource, managedClusterSet), &v1beta2.ManagedClusterSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ManagedClusterSet), err
}

// Update takes the representation of a managedClusterSet and updates it. Returns the server's representation of the managedClusterSet, and an error, if there is any.
func (c *FakeManagedClusterSets) Update(ctx context.Context, managedClusterSet *v1beta2.ManagedClusterSet, opts v1.UpdateOptions) (result *v1beta2.ManagedClusterSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(managedclustersetsResource, managedClusterSet), &v1beta2.ManagedClusterSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ManagedClusterSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedClusterSets) UpdateStatus(ctx context.Context, managedClusterSet *v1beta2.ManagedClusterSet, opts v1.UpdateOptions) (*v1beta2.ManagedClusterSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(managedclustersetsResource, "status", managedClusterSet), &v1beta2.ManagedClusterSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ManagedClusterSet), err
}

// Delete takes name of the managedClusterSet and deletes it. Returns an error if one occurs.
func (c *FakeManagedClusterSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(managedclustersetsResource, name, opts), &v1beta2.ManagedClusterSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedClusterSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(managedclustersetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.ManagedClusterSetList{})
	return err
}

// Patch applies the patch and returns the patched managedClusterSet.
func (c *FakeManagedClusterSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.ManagedClusterSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(managedclustersetsResource, name, pt, data, subresources...), &v1beta2.ManagedClusterSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ManagedClusterSet), err
}