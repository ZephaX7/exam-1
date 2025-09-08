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
	fmt.Println(CountLetters("Curry"))   // 10
	fmt.Println(CountLetters("JSP"))     // 0
	fmt.Println(CountLetters("Coucou"))  // 6
	fmt.Println(CountLetters("bonjour")) // 4
}
