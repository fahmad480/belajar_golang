package main

import "fmt"

func sayWord(word string) {
	fmt.Println(word)
}

func say2Word(word1, word2 string) {
	fmt.Println(word1, word2)
}

func main() {
	sayWord("Hello")
	hai := "Hai"
	say2Word(hai, "World")
}