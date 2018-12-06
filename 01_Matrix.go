package main

import (
  "fmt"
)

func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func updateMatrixInner(matrix, ret, visit [][]int, i, j, total int) {
  if ret[i][j] != -1 || visit[i][j] != 0 {
    return
  }
  if matrix[i][j] == 0 {
    ret[i][j] = 0
    return
  }
  // 表示正在搜寻
  visit[i][j] = 1
  t1 := 1 << 32 - 1
  t2 := 1 << 32 - 1
  t3 := 1 << 32 - 1
  t4 := 1 << 32 - 1
  if i - 1 >= 0 {
    if matrix[i - 1][j] == 0 {
      ret[i][j] = 1
      return
    }
    updateMatrixInner(matrix, ret, visit, i - 1, j, total)
    t1 = ret[i - 1][j]
  }
  if j - 1 >= 0 {
    if matrix[i][j - 1] == 0 {
      ret[i][j] = 1
      return
    }
    updateMatrixInner(matrix, ret, visit, i, j - 1, total)
    t2 = ret[i][j - 1]
  }
  if i + 1 < total {
    if matrix[i + 1][j] == 0 {
      ret[i][j] = 1
      return
    }
    updateMatrixInner(matrix, ret, visit, i + 1, j, total)
    t3 = ret[i + 1][j]
  }
  if j + 1 < total {
    if matrix[i][j + 1] == 0 {
      ret[i][j] = 1
      return
    }
    updateMatrixInner(matrix, ret, visit, i, j + 1, total)
    t4 = ret[i][j + 1]
  }
  ret[i][j] = min(min(t1, t2), min(t3, t4)) + 1
}

func updateMatrix(matrix [][]int) [][]int {
  total := len(matrix)
  ret := make([][]int, len(matrix))
  visit := make([][]int, len(matrix))
  for i := range matrix {
    ret[i] = make([]int, total)
    visit[i] = make([]int, total)
    l := ret[i][:]
    for j := range ret[i] {
      l[j] = -1
    }
  }
  for i := range matrix {
    for j := range matrix[i] {
      updateMatrixInner(matrix, ret, visit, i, j, total)
    }
  }
  return ret
}


func main() {

  a := [][]int{
    []int{0, 1, 1},
    []int{1, 1, 1},
    []int{1, 1, 1},
  }
//  a := [][]int{
//    []int{0, 0},
//    []int{1, 1},
//  }
  fmt.Println(updateMatrix(a))
}
