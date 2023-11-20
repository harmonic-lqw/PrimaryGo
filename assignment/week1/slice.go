package week1

const (
	DEFAULT_REDUCE_CAPACITY_PERCENT = 0.25
)

type order interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint |
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
	comparable
}

func Insert[T order](idx int, val T, vals []T) []T {
	if idx < 0 || idx > len(vals) {
		panic("idx not illegal")
	}
	vals = append(vals, val)
	if idx == len(vals) {
		return vals
	}
	for i := len(vals) - 1; i > idx; i-- {
		if i-1 >= 0 {
			vals[i] = vals[i-1]
		}
	}
	vals[idx] = val
	return vals
}

// Delete week1 assignment
func Delete[T order](idx int, vals []T) []T {
	if idx < 0 || idx >= len(vals) {
		panic("idx not illegal")
	}
	vals = append(vals[:idx], vals[idx+1:]...)

	// Reduce capacity when slice length is less than 0.25 of capacity.
	// Reduced to 2 times the current length
	lenVals := len(vals)
	if float64(lenVals)/float64(cap(vals)) <= DEFAULT_REDUCE_CAPACITY_PERCENT {
		newVals := make([]T, 0, 2*lenVals)
		// newVals = vals
		newVals = append(newVals, vals...)
		return newVals
	}
	return vals
}
