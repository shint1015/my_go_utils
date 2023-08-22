package utils

import (
	"errors"
	"path/filepath"
	"runtime"
)

const (
	errorTypeERROR = "ERROR"
	errorTypeWARN  = "WARN"
	errorTypeINFO  = "INFO"
	errorTypeDEBUG = "DEBUG"
	errorTypeFATAL = "FATAL"
)

type CustomError struct {
	errorType string
	msg       string
	Error     error
	file      string
	funcName  string
	line      int
	code      int
}

func (e *CustomError) errorFormat(err error, depth int) *CustomError {
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		file = "???"
		line = 0
	} else {
		file = filepath.Base(file)
	}
	e.line = line
	e.funcName = runtime.FuncForPC(pc).Name()
	e.file = file
	e.msg = err.Error()
	e.Error = err
	return e
}

func Error(errInterface interface{}, depth int, args ...interface{}) *CustomError {
	errorObj := &CustomError{}
	err := convertErrorInterface(errInterface)
	errorObj.errorFormat(err, depth)
	errorObj.errorType = errorTypeERROR
	return errorObj
}

func Warning(errInterface interface{}, depth int) *CustomError {
	errorObj := &CustomError{}
	err := convertErrorInterface(errInterface)
	errorObj.errorFormat(err, depth)
	errorObj.errorType = errorTypeWARN
	return errorObj
}

func Info(errInterface interface{}, depth int) *CustomError {
	errorObj := &CustomError{}
	err := convertErrorInterface(errInterface)
	errorObj.errorFormat(err, depth)
	errorObj.errorType = errorTypeINFO
	return errorObj
}

func (e *CustomError) Debug(errInterface interface{}, depth int) *CustomError {
	errorObj := &CustomError{}
	err := convertErrorInterface(errInterface)
	errorObj.errorFormat(err, depth)
	errorObj.errorType = errorTypeDEBUG
	return errorObj
}

func (e *CustomError) Fatal(errInterface interface{}, depth int) *CustomError {
	errorObj := &CustomError{}
	err := convertErrorInterface(errInterface)
	errorObj.errorFormat(err, depth)
	errorObj.errorType = errorTypeFATAL
	return errorObj
}

func (e *CustomError) Errorf(errInterface interface{}, depth int, format string, args ...interface{}) *CustomError {
	errorObj := &CustomError{}
	err := convertErrorInterface(errInterface)
	errorObj.errorFormat(err, depth)
	errorObj.errorType = errorTypeERROR
	errorObj.msg = format
	return errorObj
}

func convertErrorInterface(errInterface interface{}) error {
	switch errInterface.(type) {
	case error:
		return errInterface.(error)
	case string:
		return errors.New(errInterface.(string))
	default:
		return errors.New("error")
	}
}
