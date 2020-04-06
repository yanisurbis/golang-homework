package hw04_lru_cache //nolint:golint,stylecheck

type listItem struct {
	value interface{}
	next  *listItem
	prev  *listItem
}

type list struct {
	head   *listItem
	tail   *listItem
	length int
}

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *listItem {
	return l.head
}

func (l list) Back() *listItem {
	return l.tail
}

func (l *list) AddFirstElement(v interface{}) *listItem {
	newElm := &listItem{
		value: v,
		next:  nil,
		prev:  nil,
	}
	l.head = newElm
	l.tail = newElm
	l.length = 1

	return newElm
}

func (l *list) PushFront(v interface{}) *listItem {
	if l.length == 0 {
		return l.AddFirstElement(v)
	}

	head := l.head
	newElm := &listItem{
		value: v,
		next:  head,
		prev:  nil,
	}
	head.prev = newElm
	l.head = newElm
	l.length++

	return newElm
}

func (l *list) PushBack(v interface{}) *listItem {
	if l.length == 0 {
		return l.AddFirstElement(v)
	}

	tail := l.tail
	newElm := &listItem{
		value: v,
		next:  nil,
		prev:  tail,
	}
	tail.next = newElm
	l.tail = newElm
	l.length++

	return newElm
}

func (l *list) Remove(i *listItem) {
	if i == nil || l.length == 0 {
		return
	}

	if i.prev != nil && i.next != nil {
		i.prev.next = i.next
		i.next.prev = i.prev
		l.length--
		return
	}

	if i.prev == nil && i.next != nil {
		l.head = i.next
		l.head.prev = nil
		l.length--
		return
	}

	if i.prev != nil && i.next == nil {
		l.tail = i.prev
		l.tail.next = nil
		l.length--
		return
	}

	if i.prev == nil && i.next == nil {
		l.head = nil
		l.tail = nil
		l.length = 0
		return
	}
}

func (l *list) MoveToFront(i *listItem) {
	if i == nil {
		return
	}
	l.Remove(i)
	l.PushFront(i.value)
}

func NewList() List {
	return &list{}
}
