package mathutils

func Factorial(a int) int {
	if a < 0 {
		return -1
	}
	if a == 0 {
		return 1
	}
	r := 1
	for i := 1; i <= a; i++ {
		r *= i
	}
	return r
}
