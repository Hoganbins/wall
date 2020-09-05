// Copyright 2020, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package stun

import (
	"sync"
	"testing"

	"github.com/q191201771/naza/pkg/nazalog"
)

var serverAddrList = []string{
	// dial udp: lookup stun01.sipphone.com: no such host
	// ----------
	//"stun01.sipphone.com",

	// XOR-MAPPED-ADDRESS
	// ----------
	"stun.l.google.com:19302",
	"stun4.l.google.com:19302",

	// XOR-MAPPED-ADDRESS
	// MAPPED-ADDRESS
	// RESPONSE-ORIGIN
	// OTHER-ADDRESS
	// SOFTWARE
	// FINGERPRINT
	// ----------
	"stun.freeswitch.org:3478",

	// MAPPED-ADDRESS
	// SOURCE_ADDRESS
	// CHANGED_ADDRESS
	// XOR-MAPPED-ADDRESS
	// SOFTWARE
	// ----------
	"stun.xten.com",
	"stun.ekiga.net",
	"stun.schlund.de",

	// MAPPED-ADDRESS
	// SOURCE_ADDRESS
	// CHANGED_ADDRESS
	// ----------
	"stun.ideasip.com",
	"stun.voiparound.com",
	"stun.voipbuster.com",
	"stun.voipstunt.com",
}

func TestClient(t *testing.T) {
	var wg sync.WaitGroup
	for _, s := range serverAddrList {
		wg.Add(1)
		go func(ss string) {
			var c Client
			ip, port, err := c.Query(ss, 200)
			nazalog.Debugf("server=%s, addr=%s:%d, err=%+v", ss, ip, port, err)
			wg.Done()
		}(s)
	}
	wg.Wait()
}

func TestServer(t *testing.T) {
}
