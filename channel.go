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

import "errors"

type Channel struct {
	Identifier uint32
	CookieHigh uint32
	CookieLow  uint32
}

type Protocol uint32

const RpcProtocol Protocol = 0x49435052

func NewChannel(proto Protocol) *Channel {
	comm := hvCommunicate(registers{
		_bx: uint32(proto) | uint32(backdoorChannelUseCookie),
		_cx: backdoorChannelOpen<<16 | backdoorCmdMessage,
	})

	return &Channel{
		Identifier: comm.dx(),
		CookieHigh: comm.si(),
		CookieLow:  comm.di(),
	}
}

func (c *Channel) Send(msg string) error {
	return nil
}

func (c *Channel) Receive() (string, error) {
	return "", nil
}

func (c *Channel) Close() error {
	comm := hvCommunicate(registers{
		_cx: backdoorChannelClose<<16 | backdoorCmdMessage,
		_dx: c.Identifier,
		_si: c.CookieHigh,
		_di: c.CookieLow,
	})

	if (comm.cx()>>16)&backddorChannelSuccess != 0 {
		return nil
	}
	return errors.New("Failed to close channel")
}
