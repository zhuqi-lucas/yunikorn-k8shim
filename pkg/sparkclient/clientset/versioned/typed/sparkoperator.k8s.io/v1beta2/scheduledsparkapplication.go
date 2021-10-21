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

package v1beta2

import (
	"context"
	"time"

	v1beta2 "github.com/GoogleCloudPlatform/spark-on-k8s-operator/pkg/apis/sparkoperator.k8s.io/v1beta2"
	scheme "github.com/apache/incubator-yunikorn-k8shim/pkg/sparkclient/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScheduledSparkApplicationsGetter has a method to return a ScheduledSparkApplicationInterface.
// A group's client should implement this interface.
type ScheduledSparkApplicationsGetter interface {
	ScheduledSparkApplications(namespace string) ScheduledSparkApplicationInterface
}

// ScheduledSparkApplicationInterface has methods to work with ScheduledSparkApplication resources.
type ScheduledSparkApplicationInterface interface {
	Create(ctx context.Context, scheduledSparkApplication *v1beta2.ScheduledSparkApplication, opts v1.CreateOptions) (*v1beta2.ScheduledSparkApplication, error)
	Update(ctx context.Context, scheduledSparkApplication *v1beta2.ScheduledSparkApplication, opts v1.UpdateOptions) (*v1beta2.ScheduledSparkApplication, error)
	UpdateStatus(ctx context.Context, scheduledSparkApplication *v1beta2.ScheduledSparkApplication, opts v1.UpdateOptions) (*v1beta2.ScheduledSparkApplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta2.ScheduledSparkApplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta2.ScheduledSparkApplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.ScheduledSparkApplication, err error)
	ScheduledSparkApplicationExpansion
}

// scheduledSparkApplications implements ScheduledSparkApplicationInterface
type scheduledSparkApplications struct {
	client rest.Interface
	ns     string
}

// newScheduledSparkApplications returns a ScheduledSparkApplications
func newScheduledSparkApplications(c *SparkoperatorV1beta2Client, namespace string) *scheduledSparkApplications {
	return &scheduledSparkApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scheduledSparkApplication, and returns the corresponding scheduledSparkApplication object, and an error if there is any.
func (c *scheduledSparkApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.ScheduledSparkApplication, err error) {
	result = &v1beta2.ScheduledSparkApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScheduledSparkApplications that match those selectors.
func (c *scheduledSparkApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.ScheduledSparkApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta2.ScheduledSparkApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scheduledSparkApplications.
func (c *scheduledSparkApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a scheduledSparkApplication and creates it.  Returns the server's representation of the scheduledSparkApplication, and an error, if there is any.
func (c *scheduledSparkApplications) Create(ctx context.Context, scheduledSparkApplication *v1beta2.ScheduledSparkApplication, opts v1.CreateOptions) (result *v1beta2.ScheduledSparkApplication, err error) {
	result = &v1beta2.ScheduledSparkApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scheduledSparkApplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a scheduledSparkApplication and updates it. Returns the server's representation of the scheduledSparkApplication, and an error, if there is any.
func (c *scheduledSparkApplications) Update(ctx context.Context, scheduledSparkApplication *v1beta2.ScheduledSparkApplication, opts v1.UpdateOptions) (result *v1beta2.ScheduledSparkApplication, err error) {
	result = &v1beta2.ScheduledSparkApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Name(scheduledSparkApplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scheduledSparkApplication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *scheduledSparkApplications) UpdateStatus(ctx context.Context, scheduledSparkApplication *v1beta2.ScheduledSparkApplication, opts v1.UpdateOptions) (result *v1beta2.ScheduledSparkApplication, err error) {
	result = &v1beta2.ScheduledSparkApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Name(scheduledSparkApplication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scheduledSparkApplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the scheduledSparkApplication and deletes it. Returns an error if one occurs.
func (c *scheduledSparkApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scheduledSparkApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched scheduledSparkApplication.
func (c *scheduledSparkApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.ScheduledSparkApplication, err error) {
	result = &v1beta2.ScheduledSparkApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scheduledsparkapplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
