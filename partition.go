package partition

// Partition returns slice of interface slice which is divided by partSize
func Partition(data []interface{}, partSize int) [][]interface{} {
	if partSize < 1 {
		panic("partSize must be greater than 0")
	}

	length := len(data)
	mod := length % partSize
	n := length / partSize
	if mod > 0 {
		n++
	}

	var parts = make([][]interface{}, n)
	var counter = 0
	for i := range parts {
		// last item in parts aka parts[:-1],
		// if mod is bigger than 0,
		// that means: len(data[:-1]) == mod
		if i == n-1 && mod > 0 {
			partSize = mod
		}
		// make sub-slice
		parts[i] = make([]interface{}, partSize)

		// range sub-slice
		for j := range parts[i] {
			parts[i][j] = data[counter]
			counter++
		}
	}

	return parts
}
