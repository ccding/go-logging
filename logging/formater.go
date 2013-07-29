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
	"strings"
)

// pre-defined formats
const (
	BasicFormat = "%s [%6s] %30s - %s\n name,levelname,asctime,message"
	RichFormat  = "%s [%6s] %d %30s - %d - %s:%s:%d - %s\n name, levelname, seqid, asctime, thread, filename, funcName, lineno, message"
)

// generate log string from the format setting
func (logger *Logger) genLog(level Level, message string) string {
	format := strings.Split(logger.format, "\n")
	if len(format) != 2 {
		return "logging format error"
	}
	args := strings.Split(format[1], ",")
	fs := make([]interface{}, len(args))
	r := new(record)
	r.message = message
	r.level = level
	for k, v := range args {
		fs[k] = fields[strings.TrimSpace(v)](logger, r)
	}
	return fmt.Sprintf(format[0], fs...)
}
