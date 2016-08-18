// Copyright 2012-2014 The GoSNMP Authors. All rights reserved.  Use of this
// source code is governed by a BSD-style license that can be found in the
// LICENSE file.

package gosnmp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	// "testing"
	// "time"
)

var _ = fmt.Sprintf("dummy") // dummy
var _ = ioutil.Discard       // dummy
var _ = os.DevNull           // dummy
var _ = bytes.MinRead        // dummy

/*
$ snmptrap -v 2c -c public 192.168.1.10 '' SNMPv2-MIB::system SNMPv2-MIB::sysDescr.0 s "red laptop" SNMPv2-MIB::sysServices.0 i "5"

Simple Network Management Protocol
    version: v2c (1)
    community: public
    data: snmpV2-trap (7)
        snmpV2-trap
            request-id: 1271509950
            error-status: noError (0)
            error-index: 0
            variable-bindings: 5 items
                1.3.6.1.2.1.1.3.0: 1034156
                    Object Name: 1.3.6.1.2.1.1.3.0 (iso.3.6.1.2.1.1.3.0)
                    Value (Timeticks): 1034156
                1.3.6.1.6.3.1.1.4.1.0: 1.3.6.1.2.1.1 (iso.3.6.1.2.1.1)
                    Object Name: 1.3.6.1.6.3.1.1.4.1.0 (iso.3.6.1.6.3.1.1.4.1.0)
                    Value (OID): 1.3.6.1.2.1.1 (iso.3.6.1.2.1.1)
                1.3.6.1.2.1.1.1.0: 726564206c6170746f70
                    Object Name: 1.3.6.1.2.1.1.1.0 (iso.3.6.1.2.1.1.1.0)
                    Value (OctetString): 726564206c6170746f70
                        Variable-binding-string: red laptop
                1.3.6.1.2.1.1.7.0:
                    Object Name: 1.3.6.1.2.1.1.7.0 (iso.3.6.1.2.1.1.7.0)
                    Value (Integer32): 5
                1.3.6.1.2.1.1.2: 1.3.6.1.4.1.2.3.4.5 (iso.3.6.1.4.1.2.3.4.5)
                    Object Name: 1.3.6.1.2.1.1.2 (iso.3.6.1.2.1.1.2)
                    Value (OID): 1.3.6.1.4.1.2.3.4.5 (iso.3.6.1.4.1.2.3.4.5)
*/

// trap_v2 returns raw file data.
func trap_v2() []byte {
	return []byte{
		// 0     1     2     3     4     5     6     7     8     9     a     b     c     d     e     f
		0x30, 0x7f, 0x02, 0x01, 0x01, 0x04, 0x06, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0xa7, 0x72, 0x02,
		0x04, 0x4b, 0xc9, 0xb3, 0xbe, 0x02, 0x01, 0x00, 0x02, 0x01, 0x00, 0x30, 0x64, 0x30, 0x0f, 0x06,
		0x08, 0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x03, 0x00, 0x43, 0x03, 0x0f, 0xc7, 0xac, 0x30, 0x14,
		0x06, 0x0a, 0x2b, 0x06, 0x01, 0x06, 0x03, 0x01, 0x01, 0x04, 0x01, 0x00, 0x06, 0x06, 0x2b, 0x06,
		0x01, 0x02, 0x01, 0x01, 0x30, 0x16, 0x06, 0x08, 0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x01, 0x00,
		0x04, 0x0a, 0x72, 0x65, 0x64, 0x20, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x30, 0x0d, 0x06, 0x08,
		0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x07, 0x00, 0x02, 0x01, 0x05, 0x30, 0x14, 0x06, 0x07, 0x2b,
		0x06, 0x01, 0x02, 0x01, 0x01, 0x02, 0x06, 0x09, 0x2b, 0x06, 0x01, 0x04, 0x01, 0x02, 0x03, 0x04,
		0x05,
	}
}
