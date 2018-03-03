package container

type SliceType []Type

func (s *SliceType) PushBack(x Type) {
	ds := *s
	if s == nil {
		ds = SliceType{x}
		s = &ds
	} else {
		*s = append(ds, x)
	}
}

func (s *SliceType) PushFront(x Type) {
	ds := *s
	if s == nil {
		ds = SliceType{x}
		s = &ds
	} else {
		*s = append(SliceType{x}, ds...)
	}
}

func (s *SliceType) PopBack() Type {
	if s.Empty() {
		return nil
	}
	ds := *s
	f := func() { *s = ds[0 : len(ds)-1] }
	defer f()
	return s.Back()
}

func (s *SliceType) PopFront() Type {
	if s.Empty() {
		return nil
	}
	ds := *s
	f := func() { *s = ds[1:] }
	defer f()
	return s.Front()
}

func (s *SliceType) Front() Type {
	if s.Empty() {
		return nil
	}
	ds := *s
	return ds[0]
}

func (s *SliceType) Back() Type {
	if s.Empty() {
		return nil
	}
	ds := *s
	return ds[len(ds)-1]
}

func (s *SliceType) Empty() bool {
	return s.Size() == 0
}

func (s *SliceType) Size() uint {
	return uint(len(*s))
}
