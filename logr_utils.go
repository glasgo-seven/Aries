package Aries

import (
	"os"
	"log"
	"time"
)

const logFormat_error		string = "| ERROR       | %v"
const logFormat_alert		string = "| ALERT       | %v"

const defaultTimestampFormat	string = "2006-01-02_15-04-05"

func getTimestamp(_timestampFormat string) string {
	if (_timestampFormat != nil_) {
		return time.Now().Format(_timestampFormat)
	}
	return time.Now().Format(defaultTimestampFormat)
}

func SetupLogger() *os.File {
	timestamp := getTimestamp(nil_)
	f, err := os.OpenFile(
		"logs/log-" + timestamp + ".txt",
		os.O_RDWR | os.O_CREATE | os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatalf(logFormat_error, err)
	}
	// defer f.Close()

	log.SetOutput(f)
	log.Println("LOGGER ACTIVE AT [" + timestamp + "]")

	return f
}


//	Logs errors.
//	If no error found => return false, otherwise => true.
func logErr(format string, err error) bool {
	if err != nil {
		log.Printf(format, err)
		return true
	}
	return false
}