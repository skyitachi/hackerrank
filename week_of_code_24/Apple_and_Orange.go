package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) (int64, int64) {
	numbers := strings.Split(strings.TrimSpace(line), " ")
	first, _ := strconv.ParseInt(numbers[0], 10, 0)
	second, _ := strconv.ParseInt(numbers[1], 10, 0)
	return first, second
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	s, t := parseLine(line)
	line, _ = reader.ReadString('\n')
	a, b := parseLine(line)
	line, _ = reader.ReadString('\n')
	_, _ = parseLine(line)
	ans1, ans2 := 0, 0
	line, _ = reader.ReadString('\n')
	apples := strings.Split(strings.TrimSpace(line), " ")
	line, _ = reader.ReadString('\n')
	oranges := strings.Split(strings.TrimSpace(line), " ")
	for _, v := range apples {
		tmp, _ := strconv.ParseInt(v, 10, 0)
		d1 := a + tmp
		if d1 >= s && d1 <= t {
			ans1 = ans1 + 1
		}
	}
	for _, v := range oranges {
		tmp, _ := strconv.ParseInt(v, 10, 0)
		d1 := b + tmp
		if d1 >= s && d1 <= t {
			ans2 = ans2 + 1
		}
	}
	fmt.Printf("%d\n%d\n", ans1, ans2)
}
