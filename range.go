package partition

type Range struct {
	From, To int
}

func Ranges(size, partSize int) []Range {
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
