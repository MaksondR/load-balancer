package main

import (
	"math/rand"
	"strconv"
	"time"
)

type Refill struct {
	car      string
	tank     int
	fullTank chan int
}

func RefillRequest(req chan Refill) {
	fullTank := make(chan int)

	for {
		time.Sleep(time.Millisecond)
		req <- Refill{strconv.Itoa(rand.Intn(100)) + " car", rand.Intn(10), fullTank}
		<-fullTank
	}
}
