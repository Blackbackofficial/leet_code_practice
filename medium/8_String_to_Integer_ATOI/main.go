package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	if scaner.Scan() {
		str := scaner.Text()
		fmt.Println(myAtoi(str))
	}
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" || s == "0" {
		return 0
	}

	verb := 1
	if s[0] == 45 {
		verb = -1
		s = s[1:]
	} else if s[0] == 43 {
		s = s[1:]
	} else if s[0] < 48 || s[0] > 57 {
		return 0
	}

	var res int
	for _, v := range s {
		if v >= 48 && v <= 57 {
			res = 10*res + int(v) - 48
			if res > 1<<31-1 || res < -1<<31 {
				return (1<<31-1)*verb + (verb-1)>>2
			}
		} else {
			break
		}
	}
	return res * verb
}
