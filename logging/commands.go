// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
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
		fmt.Fprintln(logger.out, v...)
	}
}

func (logger *logging) logf(level Level, format string, v ...interface{}) {
	if int(level) >= int(logger.level) {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		fmt.Fprintf(logger.out, format+"\n", v...)
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
