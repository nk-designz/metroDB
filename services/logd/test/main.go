package main

import (
	logd "github.com/nk-designz/metroDB/services/logd/client"
)

func main() {
	logdServer := logd.New("127.0.0.1")
	logdServer.Connect()
	logd := logdServer.GetClient()
}
