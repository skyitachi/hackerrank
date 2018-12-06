package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d\n",&n)
    for i := 0; i < n; i++ {
        var s string
        dict := make(map[rune]bool)
        fmt.Scanf("%s\n", &s)
        ans := 0
        for _, c := range(s) {
            if !dict[c] {
                dict[c] = true
                ans++
            }
        }
        fmt.Println(ans)
    }
}
