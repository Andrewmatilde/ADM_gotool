package prange

import "math"

func RangeInt(start,end,step int) []int {
	size := int(math.Ceil(float64(end-start)/float64(step)))
	arrange := make([]int, size)
	for i := 0;i < size ;i++  {
		arrange[i] = start + i * step
	}
	return arrange
}
