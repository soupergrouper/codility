package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	var min, max float64
	for _, a := range A {
		if a < 1 {
			continue
		}
		min = math.Min(min, float64(a))
		max = math.Max(max, float64(a))
	}
	integers := make([]bool, int(max-min))

	for _, a := range A {
		if a < 1 {
			continue
		}
		integers[a-1] = true
	}

	for i := range integers {
		if !integers[i] {
			return i+1
		}
	}

	return len(integers)+1
}
