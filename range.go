package partition

// Range represents a From-To pair
type Range struct {
	From, To int
}

// Ranges returns Range slice which is divided by partSize
func Ranges(size, partSize int) []Range {
	if size < 0 {
		panic("size must be non-negative")
	}

	if partSize < 1 {
		panic("partSize must be greater than 0")
	}

	if partSize >= size {
		return []Range{{From: 0, To: size}}
	}

	remainder := size % partSize
	n := size / partSize
	nn := n
	if remainder > 0 {
		nn++
	}

	var ranges = make([]Range, nn)
	from := 0
	to := partSize
	var i int
	for i = 0; i < n; i++ {
		// fill ranges with range
		ranges[i] = Range{From: from, To: to}
		from, to = to, to+partSize
	}

	// prepare last item in ranges aka ranges[:-1],
	// if remainder is bigger than 0,
	// that means: len(ranges[:-1]) == remainder
	if nn > n {
		ranges[i] = Range{From: from, To: from + remainder}
	}

	return ranges
}

// RangesN returns Range slice which is divided by n
func RangesN(size, n int) []Range {
	if n < 1 {
		panic("n must be greater than 0")
	}

	partSize := size / n
	remainder := size % n
	if remainder > 0 {
		partSize++
	}

	return Ranges(size, partSize)
}
