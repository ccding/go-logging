// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
package logging

import ()

const (
	CRITICAL int = 50
	FATAL    int = CRITICAL
	ERROR    int = 40
	WARNING  int = 30
	WARN     int = WARNING
	INFO     int = 20
	DEBUG    int = 10
	NOTSET   int = 0
)

var levelNames = map[int]string{
	CRITICAL: "CRITICAL",
	ERROR:    "ERROR",
	WARNING:  "WARNING",
	INFO:     "INFO",
	DEBUG:    "DEBUG",
	NOTSET:   "NOTSET",
}

var levelValues = map[string]int{
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
	value int
}

const maxAddLevelCacheSize = 10

var (
	levelPairs chan *levelPair
)

func init() {
	levelPairs = make(chan *levelPair, maxAddLevelCacheSize)

	go watchLevelUpdate()
}

func GetLeveName(levelValue int) string {
	return levelNames[levelValue]
}

func GetLevelValue(levelName string) int {
	return levelValues[levelName]
}

func AddLevel(levelName string, levelValue int) {
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
