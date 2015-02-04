/*
 Copyright 2015 Bluek404

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package ukcp

import (
	"net"
	"time"

	"github.com/go-ukcp/ukcp/ikcp"
)

func NewUKCP(netType string, laddr *Addr) (*Conn, error) {
	switch netType {
	case "ukcp":
		netType = "udp"
	case "ukcp4":
		netType = "udp4"
	case "ukcp6":
		netType = "udp6"
	default:
		return nil, &net.OpError{
			Op:   "listen",
			Net:  netType,
			Addr: (*net.UDPAddr)(laddr),
			Err:  net.UnknownNetworkError(netType),
		}
	}
	conn, err := net.ListenUDP(netType, (*net.UDPAddr)(laddr))
	if err != nil {
		return nil, err
	}
	return &Conn{
		udpConn: conn,
		kcp:     ikcp.Create(uint32(time.Now().UnixNano()), nil),
	}, nil
}

type Conn struct {
	udpConn *net.UDPConn
	kcp     *ikcp.Ikcpcb
}

type Addr net.UDPAddr
