package partition

import (
	"fmt"
	"testing"
)

var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestRanges(t *testing.T) {
	ranges := Ranges(len(nums), 1)
	for _, r := range ranges {
		do(nums[r.From:r.To])
	}
}

func do(nums []int) {
	fmt.Println(nums)
}
