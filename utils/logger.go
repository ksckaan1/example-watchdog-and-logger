package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func PrintLog(text string) {
	year, month, day := time.Now().Date()
	logFileName := fmt.Sprintf("%v-%v-%v.log", year, month, day)

	logFile := checkLogFile(logFileName)

	defer logFile.Close()

	mw := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(mw)

	log.Println(text)
}

func checkLogFile(logPath string) *os.File {
	var logFile *os.File

	_, err := os.Stat(logPath)

	if err == nil {
		logFile, err = os.OpenFile(logPath, os.O_RDWR, 0644)
	} else {
		logFile, err = os.Create(logPath)
	}
	return logFile
}
