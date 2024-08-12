package Aries

import (
	"os"
	"log"
	"time"
	"strings"
)

var LogTags = map[string]string {}
const logTag_BaseFatal	= "FATAL"	// App can no longer proceed to work
const logTag_BaseError	= "ERROR"	// Functionality of the app is at risk
const logTag_BaseAlert	= "ALERT"	// Important info to be taken action on
const logTag_BaseNotify	= "NOTIFY"	// Non-important info to be aware of

var logTagSize		int = 12
var minLogTagSize	int = logTagSize

func SetLogrTagSize(_logTagSize int) {
	logTagSize = _logTagSize
}

func TagToFormat(_tag string) string {
	var tagSize int = 0
	if minLogTagSize > logTagSize {
		tagSize = minLogTagSize
	}
	return "| " + _tag + strings.Repeat(" ", tagSize - len(_tag)) + " | %v"
}

func CheckTagSizes() {
	var minTagSize = minLogTagSize
	for key := range LogTags {
		if len(key) > minTagSize {
			minTagSize = len(key)
		}
	}
	if minTagSize != minLogTagSize {
		log.Printf(TagToFormat("ALERT"), "One or more LogTags are longer than a set logTagSize")
	}
}


const DefaultTimestampFormat	string = "2006-01-02_15-04-05"

func GetTimestamp(_timestampFormat string) string {
	if (_timestampFormat != nil_) {
		return time.Now().Format(_timestampFormat)
	}
	return time.Now().Format(DefaultTimestampFormat)
}


var LogFile *os.File = nil

func SetupLogger(_toFile bool, _addBaseTags bool) {
	if _toFile {
		timestamp := GetTimestamp(nil_)
		f, err := os.OpenFile(
			"logs/log-" + timestamp + ".txt",
			os.O_RDWR | os.O_CREATE | os.O_APPEND,
			0666,
		)
		if err != nil {
			log.Fatalf(TagToFormat("FATAL"), err)
		}
		// defer f.Close()

		log.SetOutput(f)
		log.Println("LOGGER ACTIVE AT [" + timestamp + "]")

		LogFile = f
	}
	if _addBaseTags {
		LogTags["BaseFatal"]	= logTag_BaseFatal
		LogTags["BaseError"]	= logTag_BaseError
		LogTags["BaseAlert"]	= logTag_BaseAlert
		LogTags["BaseNotify"]	= logTag_BaseNotify
	}
}

func CloseLogger() {
	LogFile.Close()
}


//	Logs errors.
//	If no error found => return false, otherwise => true.
func LogErr(tag string, err error) bool {
	var format = "| " + tag + " | %v"
	if err != nil {
		log.Printf(format, err)
		return true
	}
	return false
}