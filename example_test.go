package partition_test

import (
	"fmt"
	"github.com/gocurr/partition"
)

var s = []interface{}{1, 2, 3, 4, 5, 6, 7}

func ExamplePartition() {
	parts := partition.Partition(s, 2)
	for _, p := range parts {
		fmt.Println(p)
	}
	// Output: [1 2]
	// [3 4]
	// [5 6]
	// [7]
}

func ExampleRanges() {
	ranges := partition.Ranges(len(s), 2)
	for _, r := range ranges {
		nums := s[r.From:r.To]
		fmt.Println(nums)
	}
	// Output: [1 2]
	// [3 4]
	// [5 6]
	// [7]
}

func ExampleRangesN() {
	ranges := partition.RangesN(len(s), 5)
	for _, r := range ranges {
		fmt.Println(s[r.From:r.To])
	}

	// Output: [1 2]
	// [3 4]
	// [5 6]
	// [7]
}
