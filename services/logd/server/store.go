package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
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
	LastOffset int64
	banner     string
}

func (logdlog *Log) open() {
	logdlog.banner = banner
	fmt.Println(logdlog.banner)
	log.Println(`msg="initializing logger deamon..."`)
	// Add Logger File to Logd
	file, err := os.OpenFile(
		fmt.Sprintf("%slog.db", os.Getenv("LOGD_DB_PATH")),
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
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

func (logdlog *Log) append(logValue []byte) int64 {
	var logEntry []byte
	// get size of object in Bytes
	v := uint64(len(logValue))
	logValueLength := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(logValueLength, v)
	// add size and value to the entry
	logEntry = append(logEntry, logValue...)
	logEntry = append(logEntry, logValueLength...)
	// append the data to logfile
	_, err := logdlog.File.Write(logEntry)
	if err != nil {
		log.Fatal(err)
	}
	fileInfo, err := logdlog.File.Stat()
	if err != nil {
		log.Fatal(err)
	}
	offset := fileInfo.Size()
	log.Println(fmt.Sprintf(`msg="New Log entry" size="%d" offset="%d"`, v, offset))
	logdlog.LastOffset = offset
	return offset
}

func (logdlog *Log) get(offset ...int64) []byte {
	if len(offset) == 0 {
		offset[0] = logdlog.LastOffset
	}
	lengthFieldValue := make(
		[]byte,
		binary.MaxVarintLen64)
	logdlog.File.ReadAt(
		lengthFieldValue,
		(offset[0] - binary.MaxVarintLen64))
	lengthFieldValueInt := binary.BigEndian.Uint64(lengthFieldValue)
	returnValue := make(
		[]byte,
		lengthFieldValueInt)
	logdlog.File.ReadAt(
		returnValue,
		(offset[0] - (binary.MaxVarintLen64 + int64(lengthFieldValueInt))))
	log.Println(fmt.Sprintf(`msg="Log Read" size="%d" offset="%d"`, lengthFieldValueInt, offset[0]))
	return returnValue
}
