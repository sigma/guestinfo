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

// +build !debug

package guestinfo

func hvCommunicate(hvInput hvMessage) hvMessage {
	iax := hvInput.ax()
	ibx := hvInput.bx()
	icx := hvInput.cx()
	idx := hvInput.dx()
	idi := hvInput.di()
	isi := hvInput.si()

	oax, obx, ocx, odx, odi, osi := asmCommunicate(iax, ibx, icx, idx, idi, isi)
	return registers{_ax: oax, _bx: obx, _cx: ocx, _dx: odx, _di: odi, _si: osi}
}
