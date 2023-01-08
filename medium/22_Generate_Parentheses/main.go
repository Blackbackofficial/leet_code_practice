package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	fmt.Println(generateParenthesis(n))
}

func generateParenthesis(n int) []string {

	return []string{}
}
