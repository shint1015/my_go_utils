package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	ERROR   = "[ERROR]"
	INFO    = "[INFO]"
	WARNING = "[WARNING]"
	DEBUG   = "[DEBUG]"
	FATAL   = "[FATAL]"
)

func LoggingSettings() {
	logFile := GetLogfileNameNum() + "_application.log"
	topPath := GetTopPath()
	filePath := topPath + "logs/" + logFile

	if !FileExistCheck(filePath) {
		if err := createLogFile(filePath); err != nil {
			return
		}
	}

	logfile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}

	//multiLogFile := io.MultiWriter(os.Stdout, logfile)
	multiLogFile := io.MultiWriter(logfile)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func LoggerNew(packageName string) (*log.Logger, error) {
	filePath := getLogfilePath(packageName)
	if !FileExistCheck(filePath) {
		if err := createLogFile(filePath); err != nil {
			return nil, err
		}
	}
	logfile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	//multiLogFile := io.MultiWriter(os.Stdout, logfile)
	multiLogFile := io.MultiWriter(logfile)
	logger := log.New(multiLogFile, "", log.Ldate|log.Ltime)
	return logger, nil
}

func getLogfilePath(packageName string) string {
	logFileName := GetLogfileNameNum() + "_application.log"
	topPath := GetTopPath()
	filePath := topPath + "logs/"
	if packageName == "cmd" {
		filePath += "cmd/"
		logFileName = GetLogfileNameNum() + "_cmd.log"
	} else if packageName == "external" {
		filePath += "external/"
		logFileName = GetLogfileNameNum() + "_external.log"
	} else if packageName == "exchange" {
		filePath += "external/"
		logFileName = GetLogfileNameNum() + "_exchange.log"
	} else if packageName == "kucoin" {
		filePath += "external/"
		logFileName = GetLogfileNameNum() + "_kucoin.log"
	} else if packageName == "coinService" {
		filePath += "internal/"
		logFileName = GetLogfileNameNum() + "_coinService.log"
	} else if packageName == "exchangeService" {
		filePath += "internal/"
		logFileName = GetLogfileNameNum() + "_exchangeService.log"
	}

	filePath += logFileName
	return filePath
}

func GetLogfileNameNum() string {
	return fmt.Sprintf("%04d", GetYear()) + fmt.Sprintf("%02d", GetMonth()) + fmt.Sprintf("%02d", GetMonday())
}

func createLogFile(filePath string) error {
	fp, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Printf("failed to close file: %s", err)
		}
	}()
	return nil
}

func SendLog(customLog *log.Logger, e *CustomError) {

	var msg string
	switch e.errorType {
	case errorTypeERROR:
		msg = " " + ERROR + " " + e.msg
	case errorTypeDEBUG:
		msg = " " + DEBUG + " " + e.msg
	case errorTypeINFO:
		msg = " " + INFO + " " + e.msg
	case errorTypeWARN:
		msg = " " + WARNING + " " + e.msg
	case errorTypeFATAL:
		msg = " " + FATAL + " " + e.msg
	default:
		msg = e.msg
	}
	if customLog == nil {
		log.Printf("[%s:%d] %s: %s", e.file, e.line, e.funcName, msg)
	} else {
		customLog.Printf("[%s:%d] %s: %s", e.file, e.line, e.funcName, msg)
	}
}
