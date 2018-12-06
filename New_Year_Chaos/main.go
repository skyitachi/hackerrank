package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	var ret, i int32
	var length = int32(len(q))
	var t1 = make([]int32, length+1)
	for i = 0; i < length+1; i++ {
		t1[i] = i
	}
	var t2 = make([]int32, length+1)
	copy(t2[:], t1)
	for i = 1; i <= length; i++ {
		v := q[i-1]
		if v > i+2 {
			ret = -1
			break
		} else if t2[v] > i {
			for j := t2[v] - 1; j >= i; j-- {
				t1[j+1] = t1[j]
				t2[t1[j]]++
				ret++
			}
			t2[v] = i
		}
	}
	if ret == -1 {
		fmt.Println("Too chaotic")
	} else {
		fmt.Println(ret)
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
