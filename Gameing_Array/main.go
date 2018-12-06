package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the gamingArray function below.
func gamingArray(arr []int32) int32 {
	l := len(arr)
	if l == 1 {
		return 1
	} else if arr[0] < arr[1] {
		return -gamingArray(arr[1:])
	}
	i := 1
	for ; i < l; i++ {
		if arr[i] > arr[0] {
			break
		}
	}
	if i == l {
		return 1
	}
	return -gamingArray(arr[i:])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	gTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		arrCount, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int32

		for i := 0; i < int(arrCount); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := gamingArray(arr)
		if result == 1 {
			fmt.Fprint(writer, "BOB\n")
		} else if result == -1 {
			fmt.Fprint(writer, "ANDY\n")
		}
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
