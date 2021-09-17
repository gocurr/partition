package partition

import "testing"

func TestRanges(t *testing.T) {
	ranges := Ranges(len(s), 1)
	for _, r := range ranges {
		do(s[r.From:r.To])
	}
}

func do(nums []interface{}) {
	//fmt.Println(nums)
}
