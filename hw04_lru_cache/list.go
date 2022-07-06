package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	Clear()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Size int
	Head *ListItem
	Tail *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.Size
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.Size == 0 {
		l.Head = newItem
		l.Tail = newItem
	} else {
		newItem.Next = l.Head
		l.Head.Prev = newItem
		l.Head = newItem
	}

	l.Size++

	return l.Head
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.Size == 0 {
		l.Head = newItem
		l.Tail = newItem
	} else {
		newItem.Prev = l.Tail
		l.Tail.Next = newItem
		l.Tail = newItem
	}

	l.Size++

	return l.Tail
}

func (l *list) Remove(i *ListItem) {
	if l.Size == 0 {
		return
	}

	l.Size--
	switch {
	case i.Next == nil && i.Prev == nil:
		l.Head = nil
		l.Tail = nil
		return
	case i.Next == nil:
		i.Prev.Next = nil
		l.Tail = i.Prev
	case i.Prev == nil:
		i.Next.Prev = nil
		l.Head = i.Next
	default:
		next := i.Next
		prev := i.Prev
		i.Next.Prev = prev
		i.Prev.Next = next
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) Clear() {
	l.Size = 0
	l.Head = nil
	l.Tail = nil
}
