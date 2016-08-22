package priorityq

import ()

type item struct {
	value interface{}
	cost  int32
	index int32
}

type BinaryHeap struct {
	items   []*item
	indices map[interface{}]int32
	len     int32
}
