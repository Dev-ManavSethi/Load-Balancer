package main

import (
	"net/http"
	"strconv"
	"sync"
	"time"
)

var (
	counter    int  = 0
	counterPtr *int = &counter
	ticker     *time.Ticker
	mutex      sync.Mutex
)

func main() {

	ticker = time.NewTicker(1000 * time.Millisecond)

	LogInfo("Spinning up load balancer 2...")

	LogInfo("Reading Config.yml...")
	proxy, err := ReadConfig()
	if err != nil {
		LogErr("An error occurred while trying to parse config.yml")
		LogErrAndCrash(err.Error())
	}

	go func() {

		for {
			<-ticker.C
			mutex.Lock()
			*counterPtr = 0
			mutex.Unlock()
		}

	}()

	http.HandleFunc("/", proxy.handler)

	LogInfo("Listening to requests on port: " + strconv.Itoa(proxy.Port))

	err = http.ListenAndServe(":"+strconv.Itoa(proxy.Port), nil)
	if err != nil {
		LogErr("Failed to bind to port " + strconv.Itoa(proxy.Port))
		LogErrAndCrash("Make sure the port is available and not reserved")
	}

}
