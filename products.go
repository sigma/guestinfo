// Copyright 2016 Yann Hodique
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
