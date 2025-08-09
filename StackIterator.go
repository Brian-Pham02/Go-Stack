package gostack

import (
	"errors"

	ll "github.com/Brian-Pham02/Go-Linked-List"
)

type StackIterator[T comparable] struct {
	expectedModCount int
	itPtr            *ll.Node[T]
	stack            *ListStack[T]
}

func (ls *ListStack[T]) Iterator() *StackIterator[T] {
	return &StackIterator[T]{
		itPtr:            ls.GetHead(),
		expectedModCount: ls.GetModCount(),
		stack:            ls,
	}
}

func (it *StackIterator[T]) HasNext() bool {
	return it.itPtr != nil
}

func (it *StackIterator[T]) Next() (T, error) {
	var zero T

	if !it.HasNext() {
		return zero, errors.New("NoSuchElementException")
	}

	if it.expectedModCount != it.stack.GetModCount() {
		return zero, errors.New("ConcurrentModificationException")
	}

	value := it.itPtr.Data()
	it.itPtr = it.itPtr.Next()
	return value, nil
}
