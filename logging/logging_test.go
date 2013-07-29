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
	"testing"
	"time"
)

func BenchmarkSync(b *testing.B) {
	logger, _ := RichLogger("main")
	logger.SetLevel(NOTSET)
	logger.SetSync(true)
	for i := 0; i < b.N; i++ {
		logger.Error("this is a test from error")
		logger.Debug("this is a test from debug")
		logger.Notset("orz", time.Now().UnixNano())
		logger.Error("this is a test from error")
		logger.Debug("this is a test from debug")
		logger.Notset("orz", time.Now().UnixNano())
	}
	logger.Flush()
}

func BenchmarkAsync(b *testing.B) {
	logger, _ := RichLogger("main")
	logger.SetLevel(NOTSET)
	logger.SetSync(false)
	for i := 0; i < b.N; i++ {
		logger.Error("this is a test from error")
		logger.Debug("this is a test from debug")
		logger.Notset("orz", time.Now().UnixNano())
		logger.Error("this is a test from error")
		logger.Debug("this is a test from debug")
		logger.Notset("orz", time.Now().UnixNano())
	}
	logger.Flush()
}

func BenchmarkBasicSync(b *testing.B) {
	logger, _ := BasicLogger("main")
	logger.SetLevel(NOTSET)
	logger.SetSync(true)
	for i := 0; i < b.N; i++ {
		logger.Error("this is a test from error")
		logger.Debug("this is a test from debug")
		logger.Notset("orz", time.Now().UnixNano())
		logger.Error("this is a test from error")
		logger.Debug("this is a test from debug")
		logger.Notset("orz", time.Now().UnixNano())
	}
	logger.Flush()
}

func BenchmarkBasicAsync(b *testing.B) {
	logger, _ := BasicLogger("main")
	logger.SetLevel(NOTSET)
	logger.SetSync(false)
	for i := 0; i < b.N; i++ {
		logger.Error("this is a test from error")
		logger.Debug("this is a test from debug")
		logger.Notset("orz", time.Now().UnixNano())
		logger.Error("this is a test from error")
		logger.Debug("this is a test from debug")
		logger.Notset("orz", time.Now().UnixNano())
	}
	logger.Flush()
}
