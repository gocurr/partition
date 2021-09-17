package partition

import (
	"testing"
)

var total = 1000000
var partSize = 1

func TestPartition(t *testing.T) {
	for i := 0; i < total; i++ {
		parts := Partition(s, partSize)
		for _, p := range parts {
			do(p)
		}
	}
}

func TestPartitionx(t *testing.T) {
	for i := 0; i < total; i++ {
		parts := partitionX(s, partSize)
		for _, p := range parts {
			do(p)
		}
	}
}

func TestRanges(t *testing.T) {
	ranges := Ranges(len(s), partSize)
	for _, r := range ranges {
		do(s[r.From:r.To])
	}
}

func do(nums []interface{}) {
	//fmt.Println(nums)
}
