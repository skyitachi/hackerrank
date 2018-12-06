package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
var smallest [1e5 + 1]int64
var sum int

func calc(l int) int64 {
	if smallest[l] != 0 {
		return smallest[l];
	}
	if l <= 1 {
		return 0
	} else if l == 2 {
		smallest[l] = 1
		return 1
	}
	left := (l - 1) / 2
	smallest[l] = int64(l - 1) + calc(left) + calc(l - 1 - left)
	return smallest[l]
}

func cons(ans []int64, offset, start, end int32, c int64) {
	l := end - start + 1
	if l < 1 {
		return
	}
	if l == 1 {
		ans[offset] = int64(start)
		return
	}
	if l == 2 {
		ans[offset] = int64(start)
		ans[offset + 1] = int64(end)
		return
	}
	cLeft := c - int64(l - 1)
	for i := l - 1; i >= 1; i-- {
		i64 := int64(i)
		biggest_i := i64 * (i64 - 1) / 2
		j := l - 1 - i;
		j64 := int64(j)
		biggest_j := j64 * (j64 - 1) / 2
		if cLeft >= smallest[i] + smallest[j] && cLeft <= biggest_i + biggest_j {
			x := cLeft
			for ; x >= smallest[i]; x-- {
				if cLeft - x >= smallest[j] && cLeft - x <= biggest_j {
					break
				}
			}
			if x < smallest[i] {
				continue;
			}
			//fmt.Printf("cLeft %d, offset is %d, x is %d, i is %d\n", cLeft, offset, x, i)
			ans[offset] = int64(start + i)
			cons(ans, offset + 1, start, start + i - 1, x)
			if cLeft - x >= 0 {
				cons(ans, offset + i + 1, start + i + 1, end, cLeft - x)
			}
			break
		}
	}
}

func solve(l int32, c int32) {
	c64 := int64(c)
	l64 := int64(l)
	if c64 < smallest[l] || c64 > l64 * (l64 - 1) / 2 {
		fmt.Println("-1")
		return
	}
	var ans = make([]int64, l64)
	cons(ans[:],0, 1, l, c64)
	for i := 0; i < int(l); i++ {
		if i > 0 {
			fmt.Printf(" %d", ans[i])
		} else {
			fmt.Printf("%d", ans[i])
		}
	}
	fmt.Printf("\n")

	//_, cv := lena_sort(ans)
	//if cv != c64 {
	//	sum += 1
	//	fmt.Println(l, c, cv)
	//}
	//fmt.Println("smallest: ", smallest[l])
}

func lena_sort(input []int64) ([]int64, int64) {
	if len(input) <= 1 {
		return input, 0
	}
	pivot := input[0]
	var less []int64
	var more []int64
	c := int64(0)
	for i := 1; i < len(input); i++ {
		c += 1
		if input[i] < pivot {
			less = append(less, input[i])
		} else {
			more = append(more, input[i])
		}
	}
	sorted_less, cl := lena_sort(less)
	sorted_more, cr := lena_sort(more)
	return append(append(sorted_less, pivot), sorted_more...), cl + cr + c
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)
	for i := 0; i < 1e5 + 1; i++ {
		calc(i)
		//if i < 10 {
		//	fmt.Printf("smallest %d: %d\n", i, smallest[i])
		//}
	}
	for qItr := 0; qItr < int(q); qItr++ {
		lc := strings.Split(readLine(reader), " ")

		lTemp, err := strconv.ParseInt(lc[0], 10, 64)
		checkError(err)
		l := int32(lTemp)

		cTemp, err := strconv.ParseInt(lc[1], 10, 64)
		checkError(err)
		c := int32(cTemp)


		// Write Your Code Here
		//fmt.Println("smallest: ", smallest[l])
		solve(l, c)
	}
	//fmt.Println("sum: ", sum)
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
