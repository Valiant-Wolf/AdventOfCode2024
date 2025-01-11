package linkedlist

type LinkedList[T any] struct {
	Value      T
	Prev, Next *LinkedList[T]
}

func (l *LinkedList[T]) Delete() (prev, next *LinkedList[T]) {
	prev, next = l.Prev, l.Next

	if prev != nil {
		prev.Next = next
	}

	if next != nil {
		next.Prev = prev
	}
	return
}

func (l *LinkedList[T]) InsertAfter(value T) (new *LinkedList[T]) {
	next := l.Next
	new = &LinkedList[T]{value, l, next}

	l.Next = new
	if next != nil {
		next.Prev = new
	}

	return
}

func (l *LinkedList[T]) InsertBefore(value T) (new *LinkedList[T]) {
	prev := l.Prev
	new = &LinkedList[T]{value, prev, l}

	l.Prev = new
	if prev != nil {
		prev.Next = new
	}

	return
}
