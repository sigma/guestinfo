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

type getVersionRequest struct {
	registers
}

func newGetVersionRequest() getVersionRequest {
	return getVersionRequest{
		registers{
			_ax: backdoorMagic,
			_bx: 0xFFFFFFFF,
			_cx: backdoorCmdGetVersion,
			_dx: backdoorPort,
		},
	}
}

type getVersionResponse struct {
	hvMessage
}

func (r getVersionResponse) isVMware() bool {
	return r.hvMessage.bx() == backdoorMagic
}

func (r getVersionResponse) getProduct() Product {
	code := r.hvMessage.cx()
	if code > uint32(len(productCodes)) {
		code = 0
	}
	return Product(code)
}
