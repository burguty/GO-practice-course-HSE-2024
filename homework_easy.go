package main

import (
	"fmt"
	"strings"
	"unicode"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 1
func hello_world() {
	fmt.Println("Привет, мир!")
}

// 2
func sum_two_numbers(a, b int) int {
	return a + b
}

// 3
func is_even(a int) bool {
	return a%2 == 0
}

// 4
func max_three_numbers(a, b, c int) int {
	return max(max(a, b), c)
}

// 5
func factorial(x int) int {
	var result = 1
	for i := 1; i <= x; i++ {
		result *= i
	}
	return result
}

// 6
func is_vowel(a rune) bool {
	if !unicode.IsLetter(a) {
		fmt.Println("Not a letter")
		return false
	}
	return strings.Contains("aieuoAIEUOаоуэыяёюеиАОУЭЫЯЁЮЕИ", string(a))
}

// 7
func all_primes(n int) {
	var separator []bool = make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if !separator[i] {
			fmt.Println(i)
			for j := i * 2; j <= n; j += i {
				separator[j] = true
			}
		}
	}
}

// 8
func reverse_string(str string) string {
	var result = ""
	for _, c := range str {
		result = string(c) + result
	}
	return result
}

// 9
func sum_array(array []int) int {
	var result = 0
	for _, a := range array {
		result += a
	}
	return result
}

// 10
type Rectangle struct {
	width  int
	height int
}

func get_square(rectangle Rectangle) int {
	return rectangle.width * rectangle.height
}

func main() {
	// fmt.Println(sum_array([]int{1, 3, 5}))
	rectangle := Rectangle{3, 5}
	fmt.Println(get_square(rectangle))
}
