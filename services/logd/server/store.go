package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"unsafe"
)

const (
	banner = `
	________________________________________
	/_/\     /_____/\ /______/\  /_____/\     
	\:\ \    \:::_ \ \\::::__\/__\:::_ \ \    
	 \:\ \    \:\ \ \ \\:\ /____/\\:\ \ \ \   
	  \:\ \____\:\ \ \ \\:\\_  _\/ \:\ \ \ \  
	   \:\/___/\\:\_\ \ \\:\_\ \ \  \:\/.:| | 
		\_____\/ \_____\/ \_____\/   \____/_/
	_________________________________________
	by metroDB 2020
			`
)

type Log struct {
	File       *os.File
	LastOffset []byte
	banner     string
}

func (logdlog *Log) open() {
	logdlog.banner = banner
	fmt.Println(logdlog.banner)
	log.Println(`msg="initializing logger deamon..."`)
	// Add Logger File to Logd
	file, err := os.OpenFile(
		fmt.Sprintf("%slog.db", os.Getenv("LOGD_DB_PATH")),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0600)
	if err != nil {
		log.Fatal(err)
	}
	logdlog.File = file
	log.Println(fmt.Sprintf(`msg="Logging to File: %s"`, logdlog.File.Name()))
}

func (logdlog *Log) close() {
	log.Println(`msg="shutting down logger deamon..."`)
	defer logdlog.File.Close()
}

func (logdlog *Log) append(logValue []byte) []byte {
	var logEntry []byte
	// get size of object in Bytes
	v := int64(unsafe.Sizeof(logValue))
	logValueLength := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(logValueLength, v)

	// add size and value to the entry
	for _, logValuePart := range logValue {
		logEntry = append(logEntry, logValuePart)
	}
	for _, logValueLengthPart := range logValueLength {
		logEntry = append(logEntry, logValueLengthPart)
	}
	// append the data to logfile
	offsetvalue, err := logdlog.File.Write(logEntry)
	if err != nil {
		log.Fatal(err)
	}
	lastOffset, _ := binary.Varint(logdlog.LastOffset)
	offset := make([]byte, binary.MaxVarintLen64)
	binary.PutVarint(offset, int64(lastOffset)+int64(offsetvalue))
	logdlog.LastOffset = offset
	log.Println(fmt.Sprintf(`msg="New Log entry" size="%d" offset="%x"`, v, offset))
	return offset
}

func (logdlog *Log) get(offset []byte) []byte {
	offsetInt, _ := binary.Varint(offset)
	lengthFieldValue := make([]byte, binary.MaxVarintLen64)
	logdlog.File.ReadAt(lengthFieldValue, offsetInt)
	lengthFieldValueInt, _ := binary.Varint(lengthFieldValue)
	returnValue := make([]byte, lengthFieldValueInt)
	logdlog.File.ReadAt(returnValue, offsetInt+lengthFieldValueInt)
	log.Println(fmt.Sprintf(`msg="Log Read" size="%d" offset="%x"`, lengthFieldValueInt, offset))
	return returnValue
}
