package main

import (
  "fmt"
)

type Node struct {
  L, R int
  Value int
  Delta int
  Left *Node
  Right *Node
}

func increase(root *Node, l, r, v int) {
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
    if root.Left == nil {
      root.Left = &Node{root.L, m, 0, 0, nil, nil}
    }
    increase(root.Left, l, r, v)
  } else if l > m {
    if root.Right == nil {
      root.Right = &Node{m + 1, root.R, 0, 0, nil, nil}
    }
    increase(root.Right, l, r, v)
  } else {
    if root.Left == nil {
      root.Left = &Node{root.L, m, 0, 0, nil, nil}
    }
    increase(root.Left, l, m, v)
    if root.Right == nil {
      root.Right = &Node{m + 1, root.R, 0, 0, nil, nil}
    }
    increase(root.Right, m + 1, r, v)
  }
  if root.Left != nil {
    root.Value = root.Delta + root.Left.Value
  }
  if root.Right != nil && root.Delta + root.Right.Value > root.Value {
    root.Value = root.Delta + root.Right.Value
  }
}

func main() {
  var N, M int
  fmt.Scanf("%d%d\n", &N, &M)
  root := Node{1, N, 0,0,nil, nil}
  var a, b, v int
  for i := 0;  i < M; i++ {
    fmt.Scanf("%d%d%d\n", &a, &b, &v)
    increase(&root, a, b, v)
  }
  fmt.Println(root.Value)
}
