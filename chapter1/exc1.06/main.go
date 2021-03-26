package main

import (
	"fmt"
    "time"
)

func getConfig () (bool, string, time.Time) {
    return false, "info", time.Now()
}

func main() {
    Debug, Loglevel, startUpTime := getConfig()

    fmt.Println(Debug, Loglevel, startUpTime)
}
