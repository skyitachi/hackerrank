package main

import "fmt"

func main() {
	var g int
	_, _ = fmt.Scanf("%d", &g)
	for i := 0; i < g; i++ {
		var n int
		var b string
		var sum [26]int
		empty := false
		same := true
		flag := true
		_, _ = fmt.Scanf("%d", &n)
		_, _ = fmt.Scanf("%s", &b)
		l := len(b)
		for j := 0; j < l-1; j++ {
			if b[j] == '_' {
				empty = true
				continue
			} else {
				sum[b[j]-'A']++
			}
			if j > 0 && (b[j] != b[j-1] && b[j] != b[j+1]) {
				same = false
			} else if j == 0 && b[j] != b[j+1] {
				same = false
			}
		}
		if b[l-1] == '_' {
			empty = true
		} else if l > 1 && b[l-2] != b[l-1] {
			same = false
			sum[b[l-1]-'A']++
		}
		if l == 1 && b[l-1] != '_' {
			flag = false
		} else if same {
			flag = true
		} else if !same && !empty {
			flag = false
		} else if !same {
			for _, s := range sum {
				if s == 1 {
					flag = false
					break
				}
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
