package main

import "fmt"

func main() {
	var array_numbers [3]int
	array_numbers[0] = 1
	array_numbers[0] = 2
	array_numbers[0] = 3

	fmt.Println(array_numbers)

	array_words := [2]string {"hello", "world"}

	fmt.Println(array_words)

	array_words_2 := [...]string {"hello", "capybara", ":P"}

	fmt.Println(array_words_2)

	slice_words := []string {"capybara", "is", "dev"}

	fmt.Println(slice_words)

	slice_words = append(slice_words, "coffee")

	fmt.Println(slice_words)

	slice_words3 := make([]int8, 5, 10)

	fmt.Println(slice_words3)
	fmt.Println(len(slice_words3))
	fmt.Println(cap(slice_words3))

	remove_words := []string {"a", "b", "c"}

	fmt.Println(remove_words)

	remove_words = remove_words[1:]

	fmt.Println(remove_words)
}