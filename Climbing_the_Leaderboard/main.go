package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Idx   int32
	Value int32
}

// Complete the climbingLeaderboard function below.
// use O(n)
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	var list []Item
	var lastValue int32 = -1
	var idx int32 = 1
	var ans []int32
	for i, s := range scores {
		if i == 0 {
			list = append(list, Item{Idx: idx, Value: s})
			lastValue = s
			continue
		}
		if s < lastValue {
			idx++
		}
		list = append(list, Item{Idx: idx, Value: s})
		lastValue = s
	}
	cursor := len(list) - 1
	for _, s := range alice {
		for ; cursor >= 0 && list[cursor].Value < s; cursor-- {
		}
		if cursor < 0 {
			ans = append(ans, 1)
			continue
		}
		if list[cursor].Value == s {
			ans = append(ans, list[cursor].Idx)
		} else if list[cursor].Value > s {
			ans = append(ans, list[cursor].Idx+1)
		}
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32
	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
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
