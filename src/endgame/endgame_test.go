package endgame

import "testing"

func TestPawnMoves(t *testing.T) {
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