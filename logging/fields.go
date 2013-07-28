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

type field func(*logging) interface{}

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
}

// the calling depth of these function, which is used to call the
// runtime.Caller() function to get the line number and file name
var calldepth = 5

const errorString = "???"

// GetGoId returns the id of goroutine, which is defined in ./get_go_id.c
func GetGoId() int32

func init() {
}

func (logger *logging) lname() interface{} {
	return logger.name
}

func (logger *logging) nextSeqid() interface{} {
	return atomic.AddUint64(&(logger.seqid), 1)
}

func (logger *logging) levelno() interface{} {
	return int(logger.level)
}

func (logger *logging) levelname() interface{} {
	return levelNames[logger.level]
}

func (logger *logging) pathname() interface{} {
	_, file, _, ok := runtime.Caller(calldepth)
	if !ok {
		file = errorString
	}
	return file
}

func (logger *logging) filename() interface{} {
	_, file, _, ok := runtime.Caller(calldepth)
	if !ok {
		file = errorString
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file
}

func (logger *logging) module() interface{} {
	return ""
}

func (logger *logging) lineno() interface{} {
	_, _, line, ok := runtime.Caller(calldepth)
	if !ok {
		line = 0
	}
	return line
}

func (logger *logging) funcName() interface{} {
	pc, _, _, ok := runtime.Caller(calldepth)
	if !ok {
		return errorString
	}
	return runtime.FuncForPC(pc).Name()
}

func (logger *logging) created() interface{} {
	return logger.startTime
}

func (logger *logging) asctime() interface{} {
	return time.Now().String()
}

func (logger *logging) msecs() interface{} {
	return logger.startTime % 1000
}

func (logger *logging) timestamp() interface{} {
	return time.Now().UnixNano()
}

func (logger *logging) relativeCreated() interface{} {
	return time.Now().UnixNano() - logger.startTime
}

func (logger *logging) thread() interface{} {
	return int(GetGoId())
}

func (logger *logging) threadName() interface{} {
	return fmt.Sprintf("Thread-%d", GetGoId())
}

// Process ID
func (logger *logging) process() interface{} {
	return os.Getpid()
}

func (logger *logging) message() interface{} {
	return ""
}
