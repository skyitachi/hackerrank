package main

import (
    "fmt"
    "math"
)

func main() {
    var trimmed, ans string
    var transformed []string
    fmt.Scanf("%s\n", &trimmed)
    L := len(trimmed)
    rows := (int)(math.Floor(math.Sqrt(float64(L))))
    cols := (int)(math.Ceil(math.Sqrt(float64(L))))
    if rows * cols < L {
        rows += 1
    }
    transformed = make([]string, rows)
    for i := 0; i < rows; i++ {
        transformed[i] = ""
        for j := 0; j < cols; j++ {
            if i * cols + j >= L {
                break
            }
            transformed[i] += string(trimmed[i * cols + j])
        }
    }
    for i := 0; i < cols; i++ {
        c := ""
        if i > 0 {
            ans += " "
        }
        for j := 0; j < rows; j++ {
            if j * cols + i >= L {
                break
            } else if transformed[j][i] != 0 {
                c += string(transformed[j][i])
            } else {
                break
            }
        }
        ans += c
    }
    fmt.Println(ans)
}
