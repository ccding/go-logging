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
	"strconv"
	"testing"
)

func TestAddLevel(t *testing.T) {
	for i := 0; i < 1000; i++ {
		level := Level(i)
		name := "L" + strconv.Itoa(i)
		AddLevel(name, level)
		if level.String() != name {
			t.Errorf("%v, %v, %v", level, name, level.String())
		}
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
		if level.String() != name {
			t.Errorf("%v, %v, %v", level, name, level.String())
		}
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
