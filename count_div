package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
    k := A / K
    first := K*k
    if A%K != 0 {
        first+=K
    }

    return B/K - first/K + 1
}
