// Copyright 2013, Cong Ding. All rights reserved.
// Use of this source code is governed by a GPLv2
// license that can be found in the LICENSE file.
//
// author: Cong Ding <dinggnu@gmail.com>
package logging

import (
	"strconv"
	"testing"
	"time"
)

func TestAddLevel(t *testing.T) {
	for i := 0; i < 1000; i++ {
		level := Level(i)
		name := "L" + strconv.Itoa(i)
		AddLevel(name, level)
		time.Sleep(time.Second / 1000)
		if levelNames[level] != name {
			t.Errorf("%v, %v, %v", level, name, levelNames[level])
		}
		if levelValues[name] != level {
			t.Errorf("%v, %v, %v", level, name, levelValues[name])
		}
		if GetLevelName(level) != name {
			t.Errorf("%v, %v, %v", level, name, levelNames[level])
		}
		if GetLevelValue(name) != level {
			t.Errorf("%v, %v, %v", level, name, levelValues[name])
		}
	}
}

func TestSetLevel(t *testing.T) {
	for i := 0; i < 1000; i++ {
		level := Level(i + 100)
		name := "S" + strconv.Itoa(i)
		SetLevel(name, level)
		time.Sleep(time.Second / 1000)
		if levelNames[level] != name {
			t.Errorf("%v, %v, %v", level, name, levelNames[level])
		}
		if levelValues[name] != level {
			t.Errorf("%v, %v, %v", level, name, levelValues[name])
		}
		if GetLevelName(level) != name {
			t.Errorf("%v, %v, %v", level, name, levelNames[level])
		}
		if GetLevelValue(name) != level {
			t.Errorf("%v, %v, %v", level, name, levelValues[name])
		}
	}
}
