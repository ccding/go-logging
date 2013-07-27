// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
package logging

import (
	"strconv"
	"testing"
)

func empty() {
}

func TestGetGoId(t *testing.T) {
	for i := 0; i < 1000; i++ {
		goid := int(GetGoId())
		go empty()
		goid2 := int(GetGoId())
		if goid != goid2 {
			t.Errorf("%v, %v\n", goid, goid2)
		}
	}
}

func TestSeqid(t *testing.T) {
	logger := BasicLogger("test")
	for i := 0; i < 1000; i++ {
		name := strconv.Itoa(i + 1)
		seq := logger.seqid()
		if seq != name {
			t.Errorf("%v, %v\n", seq, name)
		}
	}
}

func TestName(t *testing.T) {
	for i := 0; i < 1000; i++ {
		name := strconv.Itoa(i)
		logger := BasicLogger(name)
		if logger.name() != name {
			t.Errorf("%v, %v\n", logger.name(), name)
		}
	}
}
