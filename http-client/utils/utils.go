package main

func Total(values []int64) int64 {
	var total int64
	for _, v := range values {
		total += v
	}
	return total
}

func Average(values []int64) int64 {
	return Total(values) / int64(len(values))
}

func Dispersion(values []int64) int64 {
	var doubleTotal int64
	for _, v := range values {
		doubleTotal += v * v
	}
	doubleAverage := doubleTotal / int64(len(values))
	return doubleAverage - (Average(values) * Average(values))
}
