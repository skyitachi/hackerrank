package main

import (
  "fmt"
  "time"
)

type Node struct {
  L, R int
  Value int
  Delta int
  Left int
  Right int
}

func build(root *Node, idx int, list []*Node ) {
  if root.L == root.R {
    return
  }
  root.Left = idx * 2
  root.Right = idx * 2 + 1
  m := (root.L + root.R) / 2
  left := Node{root.L, m, 0,0,0, 0 }
  right := Node { m + 1, root.R, 0,0, 0, 0 }
  build(&left, root.Left, list)
  build(&right, root.Right, list)
  list[root.Left] = &left
  list[root.Right] = &right
}

func increase(root *Node, list []*Node, l, r, v int) {
  if root == nil {
    return
  }
  if l == root.L && r == root.R {
    root.Delta += v
    root.Value += v
    return
  }
  m := (root.L + root.R) / 2
  if r <= m {
    increase(list[root.Left], list, l, r, v)
  } else if l > m {
    increase(list[root.Right], list, l, r, v)
  } else {
    increase(list[root.Left], list, l, m, v)
    increase(list[root.Right], list, m + 1, r, v)
  }
  if root.Left != 0 {
    root.Value = root.Delta + list[root.Left].Value
  }
  if root.Right != 0 && root.Delta + list[root.Right].Value > root.Value {
    root.Value = root.Delta + list[root.Right].Value
  }
}

func main() {
  var N, M int
  fmt.Scanf("%d%d\n", &N, &M)
  root := Node{1, N, 0,0,0, 0}
  now := time.Now().Unix()
  nodeList := make([]*Node, N * 4)
  build(&root, 1, nodeList)
  fmt.Println(time.Now().Unix() - now)
  var a, b, v int
  for i := 0;  i < M; i++ {
    fmt.Scanf("%d%d%d\n", &a, &b, &v)
    increase(&root, nodeList, a, b, v)
  }
  fmt.Println(root.Value)
}
