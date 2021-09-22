package partition

import "errors"

type Range struct {
	From, To int
}

var (
	negativeErr    = errors.New("size must be non-negative")
	lessThanOneErr = errors.New("partSize must be greater than 0")
)

// Ranges return the part-range slice, divided by part-size
func Ranges(size, partSize int) ([]Range, error) {
	if size < 0 {
		return nil, negativeErr
	}

	if partSize < 1 {
		return nil, lessThanOneErr
	}

	if partSize >= size {
		return []Range{{From: 0, To: size}}, nil
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

	return ranges, nil
}

// RangesN return the part-range slice, divided by n part
func RangesN(size, n int) ([]Range, error) {
	if n < 1 {
		return nil, lessThanOneErr
	}

	partSize := size / n
	remainder := size % n
	if remainder > 0 {
		partSize++
	}

	return Ranges(size, partSize)
}
