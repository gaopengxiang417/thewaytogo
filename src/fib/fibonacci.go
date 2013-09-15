package fib

const his []int64 = make([]int, 20)

func Fab_nacci(i int) int {
	if i <= 1 {
		return 1
	}
	if his[i] != 0 {
		return his[i]
	}

	return fab_nacci(i-1) + fab_nacci(i-2)

}
