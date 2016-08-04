package priorityq

import (
	pq "github.com/rmahidhar/priorityq"
)

type item struct {
	value interface{}
	cost int
	index int
}

type BinaryHeap struct {
	items []*item
	indices map[interface{}]int
	len int
}
