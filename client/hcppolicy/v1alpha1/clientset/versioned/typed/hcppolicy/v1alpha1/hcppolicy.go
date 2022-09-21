/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "github.com/KETI-Hybrid/hcp-pkg/apis/hcppolicy/v1alpha1"
	scheme "github.com/KETI-Hybrid/hcp-pkg/client/hcppolicy/v1alpha1/clientset/versioned/scheme"
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HCPPoliciesGetter has a method to return a HCPPolicyInterface.
// A group's client should implement this interface.
type HCPPoliciesGetter interface {
	HCPPolicies(namespace string) HCPPolicyInterface
}

// HCPPolicyInterface has methods to work with HCPPolicy resources.
type HCPPolicyInterface interface {
	Create(ctx context.Context, hCPPolicy *v1alpha1.HCPPolicy, opts v1.CreateOptions) (*v1alpha1.HCPPolicy, error)
	Update(ctx context.Context, hCPPolicy *v1alpha1.HCPPolicy, opts v1.UpdateOptions) (*v1alpha1.HCPPolicy, error)
	UpdateStatus(ctx context.Context, hCPPolicy *v1alpha1.HCPPolicy, opts v1.UpdateOptions) (*v1alpha1.HCPPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HCPPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HCPPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HCPPolicy, err error)
	HCPPolicyExpansion
}

// hCPPolicies implements HCPPolicyInterface
type hCPPolicies struct {
	client rest.Interface
	ns     string
}

// newHCPPolicies returns a HCPPolicies
func newHCPPolicies(c *HcpV1alpha1Client, namespace string) *hCPPolicies {
	return &hCPPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hCPPolicy, and returns the corresponding hCPPolicy object, and an error if there is any.
func (c *hCPPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HCPPolicy, err error) {
	result = &v1alpha1.HCPPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hcppolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HCPPolicies that match those selectors.
func (c *hCPPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HCPPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HCPPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hcppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hCPPolicies.
func (c *hCPPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hcppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hCPPolicy and creates it.  Returns the server's representation of the hCPPolicy, and an error, if there is any.
func (c *hCPPolicies) Create(ctx context.Context, hCPPolicy *v1alpha1.HCPPolicy, opts v1.CreateOptions) (result *v1alpha1.HCPPolicy, err error) {
	result = &v1alpha1.HCPPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hcppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hCPPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hCPPolicy and updates it. Returns the server's representation of the hCPPolicy, and an error, if there is any.
func (c *hCPPolicies) Update(ctx context.Context, hCPPolicy *v1alpha1.HCPPolicy, opts v1.UpdateOptions) (result *v1alpha1.HCPPolicy, err error) {
	result = &v1alpha1.HCPPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hcppolicies").
		Name(hCPPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hCPPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hCPPolicies) UpdateStatus(ctx context.Context, hCPPolicy *v1alpha1.HCPPolicy, opts v1.UpdateOptions) (result *v1alpha1.HCPPolicy, err error) {
	result = &v1alpha1.HCPPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hcppolicies").
		Name(hCPPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hCPPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hCPPolicy and deletes it. Returns an error if one occurs.
func (c *hCPPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hcppolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hCPPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hcppolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hCPPolicy.
func (c *hCPPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HCPPolicy, err error) {
	result = &v1alpha1.HCPPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hcppolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
