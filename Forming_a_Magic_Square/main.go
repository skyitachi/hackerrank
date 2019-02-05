package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"math"
)

func abs(a int32) int32 {
	if a >= 0 {
		return a
	}
	return -a
}

type Permutation struct {
	Data []int32
	counter int32
	Size int32
}

func (p *Permutation) Next() []int32 {
	if p.counter == 0 {
		var i int32
		for i = 0; i < p.Size; i++ {
			if i >= 4 {
				p.Data = append(p.Data, i + 2)
			} else {
				p.Data = append(p.Data, i + 1)
			}
		}
		p.counter += 1
		return p.Data
	}
	var maxK = -1
	var intSize = int(p.Size)
	for k := 0; k + 1 < intSize; k++ {
		if p.Data[k] < p.Data[k + 1] {
			maxK = k
		}
	}
	if maxK == -1 {
		return p.Data
	}
	var minL = maxK + 1
	var minLValue = p.Data[minL]
	for i := minL + 1; i < intSize; i++ {
		if p.Data[i] > p.Data[maxK] && p.Data[i] < minLValue {
			minL = i
			minLValue = p.Data[i]
		}
	}
	// swap maxK and minL
	temp := p.Data[maxK]
	p.Data[maxK] = p.Data[minL]
	p.Data[minL] = temp

	// reverse maxK + 1 to n
	var i1 = maxK + 1
	var j1 = intSize - 1
	for ; i1 < j1; {
		temp = p.Data[i1]
		p.Data[i1] = p.Data[j1]
		p.Data[j1]= temp
		i1++
		j1--
	}
	return p.Data
}

func (p *Permutation) Show() {
	for i := 0; i < int(p.Size); i++ {
		fmt.Print(p.Data[i])
	}
	fmt.Print("\n")
}

func (p *Permutation) IsMagic() bool {
	if p.Data[0] + p.Data[1] + p.Data[2] != 15 {
		return false
	}
	if p.Data[3] + 5 + p.Data[4] != 15 {
		return false
	}
	if p.Data[5] + p.Data[6] + p.Data[7] != 15 {
		return false
	}
	if p.Data[0] + p.Data[3] + p.Data[5] != 15 {
		return false
	}
	if p.Data[1] + 5 + p.Data[6] != 15 {
		return false
	}
	if p.Data[2] + p.Data[4] + p.Data[7] != 15 {
		return false
	}
	if p.Data[0] + 5 + p.Data[7] != 15 {
		return false
	}
	if p.Data[5] + 5 + p.Data[2] != 15 {
		return false
	}
	return true
}

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) int32 {
	var ans int32 = math.MaxInt32
	count := 40320
	permutation := Permutation{Size: 8}
	for c := 0; c < count; c++ {
		permutation.Next()
		if !permutation.IsMagic() {
			continue
		}
		var temp int32
		if s[1][1] != 5 {
			temp = abs(s[1][1] - 5)
		}
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				k := i * 3 + j
				if k == 4 {
					continue
				}
				if k > 4 {
					k -= 1
				}
				temp += abs(permutation.Data[k] - s[i][j])
			}
		}
		if temp < ans {
			ans = temp
		}
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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
