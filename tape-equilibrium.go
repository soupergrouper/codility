package solution

// you can also use imports, for example:
// import "fmt"
// import "os"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    sumR := 0
    for i := 1; i < len(A); i++ {
        sumR+=A[i]
    }
    sumL := A[0]
    min := int(math.Abs(float64(sumL) - float64(sumR)))
    for i := 1; i < len(A); i++ {
        diff := int(math.Abs(float64(sumL) - float64(sumR)))
        if diff < min {
            min = diff
        }
        sumL+=A[i]
        sumR-=A[i]
    }

    return min
}
