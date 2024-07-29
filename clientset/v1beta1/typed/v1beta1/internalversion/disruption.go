// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2024 Datadog, Inc.
// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	v1beta1 "github.com/DataDog/chaos-controller/api/v1beta1"
	scheme "github.com/DataDog/chaos-controller/clientset/v1beta1/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DisruptionsGetter has a method to return a DisruptionInterface.
// A group's client should implement this interface.
type DisruptionsGetter interface {
	Disruptions(namespace string) DisruptionInterface
}

// DisruptionInterface has methods to work with Disruption resources.
type DisruptionInterface interface {
	Create(ctx context.Context, disruption *v1beta1.Disruption, opts v1.CreateOptions) (*v1beta1.Disruption, error)
	Update(ctx context.Context, disruption *v1beta1.Disruption, opts v1.UpdateOptions) (*v1beta1.Disruption, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Disruption, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DisruptionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	DisruptionExpansion
}

// disruptions implements DisruptionInterface
type disruptions struct {
	client rest.Interface
	ns     string
}

// newDisruptions returns a Disruptions
func newDisruptions(c *ChaosClient, namespace string) *disruptions {
	return &disruptions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the disruption, and returns the corresponding disruption object, and an error if there is any.
func (c *disruptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Disruption, err error) {
	result = &v1beta1.Disruption{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("disruptions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Disruptions that match those selectors.
func (c *disruptions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DisruptionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DisruptionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("disruptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested disruptions.
func (c *disruptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("disruptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a disruption and creates it.  Returns the server's representation of the disruption, and an error, if there is any.
func (c *disruptions) Create(ctx context.Context, disruption *v1beta1.Disruption, opts v1.CreateOptions) (result *v1beta1.Disruption, err error) {
	result = &v1beta1.Disruption{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("disruptions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(disruption).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a disruption and updates it. Returns the server's representation of the disruption, and an error, if there is any.
func (c *disruptions) Update(ctx context.Context, disruption *v1beta1.Disruption, opts v1.UpdateOptions) (result *v1beta1.Disruption, err error) {
	result = &v1beta1.Disruption{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("disruptions").
		Name(disruption.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(disruption).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the disruption and deletes it. Returns an error if one occurs.
func (c *disruptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("disruptions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}