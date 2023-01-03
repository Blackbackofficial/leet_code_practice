package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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
	firstLetter := false
	checkLetter := false
	for k, v := range word {
		if unicode.IsLetter(v) {
			if unicode.IsUpper(v) && k == 0 { // for the first letter
				firstLetter = true
				checkLetter = true
				continue
			} else if firstLetter && unicode.IsLower(v) { // CASE: Hello
				if k == 1 && checkLetter == false { // CASE: Hi
					continue
				}
				if k == len(word)-1 && len(word) != 2 && checkLetter == true { // CASE: CASe && outer include Hi
					return false
				}
				if checkLetter == true && k > 1 {
					return false
				}
				checkLetter = false
				continue
			} else if !firstLetter && unicode.IsLower(v) && checkLetter == false { // CASE: hello
				checkLetter = false
				continue
			} else if firstLetter && unicode.IsUpper(v) && checkLetter == true { // CASE: HELLO
				checkLetter = true
				continue
			}
		}
		return false
	}
	return true
}
