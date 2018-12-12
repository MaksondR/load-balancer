package main

import (
	"container/heap"
)

type Pool []*GasStation

type Balancer struct {
	pool     Pool
	complete chan *GasStation
}

func InitBalancer() *Balancer {
	complete := make(chan *GasStation, 5)

	b := &Balancer{make(Pool, 0, stations), complete}
	for i := 0; i < 5; i++ {
		g := &GasStation{fill: make(chan Refill, fillRequests)}

		heap.Push(&b.pool, g)
		go g.gasFill(b.complete)
	}
	return b
}

func (b *Balancer) balance(req chan Refill) {
	for {
		select {
		case fillReq := <-req:
			b.dispatch(fillReq)
		case g := <-b.complete:
			b.completed(g)
		}
		// for _, g := range b.pool {
		// 	fmt.Printf(strconv.Itoa(g.pending))
		// }
	}
}

func (b *Balancer) dispatch(req Refill) {
	g := heap.Pop(&b.pool).(*GasStation)
	g.fill <- req
	g.pending++
	heap.Push(&b.pool, g)
}

func (b *Balancer) completed(g *GasStation) {
	g.pending--
	heap.Remove(&b.pool, g.hid)
	heap.Push(&b.pool, g)
}
