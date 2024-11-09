package main

import (
	"log"
	"os"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

func init() {
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(format string, args ...interface{}) {
	infoLog.Printf(format, args...)
}

func Error(format string, args ...interface{}) {
	errorLog.Printf(format, args...)
}
