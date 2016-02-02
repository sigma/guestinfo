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

func asmCommunicate(iax, ibx, icx, idx, idi, isi uint32) (oax, obx, ocx, odx, odi, osi uint32)

type registers struct {
	_ax, _bx, _cx, _dx, _di, _si uint32
}

func (r registers) ax() uint32 { return r._ax }

func (r registers) bx() uint32 { return r._bx }

func (r registers) cx() uint32 { return r._cx }

func (r registers) dx() uint32 { return r._dx }

func (r registers) di() uint32 { return r._di }

func (r registers) si() uint32 { return r._si }

type hvMessage interface {
	ax() uint32
	bx() uint32
	cx() uint32
	dx() uint32
	di() uint32
	si() uint32
}
