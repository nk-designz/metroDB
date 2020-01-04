package main

import (
	"fmt"

	logd "github.com/nk-designz/metroDB/services/logd/client"
)

func main() {
	logdServer := logd.New("127.0.0.1")
	logdServer.Connect()
	offset := logdServer.Append([]byte("This is a test"))
	entry := logdServer.Get()
	fmt.Println(
		fmt.Sprintf(
			"%x : %x",
			offset,
			entry))

	offset1 := logdServer.Append([]byte("This is a  second test"))
	entry = logdServer.Get(offset1)
	fmt.Println(
		fmt.Sprintf(
			"%x : %x",
			offset1,
			entry))
	entry = logdServer.Get(offset)
	fmt.Println(
		fmt.Sprintf(
			"%x : %x",
			offset,
			entry))
}
