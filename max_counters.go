package solution

// you can also use imports, for example:
// import "fmt"
import "math"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	n := make([]int, N)
	var maxCur, maxGlob int
	lastMaxIndx := -1
	for i, a := range A {
		k := a-1
		if k > N {
			continue
		}
		if k < N {
			if n[k] < maxGlob {
				n[k] = maxGlob
			}
			n[k]++
			maxCur = int(math.Max(float64(maxCur), float64(n[k])))
			continue
		}
		maxGlob = maxCur
		lastMaxIndx = i
	}
	for i := range n {
		n[i] = maxGlob
	}
	for i := lastMaxIndx+1; i < len(A); i++ {
		k := A[i] - 1
		n[k]++
	}

	return n
}
