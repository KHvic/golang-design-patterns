package iterator

import (
	"container/list"
	"fmt"
)

/*
   Type: Behavioral
   Purpose: Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.
   Additional: - a uniform interface for traversing different aggregate structures (don't have to care about underlying data structure)
*/

// Iterators one for list, one for slice
// Hide underlying Aggregate/Collections data structures
type Iterator interface {
	hasNext() bool
	next() interface{}
}

// SliceIterator
type SliceIterator struct {
	slice []int
	idx   int
}

func (i *SliceIterator) hasNext() bool {
	return i.idx < len(i.slice)
}
func (i *SliceIterator) next() interface{} {
	defer func() {
		i.idx++
	}()
	return i.slice[i.idx]
}

// ListIterator
type ListIterator struct {
	cur *list.Element
}

func (i *ListIterator) hasNext() bool {
	return i.cur != nil
}
func (i *ListIterator) next() interface{} {
	defer func() {
		i.cur = i.cur.Next()
	}()
	return i.cur.Value
}

// Aggregates (Collections) with actual data structures
type Aggregate interface {
	createIterator() Iterator
}

type ListAggregate struct {
	list *list.List
}

func (a *ListAggregate) createIterator() Iterator {
	return &ListIterator{cur: a.list.Front()}
}

type SliceAggregate struct {
	slice []int
}

func (a *SliceAggregate) createIterator() Iterator {
	return &SliceIterator{slice: a.slice}
}

func main() {
	a := &ListAggregate{list.New()}
	a.list.PushBack(123)
	a.list.PushBack(456)
	a.list.PushBack(789)
	for it := a.createIterator(); it.hasNext(); {
		fmt.Println(it.next())
	}

	b := &SliceAggregate{[]int{1, 2, 3, 4, 5, 6, 7}}
	for it := b.createIterator(); it.hasNext(); {
		fmt.Println(it.next())
	}
}
