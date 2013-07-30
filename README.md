#go-logging

A high-performance logging library for golang.

* Simple: it supports only necessary operations
* Fast: it is very efficient and under improvement


## Getting Started

go-logging is used like any other Go libraries. You can simply
```go
import "github.com/ccding/go-logging"
```
to use the library.

Here is a simple example.
```go
package main

import (
	"github.com/ccding/go-logging"
)

func main() {
	logger, _ := logging.SimpleLogger("main")
	logger.Error("this is a test from error")
	logger.Destroy()
}
```

### Examples and Benchmarks
The steps below will download the library source code to
`${GOPATH}/src/github.com/ccding/go-logging`.
```bash
go get github.com/ccding/go-logging
```

Given the source code downloaded, it makes you be able to run the example
```bash
cd ${GOPATH}/src/github.com/ccding/go-logging
go run example/example.go
```
and benchmarks
```bash
go test -v -bench .
```

### Settings

#### Logging Format
The logging format is described by a string, which has two parts separated by
`\n`. The first part describes the format of the log, and the second part
lists all the fields to be shown in the log. In other word, the first part is
the first parameter `format` of `fmt.Printf(format string, v ...interface{})`,
and the second part describes the second parameter `v` of it. It is not
allowed to have `\n` in the first part.  The fields in the second part are
separated by comma `,`, while extra blank spaces are allowed.  An example of
the format string is
```
%s [%6s] %30s - %s\n name,levelname,time,message
```
which is the pre-defined `BasicFormat` used by `BasicLogger()` and
`SimpleLogger()`.

It supports the following fields for the second part of the format.
```
"name":      string   %s  name of the logger
"seqid":     uint64   %d  sequence number
"levelno":   int      %d  level number
"levelname": string   %s  level name
"created":   int64    %d  starting time of the logger
"nsecs":     int64    %d  nanosecond of the starting time
"time":      string   %s  record created time
"timestamp": int64    %d  timestamp of record
"rtime":     int64    %d  relative time since started
"filename":  string   %s  source filename of the caller
"pathname":  string   %s  filename with path
"module":    string   %s  executable filename // TODO
"lineno":    int      %d  line number in source code
"funcname":  string   %s  function name of the caller
"thread":    int32    %d  thread id
"process":   int      %d  process id
"message":   string   %d  logger message
```

#### Logger
The logger supports the following operations, all of which are not
thread-safe. If you are calling them in multiple thread, please lock them
each time before you call these functions.

It has the following functions to create a logger.
```
SimpleLogger(name string) (*Logger, error)
BasicLogger(name string) (*Logger, error)
RichLogger(name string) (*Logger, error)
FileLogger(name string, level Level, format string, file string, sync bool) (*Logger, error)
WriterLogger(name string, level Level, format string, out io.Writer, sync bool) (*Logger, error)
```
The meanings of these fields are
```
name           string         logger name
level          Level          record level higher than this will be printed
format         string         format configuration
out            io.Writer      writer
startTime      time.Time      start time of the logger
sync           bool           use sync or async way to record logs
timeFormat     string         format for time
```

There are these levels in logging.
```
CRITICAL     50
FATAL        CRITICAL
ERROR        40
WARNING      30
WARN         WARNING
INFO         20
DEBUG        10
NOTSET       0
```

It has these functions to operate on the logger:
```
Destroy()                         destroy the logger
Flush()                           flush the writer
Name() string                     get logger name
SetName(name string)              set logger name
TimeFormat() string               get time format
SetTimeFormat(format string)      set time format
Level() Level                     get level
SetLevel(level Level)             set level
Format() string                   get the first part of the format
Fargs() []string                  get the second part of the format as slice
SetFormat(format string) error    set format
Writer() io.Writer                get writer
AddWriter(out io.Writer)          add writer
AddWriters(out ...io.Writer)      add multiple writers
SetWriter(out io.Writer)          set writer
SetWriters(out ...io.Writer)      set multiple writers
Sync() bool                       get sync or async
SetSync(sync bool)                set to sync or async
```

### Logging
It supports the following operations for logging, all of which are
thread-safe.
```
Logf(level Level, format string, v ...interface{})
Log(level Level, v ...interface{})
Criticalf(format string, v ...interface{})
Critical(v ...interface{})
Fatalf(format string, v ...interface{})
Fatal(v ...interface{})
Errorf(format string, v ...interface{})
Error(v ...interface{})
Warningf(format string, v ...interface{})
Warning(v ...interface{})
Warnf(format string, v ...interface{})
Warn(v ...interface{})
Infof(format string, v ...interface{})
Info(v ...interface{})
Debugf(format string, v ...interface{})
Debug(v ...interface{})
Notsetf(format string, v ...interface{})
Notset(v ...interface{})
```

## Contributers
In alphabetical order
* Cong Ding
* Xiang Li
* Zifei Tong

## TODO List
1. logging server

2. read configuration from file
