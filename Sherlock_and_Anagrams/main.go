package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	rs := []rune(s)
	ans := make(map[string]int)
	sl := len(rs)
	ret := 0
	for l := 1; l <= sl; l++ {
		for i := 0; i+l <= sl; i++ {
			sub := make([]rune, l)
			copy(sub, rs[i:i+l])
			sort.Slice(sub, func(m, n int) bool {
				return sub[m] < sub[n]
			})
			ans[string(sub)]++
		}
	}
	for _, v := range ans {
		ret += v * (v - 1) / 2
	}
	return int32(ret)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
