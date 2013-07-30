#go-logging
A high-performance logging library for golang.
* Simple: it supports only necessary operations
* Fast: it is very efficient and under improvement

## Getting Started
### Installation
The step below will download the library source code to
`${GOPATH}/src/github.com/ccding/go-logging`.
```bash
go get github.com/ccding/go-logging
```

Given the source code downloaded, it makes you be able to run the examples,
tests, and benchmarks.
```bash
cd ${GOPATH}/src/github.com/ccding/go-logging
go get
go run example
go test -v -bench .
```

### Example
go-logging is used like any other Go libraries. You can simply use the library
in this way.
```go
import "github.com/ccding/go-logging"
```

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

### Configuration
#### Construction Functions
It has the following functions to create a logger.
```go
// with basic configuration and writing to stdout
SimpleLogger(name string) (*Logger, error)
// with basic configuration and writing to DefaultFileName 'logging.log'
BasicLogger(name string) (*Logger, error)
// with rich configuration and writing to DefaultFileName 'logging.log'
RichLogger(name string) (*Logger, error)
// with detailed configuration and writing to file
FileLogger(name string, level Level, format string, file string, sync bool) (*Logger, error)
// with detailed configuration and writing to a writer
WriterLogger(name string, level Level, format string, out io.Writer, sync bool) (*Logger, error)
```
The meanings of these fields are
```go
name           string        // logger name
level          Level         // record level higher than this will be printed
format         string        // format configuration
out            io.Writer     // writer
startTime      time.Time     // start time of the logger
sync           bool          // use sync or async way to record logs
timeFormat     string        // format for time
```

#### Logging Levels
There are these levels in logging.
```go
CRITICAL     50
FATAL        CRITICAL
ERROR        40
WARNING      30
WARN         WARNING
INFO         20
DEBUG        10
NOTSET       0
```

#### Logging
It supports the following operations for logging. All of these functions are
thread-safe.
```go
(*Logger) Logf(level Level, format string, v ...interface{})
(*Logger) Log(level Level, v ...interface{})
(*Logger) Criticalf(format string, v ...interface{})
(*Logger) Critical(v ...interface{})
(*Logger) Fatalf(format string, v ...interface{})
(*Logger) Fatal(v ...interface{})
(*Logger) Errorf(format string, v ...interface{})
(*Logger) Error(v ...interface{})
(*Logger) Warningf(format string, v ...interface{})
(*Logger) Warning(v ...interface{})
(*Logger) Warnf(format string, v ...interface{})
(*Logger) Warn(v ...interface{})
(*Logger) Infof(format string, v ...interface{})
(*Logger) Info(v ...interface{})
(*Logger) Debugf(format string, v ...interface{})
(*Logger) Debug(v ...interface{})
(*Logger) Notsetf(format string, v ...interface{})
(*Logger) Notset(v ...interface{})
```

#### Getters and Setters
The logger supports the following getter and setter operations, all of which
(except Level and SetLevel) are not thread-safe. If you are calling them in
multiple threads, please be sure locking them properly.
```go
(*Logger) Name() string                    // get name
(*Logger) SetName(name string)             // set name
(*Logger) TimeFormat() string              // get time format
(*Logger) SetTimeFormat(format string)     // set time format
(*Logger) Level() Level                    // get level  [this function is thread safe]
(*Logger) SetLevel(level Level)            // set level  [this function is thread safe]
(*Logger) Format() string                  // get the first part of the format
(*Logger) Fargs() []string                 // get the second part of the format
(*Logger) SetFormat(format string) error   // set format
(*Logger) Writer() io.Writer               // get writer
(*Logger) AddWriter(out io.Writer)         // add writer
(*Logger) AddWriters(out ...io.Writer)     // add multiple writers
(*Logger) SetWriter(out io.Writer)         // set writer
(*Logger) SetWriters(out ...io.Writer)     // set multiple writers
(*Logger) Sync() bool                      // get sync or async
(*Logger) SetSync(sync bool)               // set to sync or async
```

#### Logging Format
The logging format is described by a string, which has two parts separated by
`\n`. The first part describes the format of the log, and the second part
lists all the fields to be shown in the log. In other word, the first part is
the first parameter `format` of `fmt.Printf(format string, v ...interface{})`,
and the second part describes the second parameter `v` of it. It is not
allowed to have `\n` in the first part.  The fields in the second part are
separated by comma `,`, while extra blank spaces are allowed.  An example of
the format string is
```go
const BasicFormat = "%s [%6s] %30s - %s\n name,levelname,time,message"
```
which is the pre-defined `BasicFormat` used by `BasicLogger()` and
`SimpleLogger()`.

It supports the following fields for the second part of the format.
```go
"name"          string     %s      // name of the logger
"seqid"         uint64     %d      // sequence number
"levelno"       int32      %d      // level number
"levelname"     string     %s      // level name
"created"       int64      %d      // starting time of the logger
"nsecs"         int64      %d      // nanosecond of the starting time
"time"          string     %s      // record created time
"timestamp"     int64      %d      // timestamp of record
"rtime"         int64      %d      // relative time since started
"filename"      string     %s      // source filename of the caller
"pathname"      string     %s      // filename with path
"module"        string     %s      // executable filename
"lineno"        int        %d      // line number in source code
"funcname"      string     %s      // function name of the caller
"thread"        int32      %d      // thread id
"process"       int        %d      // process id
"message"       string     %d      // logger message
```

#### Other Operations
It has two other operations to flush the writer and destroy the logger.
```go
(*Logger) Flush()             // flush the writer
(*Logger) Destroy()           // destroy the logger
```

## Contributors
In alphabetical order
* Cong Ding ([ccding][ccding])
* Xiang Li ([xiangli-cmu][xiangli])
* Zifei Tong ([5kg][5kg])
[ccding]: //github.com/ccding
[xiangli]: //github.com/xiangli-cmu
[5kg]: //github.com/5kg

## TODO List
1. logging server

2. read configuration from file
