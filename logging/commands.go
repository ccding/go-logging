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
	"fmt"
)

// receive log request from the client, and start new goroute to record it
func (logger *logging) Logln(level Level, v ...interface{}) {
	go logger.logln(level, v...)
}

func (logger *logging) Logf(level Level, format string, v ...interface{}) {
	go logger.logf(level, format, v...)
}

// record log v... with level `level'
func (logger *logging) logln(level Level, v ...interface{}) {
	if int(level) >= int(logger.level) {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		message := fmt.Sprint(v...)
		fmt.Fprintln(logger.out, logger.genLog(level, message))
	}
}

func (logger *logging) logf(level Level, format string, v ...interface{}) {
	if int(level) >= int(logger.level) {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		message := fmt.Sprintf(format, v...)
		fmt.Fprintln(logger.out, logger.genLog(level, message))
	}
}

// other quick commands
func (logger *logging) Critical(v ...interface{}) {
	go logger.logln(CRITICAL, v...)
}

func (logger *logging) Fatal(v ...interface{}) {
	go logger.logln(CRITICAL, v...)
}

func (logger *logging) Error(v ...interface{}) {
	go logger.logln(ERROR, v...)
}

func (logger *logging) Warn(v ...interface{}) {
	go logger.logln(WARNING, v...)
}

func (logger *logging) Warning(v ...interface{}) {
	go logger.logln(WARNING, v...)
}

func (logger *logging) Info(v ...interface{}) {
	go logger.logln(INFO, v...)
}

func (logger *logging) Debug(v ...interface{}) {
	go logger.logln(DEBUG, v...)
}

func (logger *logging) Log(v ...interface{}) {
	go logger.logln(NOTSET, v...)
}
