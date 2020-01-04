package main

import (
	"fmt"
	"math/rand"

	logd "github.com/nk-designz/metroDB/services/logd/client"
)

func main() {
	logdServer := logd.New("127.0.0.1")
	logdServer.Connect()
	offset := logdServer.Append([]byte("larum ipsum dolor sid ametlarum ipsum dolor sid ametlarum ipsum dolor sid ametlarum ipsum dolor sid ametlarum ipsum dolor sid ametlarum ipsum dolor sid ametlarum ipsum dolor sid ametlarum ipsum dolor sid ametlarum ipsum dolor sid amet"))
	entry := logdServer.Get(offset)
	fmt.Println(
		fmt.Sprintf(
			"%x : %s",
			offset,
			entry))

	offset1 := logdServer.Append(
		[]byte(
			fmt.Sprintf("%d", rand.Int63())))
	entry = logdServer.Get(offset1)
	fmt.Println(
		fmt.Sprintf(
			"%x : %s",
			offset1,
			entry))
	entry = logdServer.Get(offset)
	fmt.Println(
		fmt.Sprintf(
			"%x : %s",
			offset,
			entry))
	logdServer.Close()
}
