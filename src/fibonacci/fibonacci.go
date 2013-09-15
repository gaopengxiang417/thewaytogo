package fibonacci

var his [20]int64 = [...]int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

//fab函数的使用
func Fab_nacci(i int) int64 {
	if i <= 1 {

		return 1
	}
	if his[i] != 0 {
		return his[i]
	}

	return Fab_nacci(i-1) + Fab_nacci(i-2)

}
