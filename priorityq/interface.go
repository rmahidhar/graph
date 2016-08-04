package priorityq

type Interface interface {
	Push(item interface{}, cost int)
	Pop() (interface{}, int)
	Update(item interface{}, cost int)
	Len() int
	Comparator func(int, int) bool
}
