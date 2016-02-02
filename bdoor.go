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

const (
	backdoorPort  = 0x5658
	backdoorMagic = 0x564D5868

	backdoorCmdGetVersion = 10
	backdoorCmdMessage    = 30

	backdoorChannelUseCookie = 0x80000000

	backdoorChannelOpen           = 0
	backdoorChannelSendSize       = 1
	backdoorChannelSendPayload    = 2
	backdoorChannelReceiveSize    = 3
	backdoorChannelReceivePayload = 4
	backdoorChannelReceiveStatus  = 5
	backdoorChannelClose          = 6

	backddorChannelSuccess = 0x0001
)
