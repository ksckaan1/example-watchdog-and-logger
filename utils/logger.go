package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func PrintLog(text string) {
	currentTime := time.Now()
	date := currentTime.Format("2006-02-01")
	hour, minute, second := currentTime.Hour(), currentTime.Minute(), currentTime.Second()
	logFileName := fmt.Sprintf("%s.log", date)

	logFile := checkLogFile(logFileName)

	defer logFile.Close()

	logText := fmt.Sprintf("%d:%d:%d log: %s\n", hour, minute, second, text)
	fmt.Printf(logText)
	_, err := logFile.WriteString(logText)
	if err != nil {
		log.Println(err)
	}

}

func checkLogFile(logPath string) *os.File {
	var logFile *os.File

	_, err := os.Stat(logPath)

	if err == nil {
		logFile, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		logFile, err = os.Create(logPath)
	}
	return logFile
}
