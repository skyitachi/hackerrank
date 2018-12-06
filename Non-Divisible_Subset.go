package main

import (
    "fmt"
)

func max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
func main() {
    var n, k, d int
    fmt.Scanf("%d%d\n", &n, &k)
    remainder := make(map[int]int)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &d)
        remainder[d % k]++
    }
    if k == 1 {
        fmt.Printf("1\n")
        return
    }
    var ans = 0
    for i := 0; i < k && i <= k - i; i++ {
        if i == 0 {
            ans = ans + min(remainder[0], 1)
        } else if i == k - i {
            ans = ans + min(remainder[i], 1)
        } else {
            ans = ans + max(remainder[i], remainder[k - i])
        }
    }
    fmt.Printf("%d\n", ans)
}
