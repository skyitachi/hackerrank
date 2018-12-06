package main

import (
	"fmt"
)

func main() {
  var N int
  var Q int
  var str string
  dict := make(map[string]int)
  fmt.Scanf("%d\n", &N)
  for i := 0; i < N; i++ {
    fmt.Scanf("%s\n", &str)
    dict[str] += 1
  }
  fmt.Scanf("%d\n", &Q)
  for i := 0; i < Q; i++ {
    fmt.Scanf("%s\n", &str)
    fmt.Printf("%d\n", dict[str])
  }
}
