package container

import "testing"

func TestSliceQueue_SizeAfterPush(t *testing.T) {
	cases := []struct {
		input, expected uint
	}{
		{1, 1}, {4, 2}, {5, 3},
	}

	var q SliceQueue

	for _, test_case := range cases {
		q.Push(test_case.input)
		result := q.Size()
		if result != test_case.expected {
			t.Errorf("Size after pushing %d should be %d instead of %d",
				test_case.input, test_case.expected, result)
		}
	}
}

func TestSliceQueue_SizeAfterPop(t *testing.T) {
	var q SliceQueue

	elements := []uint{12, 5, 78, 9, 10, 20}

	for _, e := range elements {
		q.Push(e)
	}

	sizes := []uint{5, 4, 3, 2, 1, 0}

	for i, size := range sizes {
		popped := q.Pop()
		result := q.Size()

		expected := elements[i]
		if popped != expected {
			t.Errorf("Popped value should have been %d instead of %d",
				expected, popped)
		}

		if result != size {
			t.Errorf("Size after popping %d should be %d instead of %d",
				popped, size, result)
		}
	}

	result := q.Pop()
	if result != nil {
		t.Errorf("Popping empty queue should return nil")
	}
}
