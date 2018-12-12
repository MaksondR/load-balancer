package main

const fillRequests = 10
const stations = 5

func main() {
	gas := make(chan Refill)
	for i := 0; i < fillRequests; i++ {
		go RefillRequest(gas)
	}
	InitBalancer().balance(gas)
}
