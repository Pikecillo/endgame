package container

type Stack interface {
	Push(x Type)
	Pop() Type
	Top() Type
	IsEmpty() bool
	Size() uint
}

type SliceStack struct {
	slice SliceType
}

func (s *SliceStack) Push(x Type) {
	s.slice.PushBack(x)
}

func (s *SliceStack) Pop() Type {
	return s.slice.PopBack()
}

func (s *SliceStack) Top() Type {
	return s.slice.Back()
}

func (s *SliceStack) Empty() bool {
	return s.slice.Empty()
}

func (s *SliceStack) Size() uint {
	return s.slice.Size()
}
