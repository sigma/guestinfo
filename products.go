package guestinfo

var (
	productCodes = []string{
		"something weird that seems to act as VMware stuff",
		"VMware Express",
		"VMware ESX(i)",
		"VMware Server",
		"VMware Workstation/Fusion",
		"VMware ACE",
	}
)

// Product represents a flavor of VMware platform
type Product int

const (
	// Express represents VMware Express
	Express Product = iota
	// ESX represents VMware ESX(i)
	ESX
	// Server represents VMware Server
	Server
	// WS represents VMware Workstation or Fusion
	WS
	// ACE represents VMware ACE
	ACE
)

// String returns the name of the product
func (p Product) String() string {
	return productCodes[p]
}
