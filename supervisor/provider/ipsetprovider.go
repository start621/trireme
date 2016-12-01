package provider

import "github.com/bvandewalle/go-ipset/ipset"

// IpsetProvider returns a fabric for Ipset.
type IpsetProvider interface {
	NewIPset(name string, hasht string, p *ipset.Params) (Ipset, error)
	DestroyAll() error
}

// Ipset is an abstraction of all the methods an implementation of userspace
// ipsets need to provide.
type Ipset interface {
	Add(entry string, timeout int) error
	AddOption(entry string, option string, timeout int) error
	Del(entry string) error
	Destroy() error
	Flush() error
	Test(entry string) (bool, error)
}

type goIpset struct{}

// NewIPset returns an IpsetProvider interface based on the go-ipset
// external package.
func (i *goIpset) NewIPset(name string, hasht string, p *ipset.Params) (Ipset, error) {
	return ipset.New(name, hasht, p)
}

func (i *goIpset) DestroyAll() error {
	return ipset.DestroyAll()
}

// NewGoIPsetProvider Return a Go IPSet Provider
func NewGoIPsetProvider() IpsetProvider {
	return &goIpset{}
}
