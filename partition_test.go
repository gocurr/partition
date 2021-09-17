package partition

import "testing"

func Test_Partition(t *testing.T) {
	var s = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	partSlice := partition(s, 5)
	for _, p := range partSlice {
		t.Logf("%v", p)
	}
}
