// Copyright 2012-2016 The GoSNMP Authors. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

package gosnmp

import (
	"fmt"
	"log"
	"net"
)

//
// Sending Traps ie GoSNMP acting as an Agent
//

// SendTrap sends a SNMP Trap(currently v2c/v3 only)/
func (x *GoSNMP) SendTrap(pdus []SnmpPDU) (result *SnmpPacket, err error) {
	switch x.Version {
	case Version2c, Version3:
		// x.mkSnmpPacket doesn't exist
		// will send even work??
		packetOut := x.mkSnmpPacket(SNMPv2Trap, 0, 0)
		return x.send(pdus, packetOut)
	default:
		err = fmt.Errorf("SendTrap doesn't support %s", x.Version)
		return nil, err
	}
}

//
// Receiving Traps ie GoSNMP acting as an NMS
//
// GoSNMP.unmarshal() currently only handles SNMPv2Trap (ie v2c, v3)
//

// A TrapListener defineds parameters for running a SNMP Trap receiver.
// nil values will be replaced by default values.
type TrapListener struct {
	OnNewTrap func(s *SnmpPacket, u *net.UDPAddr)
	Params    *GoSNMP
}

// Listen listens on the UDP address addr and calls the OnNewTrap
// function specified in *TrapListener for every trap recieved.
func (t *TrapListener) Listen(addr string) (err error) {
	if t.Params == nil {
		t.Params = Default
	}

	if t.OnNewTrap == nil {
		t.OnNewTrap = debugTrapHandler
	}

	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	for {
		var buf [4096]byte
		rlen, remote, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			if t.Params.loggingEnabled {
				t.Params.Logger.Printf("TrapListener: error in read %s\n", err)
			}
		}

		msg := buf[:rlen]
		traps := t.Params.unmarshalTrap(msg)
		t.OnNewTrap(traps, remote)
	}
}

// Default trap handler
func debugTrapHandler(s *SnmpPacket, u *net.UDPAddr) {
	log.Printf("got trapdata from %+v: %+v\n", u, s)
}

// Unmarshal SNMP Trap
func (x *GoSNMP) unmarshalTrap(trap []byte) (result *SnmpPacket) {
	result = new(SnmpPacket)
	err := x.unmarshal(trap, result)
	if err != nil {
		if x.loggingEnabled {
			x.Logger.Printf("unmarshalTrap: %s\n", err)
		}
	}

	return result
}
