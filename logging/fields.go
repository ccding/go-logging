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
	"os"
	"path"
	"runtime"
	"sync/atomic"
	"time"
)

// the struct for each log record
type record struct {
	level      Level
	seqid      uint64
	pathname   string
	filename   string
	module     string
	lineno     int
	funcName   string
	thread     int
	threadName string
	process    int
	message    string
	time       time.Time
}

// this variable maps fields in format to relavent function signatures
var fields = map[string]func(*Logger, *record) interface{}{
	"name":            (*Logger).lname,
	"seqid":           (*Logger).nextSeqid,
	"levelno":         (*Logger).levelno,
	"levelname":       (*Logger).levelname,
	"pathname":        (*Logger).pathname,
	"filename":        (*Logger).filename,
	"module":          (*Logger).module,
	"lineno":          (*Logger).lineno,
	"funcName":        (*Logger).funcName,
	"created":         (*Logger).created,
	"asctime":         (*Logger).asctime,
	"msecs":           (*Logger).msecs,
	"relativeCreated": (*Logger).relativeCreated,
	"thread":          (*Logger).thread,
	"threadName":      (*Logger).threadName,
	"process":         (*Logger).process,
	"message":         (*Logger).message,
	"timestamp":       (*Logger).timestamp,
}

// if it fails to get some fields with string type, these fields are set to
// errString value
const errString = "???"

// GetGoId returns the id of goroutine, which is defined in ./get_go_id.c
func GetGoId() int32

// generate the runtime information, including pathname, function name,
// filename, line number.
func genRuntime(r *record) {
	calldepth := 5
	pc, file, line, ok := runtime.Caller(calldepth)
	if ok {
		// generate short function name
		fname := runtime.FuncForPC(pc).Name()
		fshort := fname
		for i := len(fname) - 1; i > 0; i-- {
			if fname[i] == '.' {
				fshort = fname[i+1:]
				break
			}
		}

		r.pathname = file
		r.funcName = fshort
		r.filename = path.Base(file)
		r.lineno = line
	} else {
		r.pathname = errString
		r.funcName = errString
		r.filename = errString
		// here we uses -1 rather than 0, because the default value in
		// golang is 0 and we should know the value is uninitialized
		// or failed to get
		r.lineno = -1
	}
}

// logger name
func (logger *Logger) lname(r *record) interface{} {
	return logger.name
}

// next sequence number
func (logger *Logger) nextSeqid(r *record) interface{} {
	if r.seqid == 0 {
		r.seqid = atomic.AddUint64(&(logger.seqid), 1)
	}
	return r.seqid
}

// log level number
func (logger *Logger) levelno(r *record) interface{} {
	return int(r.level)
}

// log level name
func (logger *Logger) levelname(r *record) interface{} {
	return levelNames[r.level]
}

// file name of calling logger, with whole path
func (logger *Logger) pathname(r *record) interface{} {
	if r.pathname == "" {
		genRuntime(r)
	}
	return r.pathname
}

// file name of calling logger
func (logger *Logger) filename(r *record) interface{} {
	if r.filename == "" {
		genRuntime(r)
	}
	return r.filename
}

// TODO: module name
func (logger *Logger) module(r *record) interface{} {
	return ""
}

// line number
func (logger *Logger) lineno(r *record) interface{} {
	if r.lineno == 0 {
		genRuntime(r)
	}
	return r.lineno
}

// function name
func (logger *Logger) funcName(r *record) interface{} {
	if r.funcName == "" {
		genRuntime(r)
	}
	return r.funcName
}

// timestamp of starting time
func (logger *Logger) created(r *record) interface{} {
	return logger.startTime.UnixNano()
}

// RFC3339Nano time
func (logger *Logger) asctime(r *record) interface{} {
	if r.time.IsZero() {
		r.time = time.Now()
	}
	return r.time.Format("2006-01-02 15:04:05.999999999")
}

// nanosecond of starting time
func (logger *Logger) msecs(r *record) interface{} {
	return logger.startTime.Nanosecond()
}

// nanosecond timestamp
func (logger *Logger) timestamp(r *record) interface{} {
	if r.time.IsZero() {
		r.time = time.Now()
	}
	return r.time.UnixNano()
}

// nanoseconds since logger created
func (logger *Logger) relativeCreated(r *record) interface{} {
	if r.time.IsZero() {
		r.time = time.Now()
	}
	return r.time.Sub(logger.startTime).Nanoseconds()
}

// thread id
func (logger *Logger) thread(r *record) interface{} {
	if r.thread == 0 {
		r.thread = int(GetGoId())
	}
	return r.thread
}

// thread name
func (logger *Logger) threadName(r *record) interface{} {
	if r.threadName == "" {
		r.threadName = fmt.Sprintf("Thread-%d", GetGoId())
	}
	return r.threadName
}

// Process ID
func (logger *Logger) process(r *record) interface{} {
	if r.process == 0 {
		r.process = os.Getpid()
	}
	return r.process
}

// the log message
func (logger *Logger) message(r *record) interface{} {
	return r.message
}
