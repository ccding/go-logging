// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
package logging

import (
	"io"
	"os"
	"sync"
	"time"
)

// pre-defined formats
const (
	basicFormat     = "%(name)s [%(levelname)s] %(asctime)s - %(message)s"
	richFormat      = "%(name)s [%(levelname)s] %(asctime)s - %(thread)d - %(module)s:%(filename)s:%(funcName)s:%(lineno)d- %(message)s"
	defaultFileName = "logging.log"
	configFileName  = "logging.conf"
)

// the logging struct
type logging struct {
	Name      string
	Level     int
	Format    string
	Out       io.Writer
	lock      sync.Mutex
	startTime int64
	Seqid     uint64
}

// create a new logger with simple configuration
func SimpleLogger(name string) *logging {
	return Logger(name, WARNING, basicFormat, os.Stdout)
}

// create a new logger with basic configuration
func BasicLogger(name string) *logging {
	return SimpleLogger(name)
}

// create a new logger with simple configuration
func RichLogger(name string) *logging {
	return FileLogger(name, NOTSET, richFormat, defaultFileName)
}

// create a new logger with file output
func FileLogger(name string, level int, format string, file string) *logging {
	out, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	return Logger(name, level, format, out)
}

// create a new logger
func Logger(name string, level int, format string, out io.Writer) *logging {
	logger := new(logging)
	logger.Name = name
	logger.Level = level
	logger.Format = format
	logger.Out = out
	logger.Seqid = 0

	logger.init()
	return logger
}

// initialize the logger
func (logger *logging) init() {
	logger.startTime = time.Now().UnixNano()
}

// get and set the configuration of the logger
func (logger *logging) GetName() string {
	return logger.Name
}

func (logger *logging) SetName(name string) {
	logger.Name = name
}

func (logger *logging) GetLevel() int {
	return logger.Level
}

func (logger *logging) SetLevel(level int) {
	logger.Level = level
}

func (logger *logging) GetLevelName() (string, bool) {
	name, ok := levelNames[logger.Level]
	return name, ok
}

func (logger *logging) SetLevelName(name string) {
	level, ok := levelValues[name]
	if ok {
		logger.Level = level
	}
}

func (logger *logging) GetFormat() string {
	return logger.Format
}

func (logger *logging) SetFormat(format string) {
	logger.Format = format
}

func (logger *logging) GetWriger() io.Writer {
	return logger.Out
}

func (logger *logging) SetWriter(out io.Writer) {
	logger.Out = out
}
