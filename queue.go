package queue

type Queue[V any] interface {
	Push(value V)
	Pop() (value V, ok bool)
	Peek() (value V, ok bool)
	Len() int
}
