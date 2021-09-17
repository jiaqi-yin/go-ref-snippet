package main

import "fmt"

type Logging interface {
	Info()
	Error()
}

type LogManager struct {
	Logging
}

func NewLogManager(logging Logging) *LogManager {
	return &LogManager{logging}
}

type FileLogging struct{}

func (f *FileLogging) Info() {
	fmt.Println("FileLogging Info")
}

func (f *FileLogging) Error() {
	fmt.Println("FileLogging Error")
}

type DBLogging struct{}

func (db *DBLogging) Info() {
	fmt.Println("DBLogging Info")
}

func (db *DBLogging) Error() {
	fmt.Println("DBLogging Error")
}

func main() {
	fileLogging := &FileLogging{}
	logManager := NewLogManager(fileLogging)
	logManager.Info()
	logManager.Error()

	dbLogging := &DBLogging{}
	logManager.Logging = dbLogging
	logManager.Info()
	logManager.Error()
}
