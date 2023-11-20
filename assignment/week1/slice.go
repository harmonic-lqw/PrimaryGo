package week1

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
