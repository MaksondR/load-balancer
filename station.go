package main

import (
	"math/rand"
)

type GasStation struct {
	hid     int
	fill    chan Refill
	pending int
}

func (g *GasStation) gasFill(complete chan *GasStation) {

	for {
		req := <-g.fill
		req.fullTank <- req.tank * rand.Intn(5)
		complete <- g
	}
}
