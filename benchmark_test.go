package partition

import "testing"

var s = []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func BenchmarkPartition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Partition(s, 1)
	}
}

func BenchmarkPartitionx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partitionX(s, 1)
	}
}

func BenchmarkRanges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Ranges(len(s), 1)
	}
}
