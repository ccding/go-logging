// Copyright 2013, Cong Ding. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// author: Cong Ding <dinggnu@gmail.com>
//
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
		seq := logger.nextSeqid()
		if seq != name {
			t.Errorf("%v, %v\n", seq, name)
		}
	}
}

func TestName(t *testing.T) {
	for i := 0; i < 1000; i++ {
		name := strconv.Itoa(i)
		logger := BasicLogger(name)
		if logger.Name() != name {
			t.Errorf("%v, %v\n", logger.Name(), name)
		}
	}
}
