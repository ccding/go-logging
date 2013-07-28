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
	"runtime"
	"sync/atomic"
	"time"
)

type log struct {
	level           Level
	name            string
	nextSeqid       uint64
	levelno         int
	levelname       string
	pathname        string
	filename        string
	module          string
	lineno          int
	funcName        string
	created         int64
	asctime         string
	msecs           int64
	relativeCreated int64
	thread          int
	threadName      string
	process         int
	message         string
	timestamp       int64
}

type field func(*logging, *log) interface{}

var fields = map[string]field{
	"name":            (*logging).lname,
	"nextSeqid":       (*logging).nextSeqid,
	"levelno":         (*logging).levelno,
	"levelname":       (*logging).levelname,
	"pathname":        (*logging).pathname,
	"filename":        (*logging).filename,
	"module":          (*logging).module,
	"lineno":          (*logging).lineno,
	"funcName":        (*logging).funcName,
	"created":         (*logging).created,
	"asctime":         (*logging).asctime,
	"msecs":           (*logging).msecs,
	"relativeCreated": (*logging).relativeCreated,
	"thread":          (*logging).thread,
	"threadName":      (*logging).threadName,
	"process":         (*logging).process,
	"message":         (*logging).message,
	"timestamp":       (*logging).timestamp,
}

const errString = "???"

// GetGoId returns the id of goroutine, which is defined in ./get_go_id.c
func GetGoId() int32

// generate the runtime information, including pathname, function name,
// filename, line number.
func genRuntime(l *log) {
	calldepth := 5
	pc, file, line, ok := runtime.Caller(calldepth)
	if ok {
		// generate short filename
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]
				break
			}
		}

		// generate short function name
		fname := runtime.FuncForPC(pc).Name()
		fshort := fname
		for i := len(fname) - 1; i > 0; i-- {
			if fname[i] == '.' {
				fshort = fname[i+1:]
				break
			}
		}

		l.pathname = file
		l.funcName = fshort
		l.filename = short
		l.lineno = line
	} else {
		l.pathname = errString
		l.funcName = errString
		l.filename = errString
		l.lineno = 0
	}
}

func (logger *logging) lname(l *log) interface{} {
	return logger.name
}

func (logger *logging) nextSeqid(l *log) interface{} {
	return atomic.AddUint64(&(logger.seqid), 1)
}

func (logger *logging) levelno(l *log) interface{} {
	return int(l.level)
}

func (logger *logging) levelname(l *log) interface{} {
	return levelNames[l.level]
}

func (logger *logging) pathname(l *log) interface{} {
	if l.pathname == "" {
		genRuntime(l)
	}
	return l.pathname
}

func (logger *logging) filename(l *log) interface{} {
	if l.filename == "" {
		genRuntime(l)
	}
	return l.filename
}

func (logger *logging) module(l *log) interface{} {
	return ""
}

func (logger *logging) lineno(l *log) interface{} {
	if l.lineno == 0 {
		genRuntime(l)
	}
	return l.lineno
}

func (logger *logging) funcName(l *log) interface{} {
	if l.funcName == "" {
		genRuntime(l)
	}
	return l.funcName
}

func (logger *logging) created(l *log) interface{} {
	return logger.startTime
}

func (logger *logging) asctime(l *log) interface{} {
	if l.asctime == "" {
		l.asctime = time.Now().String()
	}
	return l.asctime
}

func (logger *logging) msecs(l *log) interface{} {
	if l.msecs == 0 {
		l.msecs = logger.startTime % 1000
	}
	return l.msecs
}

func (logger *logging) timestamp(l *log) interface{} {
	if l.timestamp == 0 {
		l.timestamp = time.Now().UnixNano()
	}
	return l.timestamp
}

func (logger *logging) relativeCreated(l *log) interface{} {
	if l.relativeCreated == 0 {
		l.relativeCreated = time.Now().UnixNano() - logger.startTime
	}
	return l.relativeCreated
}

func (logger *logging) thread(l *log) interface{} {
	if l.thread == 0 {
		l.thread = int(GetGoId())
	}
	return l.thread
}

func (logger *logging) threadName(l *log) interface{} {
	if l.threadName == "" {
		l.threadName = fmt.Sprintf("Thread-%d", GetGoId())
	}
	return l.threadName
}

// Process ID
func (logger *logging) process(l *log) interface{} {
	if l.process == 0 {
		l.process = os.Getpid()
	}
	return l.process
}

func (logger *logging) message(l *log) interface{} {
	return l.message
}
