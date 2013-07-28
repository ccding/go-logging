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
	"io"
	"os"
	"sync"
	"time"
)

// pre-defined formats
const (
	defaultFileName = "logging.log"
	configFileName  = "logging.conf"
)

// the logging struct
type logging struct {
	name      string
	level     Level
	format    string
	out       io.Writer
	lock      sync.Mutex
	startTime time.Time
	seqid     uint64
}

// create a new logger with simple configuration
func SimpleLogger(name string) *logging {
	return Logger(name, WARNING, BasicFormat, os.Stdout)
}

// create a new logger with basic configuration
func BasicLogger(name string) *logging {
	return SimpleLogger(name)
}

// create a new logger with simple configuration
func RichLogger(name string) *logging {
	return FileLogger(name, NOTSET, RichFormat, defaultFileName)
}

// create a new logger with file output
func FileLogger(name string, level Level, format string, file string) *logging {
	out, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	return Logger(name, level, format, out)
}

// create a new logger
func Logger(name string, level Level, format string, out io.Writer) *logging {
	logger := new(logging)
	logger.name = name
	logger.level = level
	logger.format = format
	logger.out = out
	logger.seqid = 0

	logger.init()
	return logger
}

// initialize the logger
func (logger *logging) init() {
	logger.startTime = time.Now()
}

// get and set the configuration of the logger
func (logger *logging) Name() string {
	return logger.name
}

func (logger *logging) SetName(name string) {
	logger.name = name
}

func (logger *logging) Level() Level {
	return logger.level
}

func (logger *logging) SetLevel(level Level) {
	logger.level = Level(level)
}

func (logger *logging) LevelName() string {
	name, _ := levelNames[logger.level]
	return name
}

func (logger *logging) SetLevelName(name string) {
	level, ok := levelValues[name]
	if ok {
		logger.level = level
	}
}

func (logger *logging) Format() string {
	return logger.format
}

func (logger *logging) SetFormat(format string) {
	logger.format = format
}

func (logger *logging) Writer() io.Writer {
	return logger.out
}

func (logger *logging) SetWriter(out io.Writer) {
	logger.out = out
}
