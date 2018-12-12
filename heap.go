package main

func (p Pool) Len() int { return len(p) }

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p *Pool) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
	a[i].hid = i
	a[j].hid = j
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	item := x.(*GasStation)
	item.hid = n
	*p = append(*p, item)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	item.hid = -1
	*p = old[0 : n-1]
	return item
}
