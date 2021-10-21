/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta2 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSparkApplications implements SparkApplicationInterface
type FakeSparkApplications struct {
	Fake *FakeSparkoperatorV1beta2
	ns   string
}

var sparkapplicationsResource = schema.GroupVersionResource{Group: "sparkoperator.k8s.io", Version: "v1beta2", Resource: "sparkapplications"}

var sparkapplicationsKind = schema.GroupVersionKind{Group: "sparkoperator.k8s.io", Version: "v1beta2", Kind: "SparkApplication"}

// Get takes name of the sparkApplication, and returns the corresponding sparkApplication object, and an error if there is any.
func (c *FakeSparkApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.SparkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sparkapplicationsResource, c.ns, name), &v1beta2.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SparkApplication), err
}

// List takes label and field selectors, and returns the list of SparkApplications that match those selectors.
func (c *FakeSparkApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.SparkApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sparkapplicationsResource, sparkapplicationsKind, c.ns, opts), &v1beta2.SparkApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.SparkApplicationList{ListMeta: obj.(*v1beta2.SparkApplicationList).ListMeta}
	for _, item := range obj.(*v1beta2.SparkApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sparkApplications.
func (c *FakeSparkApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sparkapplicationsResource, c.ns, opts))

}

// Create takes the representation of a sparkApplication and creates it.  Returns the server's representation of the sparkApplication, and an error, if there is any.
func (c *FakeSparkApplications) Create(ctx context.Context, sparkApplication *v1beta2.SparkApplication, opts v1.CreateOptions) (result *v1beta2.SparkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sparkapplicationsResource, c.ns, sparkApplication), &v1beta2.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SparkApplication), err
}

// Update takes the representation of a sparkApplication and updates it. Returns the server's representation of the sparkApplication, and an error, if there is any.
func (c *FakeSparkApplications) Update(ctx context.Context, sparkApplication *v1beta2.SparkApplication, opts v1.UpdateOptions) (result *v1beta2.SparkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sparkapplicationsResource, c.ns, sparkApplication), &v1beta2.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SparkApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSparkApplications) UpdateStatus(ctx context.Context, sparkApplication *v1beta2.SparkApplication, opts v1.UpdateOptions) (*v1beta2.SparkApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sparkapplicationsResource, "status", c.ns, sparkApplication), &v1beta2.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SparkApplication), err
}

// Delete takes name of the sparkApplication and deletes it. Returns an error if one occurs.
func (c *FakeSparkApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sparkapplicationsResource, c.ns, name), &v1beta2.SparkApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSparkApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sparkapplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.SparkApplicationList{})
	return err
}

// Patch applies the patch and returns the patched sparkApplication.
func (c *FakeSparkApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.SparkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sparkapplicationsResource, c.ns, name, pt, data, subresources...), &v1beta2.SparkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SparkApplication), err
}
