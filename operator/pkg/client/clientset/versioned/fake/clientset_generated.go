// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/banzaicloud/bank-vaults/operator/pkg/client/clientset/versioned"
	vaultv1alpha1 "github.com/banzaicloud/bank-vaults/operator/pkg/client/clientset/versioned/typed/vault/v1alpha1"
	fakevaultv1alpha1 "github.com/banzaicloud/bank-vaults/operator/pkg/client/clientset/versioned/typed/vault/v1alpha1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

var _ clientset.Interface = &Clientset{}

// VaultV1alpha1 retrieves the VaultV1alpha1Client
func (c *Clientset) VaultV1alpha1() vaultv1alpha1.VaultV1alpha1Interface {
	return &fakevaultv1alpha1.FakeVaultV1alpha1{Fake: &c.Fake}
}

// Vault retrieves the VaultV1alpha1Client
func (c *Clientset) Vault() vaultv1alpha1.VaultV1alpha1Interface {
	return &fakevaultv1alpha1.FakeVaultV1alpha1{Fake: &c.Fake}
}