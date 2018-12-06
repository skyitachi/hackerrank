package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
	"strings"
)

func parseLine(numberStr string) (int64) {
	number, _ := strconv.ParseInt(strings.TrimSpace(numberStr), 10, 0)
	return number
}

func solve(a string, b string) string {
    m := make(map [rune]bool)
    for _, v := range a {
        m[v] = true
    }
    for _, v := range b {
        if m[v] {
            return "YES"
        }
    }
    return "NO"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
    a := parseLine(line)
    var i int64
    for i = 0; i < a; i++ {
        l1, _ := reader.ReadString('\n')
        s1 := strings.TrimSpace(reader.ReadString('\n'))
        s2 := strings.TrimSpace(reader.ReadString('\n'))
        fmt.Println(solve(s1, s2))
    }
}
