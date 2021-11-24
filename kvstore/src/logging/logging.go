package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

var logger *logg
var logqueue = make(chan msg, 10000)

const logDefaultLevel int = 3

// Cancelled, Stopped, Error const
const (
	Cancelled string = "Cancelled"
	Stopped   string = "Stopped"
	Error     string = "Error"
	Input     string = "Input"
	Output    string = "Output"
	Payload   string = "Payload"
)

// Logg struct
type logg struct {
	*log.Logger
	level int
	close chan struct{}
}

// Msg like e.g. <when><where><what><supporting data>
type msg struct {
	id     string
	action string
	key    string
	msg    string
	fields Field
}

// Field type
type Field [][]interface{}

// Msg func
func Msg(id string, action string, key string, message string) {
	logToQueue(&msg{id: id, action: action, key: key, msg: message})
}

// Msgf func
func Msgf(id string, action string, key string, format string, v ...interface{}) {
	Msg(id, action, key, fmt.Sprintf(format, v...))
}

// MsgFields func
func MsgFields(id string, action string, key string, fields Field) {
	logToQueue(&msg{id: id, action: action, key: key, fields: fields})
}

// Fatal func
func Fatal(id string, action string, key string, message string) {
	Msg(id, action, key, message)
	Close()
	os.Exit(1)
}

// Fatalf func
func Fatalf(id string, action string, key string, format string, v ...interface{}) {
	Fatal(id, action, key, fmt.Sprintf(format, v...))
}

// Log func
func logToQueue(msg *msg) {
	select {
	case <-logger.close:
	default:
		logqueue <- *msg
	}
}

// Close all goroutines
func Close() {
	select {
	case <-logger.close:
	default:
		close(logger.close)
	}
}

// Process logs using alike dispacher pattern
func processLogs() {
	go func() {
		defer close(logqueue)
		for log := range logqueue {
			select {
			case <-logger.close:
			default:
				go parseLog(log, logger.level)
			}
		}
	}()
}

func parseLog(log msg, loglevel int) {

	if log.key == Cancelled || log.key == Stopped || log.key == Error {
		filterLog(1, loglevel, "%s::%s::%s::%s", log.id, log.action, log.key, log.msg)
	} else if log.key == Input || log.key == Output {
		if log.fields != nil && len(log.fields) > 0 {
			mField, _ := json.Marshal(log.fields)
			filterLog(2, loglevel, "%s::%s::%s", log.id, log.action, mField)
		}
	} else {
		filterLog(0, loglevel, "%s::%s::%s::%s", log.id, log.action, log.key, log.msg)
	}
}

func filterLog(level int, loglevel int, format string, v ...interface{}) {
	if level <= loglevel {
		logger.Printf(format, v...)
	}
}

// UUID kind of doesn't conform to RFC 4122
func UUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "UUID-Error"
	}
	return fmt.Sprintf("%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10])
}

// CreateLogger startup func
func CreateLogger(args []string) {
	if logger == nil {

		var argLevel int = logDefaultLevel
		var wrt io.Writer

		if args != nil && len(args) > 0 {
			level, err := strconv.Atoi(args[0])
			if err == nil {
				argLevel = level
			}
			wrt = io.MultiWriter(os.Stdout, os.Stderr)
		} else {
			wrt = io.MultiWriter(os.Stdout)
		}

		// Logger encapsulation
		logger = &logg{log.New(wrt, "", log.Ldate+log.Lmicroseconds),
			argLevel,
			make(chan struct{})}

		// Start dispatcher of logs
		processLogs()

		Msgf(
			UUID(),
			"CreateLogger",
			"Logging",
			"log level %d, log address %+v", argLevel, &logger)
	}
}
