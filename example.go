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
package main

import (
	"github.com/ccding/go-logging/logging"
	"time"
)

func main() {
	logger := logging.SimpleLogger("main")
	logger.SetLevel(logging.NOTSET)
	logger.Error("this is a test from error")
	logger.Debug("this is a test from debug")
	logger.Log(time.Now().UnixNano())
	time.Sleep(time.Second)
}
