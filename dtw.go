package dynamic_time_warping

import "math"

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func DynamicTimeWarping[NUMERIC Numeric](a, b []NUMERIC) uint64 {
	matrix := makeMatrix(len(a)+1, len(b)+1)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			matrix[i][j] = math.MaxUint64
		}
	}
	matrix[0][0] = 0

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			changeMin(&matrix[i+1][j+1], matrix[i][j])
			changeMin(&matrix[i+1][j+1], matrix[i+1][j])
			changeMin(&matrix[i+1][j+1], matrix[i][j+1])
			matrix[i+1][j+1] += distance(a[i], b[j])
		}
	}

	return matrix[len(a)][len(b)]
}

func distance[NUMERIC Numeric](a, b NUMERIC) uint64 {
	return uint64((a - b) * (a - b))
}

func changeMin[NUMERIC Numeric](target *NUMERIC, value NUMERIC) {
	if *target > value {
		*target = value
	}
}

func makeMatrix(n, m int) [][]uint64 {
	matrix := make([][]uint64, n)
	rows := make([]uint64, n*m)
	for i := 0; i < n; i++ {
		matrix[i] = rows[i*m : (i+1)*m]
	}
	return matrix
}
