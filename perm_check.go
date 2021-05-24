package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
    m := make(map[int]bool, len(A))

    for _, a := range A {
        if m[a] {
            return 0
        }
        m[a] = true
    }

    for i := 1; i < len(A) + 1; i++ {
        if !m[i] {
            return 0
        }
    }

    return 1
}
