package dtls

import (
	"fmt"
	"log"
	"time"
)

const (
	LogLevelError string = "error"
	LogLevelWarn  string = "warn"
	LogLevelInfo  string = "info"
	LogLevelDebug string = "debug"
)

type LogFunc func(ts time.Time, level string, peer *Peer, err error, msg string)

var logFunc LogFunc = defaultLogFunc
var logLevel int = 0

func SetLogFunc(lf LogFunc) {
	logFunc = lf
}

func SetLogLevel(level string) {
	switch level {
	case LogLevelError:
		logLevel = 1
	case LogLevelWarn:
		logLevel = 2
	case LogLevelInfo:
		logLevel = 3
	case LogLevelDebug:
		logLevel = 4
	default:
		logLevel = 0
	}
}

func defaultLogFunc(ts time.Time, level string, peer *Peer, err error, msg string) {
	if err != nil {
		log.Printf(" [" + level + "] [" + peer.RemoteAddr() + "] " + msg + "(err: " + err.Error() + ")")
	} else {
		log.Printf(" [" + level + "] [" + peer.RemoteAddr() + "] " + msg)
	}
}

func logError(peer *Peer, err error, f string, args ...interface{}) {
	if logLevel < 1 {
		return
	}
	logFunc(time.Now(), LogLevelError, peer, err, fmt.Sprintf(f, args...))
}

func logWarn(peer *Peer, err error, f string, args ...interface{}) {
	if logLevel < 2 {
		return
	}
	logFunc(time.Now(), LogLevelWarn, peer, err, fmt.Sprintf(f, args...))
}

func logInfo(peer *Peer, f string, args ...interface{}) {
	if logLevel < 3 {
		return
	}
	logFunc(time.Now(), LogLevelInfo, peer, nil, fmt.Sprintf(f, args...))
}

func logDebug(peer *Peer, f string, args ...interface{}) {
	if logLevel < 4 {
		return
	}
	logFunc(time.Now(), LogLevelDebug, peer, nil, fmt.Sprintf(f, args...))
}
