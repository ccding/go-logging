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

// The struct for each log record
type record struct {
	level      Level
	seqid      uint64
	pathname   string
	filename   string
	module     string
	lineno     int
	funcname   string
	thread     int
	threadName string
	process    int
	message    string
	time       time.Time
}

// This variable maps fields in format to relavent function signatures
var fields = map[string]func(*Logger, *record) interface{}{
	"name":            (*Logger).lname,
	"seqid":           (*Logger).nextSeqid,
	"levelno":         (*Logger).levelno,
	"levelname":       (*Logger).levelname,
	"pathname":        (*Logger).pathname,
	"filename":        (*Logger).filename,
	"module":          (*Logger).module,
	"lineno":          (*Logger).lineno,
	"funcname":        (*Logger).funcname,
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

// If it fails to get some fields with string type, these fields are set to
// errString value.
const errString = "???"

// GetGoID returns the id of goroutine, which is defined in ./get_go_id.c
func GetGoID() int32

// genRuntime generates the runtime information, including pathname, function
// name, filename, line number.
func genRuntime(r *record) {
	calldepth := 5
	pc, file, line, ok := runtime.Caller(calldepth)
	if ok {
		// Generate short function name
		fname := runtime.FuncForPC(pc).Name()
		fshort := fname
		for i := len(fname) - 1; i > 0; i-- {
			if fname[i] == '.' {
				fshort = fname[i+1:]
				break
			}
		}

		r.pathname = file
		r.funcname = fshort
		r.filename = path.Base(file)
		r.lineno = line
	} else {
		r.pathname = errString
		r.funcname = errString
		r.filename = errString
		// Here we uses -1 rather than 0, because the default value in
		// golang is 0 and we should know the value is uninitialized
		// or failed to get
		r.lineno = -1
	}
}

// Logger name
func (logger *Logger) lname(r *record) interface{} {
	return logger.name
}

// Next sequence number
func (logger *Logger) nextSeqid(r *record) interface{} {
	if r.seqid == 0 {
		r.seqid = atomic.AddUint64(&(logger.seqid), 1)
	}
	return r.seqid
}

// Log level number
func (logger *Logger) levelno(r *record) interface{} {
	return int(r.level)
}

// Log level name
func (logger *Logger) levelname(r *record) interface{} {
	return levelNames[r.level]
}

// File name of calling logger, with whole path
func (logger *Logger) pathname(r *record) interface{} {
	if r.pathname == "" {
		genRuntime(r)
	}
	return r.pathname
}

// File name of calling logger
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

// Line number
func (logger *Logger) lineno(r *record) interface{} {
	if r.lineno == 0 {
		genRuntime(r)
	}
	return r.lineno
}

// Function name
func (logger *Logger) funcname(r *record) interface{} {
	if r.funcname == "" {
		genRuntime(r)
	}
	return r.funcname
}

// Timestamp of starting time
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

// Nanosecond of starting time
func (logger *Logger) msecs(r *record) interface{} {
	return logger.startTime.Nanosecond()
}

// Nanosecond timestamp
func (logger *Logger) timestamp(r *record) interface{} {
	if r.time.IsZero() {
		r.time = time.Now()
	}
	return r.time.UnixNano()
}

// Nanoseconds since logger created
func (logger *Logger) relativeCreated(r *record) interface{} {
	if r.time.IsZero() {
		r.time = time.Now()
	}
	return r.time.Sub(logger.startTime).Nanoseconds()
}

// Thread ID
func (logger *Logger) thread(r *record) interface{} {
	if r.thread == 0 {
		r.thread = int(GetGoID())
	}
	return r.thread
}

// Thread name
func (logger *Logger) threadName(r *record) interface{} {
	if r.threadName == "" {
		r.threadName = fmt.Sprintf("Thread-%d", GetGoID())
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

// The log message
func (logger *Logger) message(r *record) interface{} {
	return r.message
}
