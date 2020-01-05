package main

import (
	"fmt"
	"log"
	"os"

	logd "github.com/nk-designz/metroDB/services/logd/client"
)

type Logds struct {
	logds []*logd.Logd
}

func (logdS *Logds) Init() {
	logds := make([]*logd.Logd, len(os.Args))
	for i := 0; i <= len(os.Args); i++ {
		logds[i] = logd.New(os.Args[i])
		logds[i].Connect()
		log.Println(fmt.Sprintf(`msq="connecting logd" address="%s"`, os.Args[i]))
	}
	logdS.logds = logds
}

func (logdS *Logds) Close() {
	for i := 0; i <= len(logdS.logds); i++ {
		logdS.logds[i].Close()
		log.Println(`msg="closing logd"`)
	}
}
