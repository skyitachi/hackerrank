package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sum(n int64) int64 {
	if n < 10 {
		return n
	}
	var ret int64
	for n > 0 {
		ret += n % 10
		n /= 10
	}
	return sum(ret)
}

// Complete the superDigit function below.
func superDigit(n string, k int32) int32 {
	var s int64
	for _, d := range n {
		s += int64(d - 0x30)
	}
	return int32(sum(s * int64(k)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	n := nk[0]

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

	fmt.Fprintf(writer, "%d\n", result)

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
