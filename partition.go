package partition

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
		// make subslice
		parts[i] = make([]interface{}, partSize)

		// range subslice
		for j := range parts[i] {
			parts[i][j] = data[counter]
			counter++
		}
	}

	return parts
}

func partitionX(data []interface{}, partSize int) [][]interface{} {
	if partSize < 1 {
		panic("partSize must be greater than 0")
	}

	var counter int
	var part []interface{}
	var parts [][]interface{}
	for idx, d := range data {
		if counter == 0 {
			part = make([]interface{}, 0)
		}
		part = append(part, d)
		counter++
		// end of subslice
		if counter == partSize || idx+1 == len(data) {
			counter = 0
			parts = append(parts, part)
		}
	}
	return parts
}
