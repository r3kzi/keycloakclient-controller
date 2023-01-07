/*
Copyright 2022.

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

	v1alpha1 "github.com/christianwoehrle/keycloakclient-controller/api/v1alpha1"
	scheme "github.com/christianwoehrle/keycloakclient-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KeycloakClientsGetter has a method to return a KeycloakClientInterface.
// A group's client should implement this interface.
type KeycloakClientsGetter interface {
	KeycloakClients(namespace string) KeycloakClientInterface
}

// KeycloakClientInterface has methods to work with KeycloakClient resources.
type KeycloakClientInterface interface {
	Create(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.CreateOptions) (*v1alpha1.KeycloakClient, error)
	Update(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.UpdateOptions) (*v1alpha1.KeycloakClient, error)
	UpdateStatus(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.UpdateOptions) (*v1alpha1.KeycloakClient, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.KeycloakClient, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.KeycloakClientList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeycloakClient, err error)
	KeycloakClientExpansion
}

// keycloakClients implements KeycloakClientInterface
type keycloakClients struct {
	client rest.Interface
	ns     string
}

// newKeycloakClients returns a KeycloakClients
func newKeycloakClients(c *KeycloakV1alpha1Client, namespace string) *keycloakClients {
	return &keycloakClients{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the keycloakClient, and returns the corresponding keycloakClient object, and an error if there is any.
func (c *keycloakClients) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KeycloakClient, err error) {
	result = &v1alpha1.KeycloakClient{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keycloakclients").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KeycloakClients that match those selectors.
func (c *keycloakClients) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeycloakClientList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KeycloakClientList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keycloakclients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested keycloakClients.
func (c *keycloakClients) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("keycloakclients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a keycloakClient and creates it.  Returns the server's representation of the keycloakClient, and an error, if there is any.
func (c *keycloakClients) Create(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.CreateOptions) (result *v1alpha1.KeycloakClient, err error) {
	result = &v1alpha1.KeycloakClient{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("keycloakclients").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keycloakClient).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a keycloakClient and updates it. Returns the server's representation of the keycloakClient, and an error, if there is any.
func (c *keycloakClients) Update(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.UpdateOptions) (result *v1alpha1.KeycloakClient, err error) {
	result = &v1alpha1.KeycloakClient{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keycloakclients").
		Name(keycloakClient.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keycloakClient).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *keycloakClients) UpdateStatus(ctx context.Context, keycloakClient *v1alpha1.KeycloakClient, opts v1.UpdateOptions) (result *v1alpha1.KeycloakClient, err error) {
	result = &v1alpha1.KeycloakClient{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keycloakclients").
		Name(keycloakClient.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(keycloakClient).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the keycloakClient and deletes it. Returns an error if one occurs.
func (c *keycloakClients) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keycloakclients").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *keycloakClients) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keycloakclients").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched keycloakClient.
func (c *keycloakClients) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeycloakClient, err error) {
	result = &v1alpha1.KeycloakClient{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("keycloakclients").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
