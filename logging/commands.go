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
func (logger *logging) Logln(level int, v ...interface{}) {
	go logger.logln(level, v...)
}

func (logger *logging) Logf(level int, format string, v ...interface{}) {
	go logger.logf(level, format, v...)
}

// record log v... with level `level'
func (logger *logging) logln(level int, v ...interface{}) {
	if level >= logger.Level {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		fmt.Fprintln(logger.Out, v...)
	}
}

func (logger *logging) logf(level int, format string, v ...interface{}) {
	if level >= logger.Level {
		logger.lock.Lock()
		defer logger.lock.Unlock()
		fmt.Fprintf(logger.Out, format+"\n", v...)
	}
}

// other quick commands
func (logger *logging) Critical(v ...interface{}) {
	logger.Logln(CRITICAL, v...)
}

func (logger *logging) Fatal(v ...interface{}) {
	logger.Logln(CRITICAL, v...)
}

func (logger *logging) Error(v ...interface{}) {
	logger.Logln(ERROR, v...)
}

func (logger *logging) Warn(v ...interface{}) {
	logger.Logln(WARNING, v...)
}

func (logger *logging) Warning(v ...interface{}) {
	logger.Logln(WARNING, v...)
}

func (logger *logging) Info(v ...interface{}) {
	logger.Logln(INFO, v...)
}

func (logger *logging) Debug(v ...interface{}) {
	logger.Logln(DEBUG, v...)
}

func (logger *logging) Log(v ...interface{}) {
	logger.Logln(NOTSET, v...)
}
