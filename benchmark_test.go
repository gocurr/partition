package partition

import "testing"

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
		_, _ = Ranges(len(s), 1)
	}
}
