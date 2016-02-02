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

import "github.com/sigma/bdoor"

var (
	virtualWorld = false
)

func init() {
	virtualWorld = bdoor.HypervisorPortCheck()
}

// IsVirtualWorld returns whether the code is running in a VMware virtual machine or no
func IsVirtualWorld() bool {
	return virtualWorld
}

func VMwareProduct() Product {
	return getVersionResponse{hvCommunicate(newGetVersionRequest())}.getProduct()
}
