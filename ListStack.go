package gostack

import (
	ll "github.com/Brian-Pham02/Go-Linked-List"
)

type ListStack[T comparable] struct {
	list *ll.LinkedList[T]
}

func InitStack[T comparable]() *ListStack[T] {
	return &ListStack[T]{
		list: ll.NewList[T](),
	}
}

func (ls *ListStack[T]) IsEmpty() bool {
	return ls.list.IsEmpty()
}

func (ls *ListStack[T]) Count() int {
	return ls.list.Size()
}

func (ls *ListStack[T]) Push(data T) {
	ls.list.AddHead(data)
}

func (ls *ListStack[T]) Pop() T {
	if ls.IsEmpty() {
		panic("Empty Stack")
	}

	return ls.list.RemoveHead()
}

func (ls *ListStack[T]) Peek() T {
	if ls.IsEmpty() {
		panic("Empty Stack")
	}

	return ls.list.Peek()
}

func (ls *ListStack[T]) String() string {
	return ls.list.String()
}
