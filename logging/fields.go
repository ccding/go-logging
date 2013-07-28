// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
package logging

import (
	"os"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"
	"path"
)

type field func(*logging) string

var fields = map[string] field {
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

// GetGoId returns the id of goroutine, which is defined in ./get_go_id.c
func GetGoId() int32

func (logger *logging) nextSeqid() string {
	return strconv.FormatUint(atomic.AddUint64(&(logger.seqid), 1), 10)
}

func (logger *logging) levelno() string {
	return strconv.Itoa(int(logger.level))
}

func (logger *logging) levelname() string {
	return levelNames[logger.level]
}

func (logger *logging) pathname() string {
	_, file, _, ok := runtime.Caller(calldepth)
	if !ok {
		file = "???"
	}
	return file
}

func (logger *logging) filename() string {
	return path.Base(logger.pathname())
}

func (logger *logging) module() string {
	return ""
}

func (logger *logging) lineno() string {
	_, _, line, ok := runtime.Caller(calldepth)
	if !ok {
		line = 0
	}
	return strconv.Itoa(line)
}

func (logger *logging) funcName() string {
	pc, _, _, ok := runtime.Caller(calldepth)
	if !ok {
		return "???"
	}
	return runtime.FuncForPC(pc).Name()
}

func (logger *logging) created() string {
	return strconv.FormatInt(logger.startTime, 10)
}

func (logger *logging) asctime() string {
	return time.Now().String()
}

func (logger *logging) msecs() string {
	return strconv.Itoa(int(logger.startTime % 1000))
}

func (logger *logging) timestamp() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func (logger *logging) relativeCreated() string {
	return strconv.FormatInt(time.Now().UnixNano()-logger.startTime, 10)
}

func (logger *logging) thread() string {
	return strconv.Itoa(int(GetGoId()))
}

func (logger *logging) threadName() string {
	return strconv.Itoa(int(GetGoId()))
}

// Process ID
func (logger *logging) process() string {
	return strconv.Itoa(os.Getpid())
}

func (logger *logging) message() string {
	return ""
}
