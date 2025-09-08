package main

import "fmt"

func CountLetters(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountLetters("Curry"))
	fmt.Println(CountLetters("JSP"))
	fmt.Println(CountLetters("Coucou"))
	fmt.Println(CountLetters("bonjour"))
}
