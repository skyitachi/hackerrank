package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Value      int32
	Left       *TreeNode
	LeftIndex  int32
	Right      *TreeNode
	RightIndex int32
	Height     int32
}

func buildHeight(root *TreeNode, currentHeight int32) int32 {
	var leftHeight, rightHeight int32
	root.Height = currentHeight
	if root.Left != nil {
		leftHeight = buildHeight(root.Left, currentHeight+1)
	}
	if root.Right != nil {
		rightHeight = buildHeight(root.Right, currentHeight+1)
	}
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

func swap(root *TreeNode, height int32) {
	if root.Height%height == 0 {
		root.Left, root.Right = root.Right, root.Left
	}
	if root.Left != nil {
		swap(root.Left, height)
	}
	if root.Right != nil {
		swap(root.Right, height)
	}
}

func inorder(root *TreeNode, visitList []int32, visited int) int {
	var leftCount, rightCount int
	if root.Left != nil {
		leftCount = inorder(root.Left, visitList, visited)
		visited += leftCount
	}
	visitList[visited] = root.Value
	if root.Right != nil {
		rightCount = inorder(root.Right, visitList, visited+1)
	}
	return leftCount + rightCount + 1
}

/*
 * Complete the swapNodes function below.
 */
func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	/*
	 * Write your code here.
	 */
	treeNodes := make([]TreeNode, len(indexes)+1)
	for idx, index := range indexes {
		treeNodes[idx+1] = TreeNode{
			Value:      int32(idx) + 1,
			LeftIndex:  index[0],
			RightIndex: index[1],
		}
	}
	for i := 1; i < len(treeNodes); i++ {
		node := &treeNodes[i]
		if node.LeftIndex != -1 {
			node.Left = &treeNodes[node.LeftIndex]
		}
		if node.RightIndex != -1 {
			node.Right = &treeNodes[node.RightIndex]
		}
	}
	buildHeight(&treeNodes[1], 1)
	ans := make([][]int32, len(queries))
	for idx, query := range queries {
		swap(&treeNodes[1], query)
		ret := make([]int32, len(indexes))
		inorder(&treeNodes[1], ret, 0)
		ans[idx] = ret
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var indexes [][]int32
	for indexesRowItr := 0; indexesRowItr < int(n); indexesRowItr++ {
		indexesRowTemp := strings.Split(readLine(reader), " ")

		var indexesRow []int32
		for _, indexesRowItem := range indexesRowTemp {
			indexesItemTemp, err := strconv.ParseInt(indexesRowItem, 10, 64)
			checkError(err)
			indexesItem := int32(indexesItemTemp)
			indexesRow = append(indexesRow, indexesItem)
		}

		if len(indexesRow) != int(2) {
			panic("Bad input")
		}

		indexes = append(indexes, indexesRow)
	}

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int32

	for queriesItr := 0; queriesItr < int(queriesCount); queriesItr++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := swapNodes(indexes, queries)

	for resultRowItr, rowItem := range result {
		for resultColumnItr, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if resultColumnItr != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if resultRowItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
