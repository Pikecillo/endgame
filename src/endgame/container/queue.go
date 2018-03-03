package container

type Queue interface {
	Push(x Type)
	Pop() Type
	Front() Type
	Back() Type
	IsEmpty() bool
	Size() uint
}

type SliceQueue struct {
	slice SliceType
}

func (q *SliceQueue) Push(x Type) {
	q.slice.PushBack(x)
}

func (q *SliceQueue) Pop() Type {
	return q.slice.PopFront()
}

func (q *SliceQueue) Front() Type {
	return q.slice.Front()
}

func (q *SliceQueue) Back() Type {
	return q.slice.Back()
}

func (q *SliceQueue) Empty() bool {
	return q.slice.Empty()
}

func (q *SliceQueue) Size() uint {
	return q.slice.Size()
}
