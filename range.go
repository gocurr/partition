package partition

type Range struct {
	From, To int
}

func Ranges(size, partSize int) []Range {
	if partSize < 1 {
		panic("partSize must be greater than 0")
	}

	if partSize > size {
		return []Range{{From: 0, To: size}}
	}

	remainder := size % partSize
	n := size / partSize
	if remainder > 0 {
		n++
	}

	var ranges = make([]Range, n)
	from := 0
	to := partSize
	for i := range ranges {
		// fill ranges with range
		ranges[i] = Range{From: from, To: to}

		// prepare last item in ranges aka ranges[:-1],
		// if remainder is bigger than 0,
		// that means: len(data[:-1]) == remainder
		if i == n-2 && remainder > 0 {
			partSize = remainder
		}
		from, to = to, to+partSize
	}

	return ranges
}
