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
	"context"
	"time"

	v1alpha1 "github.com/hdkshingala/deploymentcreator/pkg/apis/hardik.dev/v1alpha1"
	scheme "github.com/hdkshingala/deploymentcreator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeploymentCreatorsGetter has a method to return a DeploymentCreatorInterface.
// A group's client should implement this interface.
type DeploymentCreatorsGetter interface {
	DeploymentCreators(namespace string) DeploymentCreatorInterface
}

// DeploymentCreatorInterface has methods to work with DeploymentCreator resources.
type DeploymentCreatorInterface interface {
	Create(ctx context.Context, deploymentCreator *v1alpha1.DeploymentCreator, opts v1.CreateOptions) (*v1alpha1.DeploymentCreator, error)
	Update(ctx context.Context, deploymentCreator *v1alpha1.DeploymentCreator, opts v1.UpdateOptions) (*v1alpha1.DeploymentCreator, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DeploymentCreator, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DeploymentCreatorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeploymentCreator, err error)
	DeploymentCreatorExpansion
}

// deploymentCreators implements DeploymentCreatorInterface
type deploymentCreators struct {
	client rest.Interface
	ns     string
}

// newDeploymentCreators returns a DeploymentCreators
func newDeploymentCreators(c *HardikV1alpha1Client, namespace string) *deploymentCreators {
	return &deploymentCreators{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deploymentCreator, and returns the corresponding deploymentCreator object, and an error if there is any.
func (c *deploymentCreators) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeploymentCreator, err error) {
	result = &v1alpha1.DeploymentCreator{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentcreators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeploymentCreators that match those selectors.
func (c *deploymentCreators) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeploymentCreatorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DeploymentCreatorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymentcreators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deploymentCreators.
func (c *deploymentCreators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deploymentcreators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deploymentCreator and creates it.  Returns the server's representation of the deploymentCreator, and an error, if there is any.
func (c *deploymentCreators) Create(ctx context.Context, deploymentCreator *v1alpha1.DeploymentCreator, opts v1.CreateOptions) (result *v1alpha1.DeploymentCreator, err error) {
	result = &v1alpha1.DeploymentCreator{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deploymentcreators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deploymentCreator).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deploymentCreator and updates it. Returns the server's representation of the deploymentCreator, and an error, if there is any.
func (c *deploymentCreators) Update(ctx context.Context, deploymentCreator *v1alpha1.DeploymentCreator, opts v1.UpdateOptions) (result *v1alpha1.DeploymentCreator, err error) {
	result = &v1alpha1.DeploymentCreator{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymentcreators").
		Name(deploymentCreator.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deploymentCreator).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deploymentCreator and deletes it. Returns an error if one occurs.
func (c *deploymentCreators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentcreators").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deploymentCreators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymentcreators").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deploymentCreator.
func (c *deploymentCreators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeploymentCreator, err error) {
	result = &v1alpha1.DeploymentCreator{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deploymentcreators").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}