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

// Logln receives log request from the client. The request includes a set of
// variables.
func (logger *Logger) Logln(level Level, v ...interface{}) {
	logger.logln(level, v...)
}

// Logf receives log request from the client. The request has a string
// parameter to describe the format of output.
func (logger *Logger) Logf(level Level, format string, v ...interface{}) {
	logger.logf(level, format, v...)
}

// logln records log v... with level `level'.
func (logger *Logger) logln(level Level, v ...interface{}) {
	if int(level) >= int(logger.level) {
		message := fmt.Sprint(v...)
		message = logger.genLog(level, message)
		if logger.sync {
			logger.printLog(message)
		} else {
			go logger.printLog(message)
		}
	}
}

// logf records log v... with level `level'.
func (logger *Logger) logf(level Level, format string, v ...interface{}) {
	if int(level) >= int(logger.level) {
		message := fmt.Sprintf(format, v...)
		message = logger.genLog(level, message)
		if logger.sync {
			logger.printLog(message)
		} else {
			go logger.printLog(message)
		}
	}
}

// printLog is to print log to file, stdout, or others.
func (logger *Logger) printLog(message string) {
	logger.lock.Lock()
	defer logger.lock.Unlock()
	fmt.Fprintln(logger.out, message)
}

// other quick commands for different level

func (logger *Logger) Critical(v ...interface{}) {
	logger.logln(CRITICAL, v...)
}

func (logger *Logger) Fatal(v ...interface{}) {
	logger.logln(CRITICAL, v...)
}

func (logger *Logger) Error(v ...interface{}) {
	logger.logln(ERROR, v...)
}

func (logger *Logger) Warn(v ...interface{}) {
	logger.logln(WARNING, v...)
}

func (logger *Logger) Warning(v ...interface{}) {
	logger.logln(WARNING, v...)
}

func (logger *Logger) Info(v ...interface{}) {
	logger.logln(INFO, v...)
}

func (logger *Logger) Debug(v ...interface{}) {
	logger.logln(DEBUG, v...)
}

func (logger *Logger) Log(v ...interface{}) {
	logger.logln(NOTSET, v...)
}
