func Solution(X int, A []int) int {
	c := make([]bool, X)
	s := X
	for i, a := range A {
		if a > X {
			continue
		}
		if !c[a-1] {
			s--
			c[a-1]=true
		}
		if s == 0 {
			return i
		}
	}

	return -1
}
