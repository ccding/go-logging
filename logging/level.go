// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
package logging

import ()

type Level int

const (
	CRITICAL Level = 50
	FATAL    Level = CRITICAL
	ERROR    Level = 40
	WARNING  Level = 30
	WARN     Level = WARNING
	INFO     Level = 20
	DEBUG    Level = 10
	NOTSET   Level = 0
)

var levelNames = map[Level]string{
	CRITICAL: "CRITICAL",
	ERROR:    "ERROR",
	WARNING:  "WARNING",
	INFO:     "INFO",
	DEBUG:    "DEBUG",
	NOTSET:   "NOTSET",
}

var levelValues = map[string]Level{
	"CRITICAL": CRITICAL,
	"ERROR":    ERROR,
	"WARN":     WARNING,
	"WARNING":  WARNING,
	"INFO":     INFO,
	"DEBUG":    DEBUG,
	"NOTSET":   NOTSET,
}

type levelPair struct {
	name  string
	value Level
}

const maxAddLevelCacheSize = 10

var (
	levelPairs chan *levelPair
)

func init() {
	levelPairs = make(chan *levelPair, maxAddLevelCacheSize)

	go watchLevelUpdate()
}

func (level *Level) String() string {
	return levelNames[*level]
}

func GetLevelName(levelValue Level) string {
	return levelNames[levelValue]
}

func GetLevelValue(levelName string) Level {
	return levelValues[levelName]
}

func AddLevel(levelName string, levelValue Level) {
	SetLevel(levelName, levelValue)
}

func SetLevel(levelName string, levelValue Level) {
	level := new(levelPair)
	level.name = levelName
	level.value = levelValue
	levelPairs <- level
}

func watchLevelUpdate() {
	for {
		level := <-levelPairs
		levelValues[level.name] = level.value
		levelNames[level.value] = level.name
	}
}
