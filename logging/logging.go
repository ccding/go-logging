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

// Package logging implements log library for other applications. It provides
// functions Debug, Info, Warning, Error, Critical, and formatting version
// Logf.
//
// Example:
//
//	logger := logging.SimpleLogger("main")
//	logger.SetLevel(logging.WARNING)
//	logger.Error("test for error")
//	logger.Warning("test for warning", "second parameter")
//	logger.Debug("test for debug")
//
package logging

import (
	"errors"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

// Pre-defined formats
const (
	defaultFileName = "logging.log"
	configFileName  = "logging.conf"
)

// Logger is the logging struct.
type Logger struct {
	// Be careful of the alignment issue of the variable seqid because it
	// uses the sync/atomic.AddUint64() operation. If the alignment is
	// wrong, it will cause a panic. To solve the alignment issue in an
	// easy way, we put seqid to the beginning of the structure.
	seqid     uint64
	name      string
	level     Level
	format    string
	fargs     []string
	out       io.Writer
	lock      sync.Mutex
	startTime time.Time
	sync      bool
	queue     chan string
	flush     chan bool
}

// SimpleLogger creates a new logger with simple configuration.
func SimpleLogger(name string) (*Logger, error) {
	return createLogger(name, WARNING, BasicFormat, os.Stdout, true)
}

// BasicLogger creates a new logger with basic configuration.
func BasicLogger(name string) (*Logger, error) {
	return FileLogger(name, WARNING, BasicFormat, defaultFileName, true)
}

// RichLogger creates a new logger with simple configuration.
func RichLogger(name string) (*Logger, error) {
	return FileLogger(name, NOTSET, RichFormat, defaultFileName, true)
}

// FileLogger creates a new logger with file output.
func FileLogger(name string, level Level, format string, file string, sync bool) (*Logger, error) {
	out, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	return createLogger(name, level, format, out, sync)
}

// createLogger create a new logger
func createLogger(name string, level Level, format string, out io.Writer, sync bool) (*Logger, error) {
	logger := new(Logger)

	err := logger.SetFormat(format)
	if err != nil {
		return logger, err
	}

	// asign values to logger
	logger.name = name
	logger.level = level
	logger.out = out
	logger.seqid = 0
	logger.sync = sync
	logger.queue = make(chan string)
	logger.flush = make(chan bool)
	logger.startTime = time.Now()

	// start watcher and timer
	go logger.watcher()
	go logger.timer()

	return logger, nil
}

func (logger *Logger) Flush() {
	logger.flush <- true
}

// Get and set the configuration of the logger

func (logger *Logger) Name() string {
	return logger.name
}

func (logger *Logger) SetName(name string) {
	logger.name = name
}

func (logger *Logger) Level() Level {
	return logger.level
}

func (logger *Logger) SetLevel(level Level) {
	logger.level = Level(level)
}

func (logger *Logger) LevelName() string {
	name, _ := levelNames[logger.level]
	return name
}

func (logger *Logger) SetLevelName(name string) {
	level, ok := levelValues[name]
	if ok {
		logger.level = level
	}
}

func (logger *Logger) Format() string {
	return logger.format
}

func (logger *Logger) Fargs() []string {
	return logger.fargs
}

func (logger *Logger) SetFormat(format string) error {
	// partially check the legality of format
	fts := strings.Split(format, "\n")
	if len(fts) != 2 {
		return errors.New("logging format error")
	}
	logger.format = fts[0]
	logger.fargs = strings.Split(fts[1], ",")
	for k, v := range logger.fargs {
		tv := strings.TrimSpace(v)
		_, ok := fields[tv]
		if ok == false {
			return errors.New("logging format error")
		}
		logger.fargs[k] = tv
	}
	return nil
}

func (logger *Logger) Writer() io.Writer {
	return logger.out
}

func (logger *Logger) AddWriter(out io.Writer) {
	logger.out = io.MultiWriter(logger.out, out)
}

func (logger *Logger) AddWriters(out ...io.Writer) {
	logger.out = io.MultiWriter(append(out, logger.out)...)
}

func (logger *Logger) SetWriter(out io.Writer) {
	logger.out = out
}

func (logger *Logger) SetWriters(out ...io.Writer) {
	logger.out = io.MultiWriter(out...)
}

func (logger *Logger) Sync() bool {
	return logger.sync
}

func (logger *Logger) SetSync(sync bool) {
	logger.sync = sync
}
