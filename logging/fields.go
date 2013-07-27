// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
package logging

import (
	"strconv"
	"time"
)

type field func(*logging) string

var fields = map[string]field{
	"name":            (*logging).name,
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

func init() {
}

func (logger *logging) name() string {
	return logger.Name
}

func (logger *logging) levelno() string {
	return strconv.Itoa(logger.Level)
}

func (logger *logging) levelname() string {
	return levelNames[logger.Level]
}

func (logger *logging) pathname() string {
	return ""
}

func (logger *logging) filename() string {
	return ""
}

func (logger *logging) module() string {
	return ""
}

func (logger *logging) lineno() string {
	return ""
}

func (logger *logging) funcName() string {
	return ""
}

func (logger *logging) created() string {
	return strconv.FormatInt(logger.startTime, 10)
}

func (logger *logging) asctime() string {
	return ""
}

func (logger *logging) msecs() string {
	return ""
}

func (logger *logging) timestamp() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func (logger *logging) relativeCreated() string {
	return strconv.FormatInt(time.Now().UnixNano()-logger.startTime, 10)
}

func (logger *logging) thread() string {
	return ""
}

func (logger *logging) threadName() string {
	return ""
}

func (logger *logging) process() string {
	return ""
}

func (logger *logging) message() string {
	return ""
}
