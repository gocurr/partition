package partition

import (
	"fmt"
	"testing"
)

var partSize = 6

func TestPartition(t *testing.T) {
	parts := Partition(s, partSize)
	for _, p := range parts {
		do(p)
	}
}

func TestPartitionx(t *testing.T) {
	parts := partitionX(s, partSize)
	for _, p := range parts {
		do(p)
	}
}

func TestRanges(t *testing.T) {
	ranges := Ranges(len(s), partSize)
	for _, r := range ranges {
		do(s[r.From:r.To])
	}
}

func do(nums []interface{}) {
	fmt.Println(nums)
}
