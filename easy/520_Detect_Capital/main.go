package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = text[0 : len(text)-1]
	b := DetectCapitalUse(text)
	fmt.Print(b)
}

func DetectCapitalUse(word string) bool {
	capital := false
	nonCapital := false
	for index, runeValue := range word {
		if 97 <= runeValue && runeValue <= 122 || 65 <= runeValue && runeValue <= 90 {
			if runeValue-97 >= 0 {
				nonCapital = true
			} else if index > 0 {
				capital = true
			}

			if capital && nonCapital {
				return false
			}
			continue
		}
		return false
	}
	return true
}
