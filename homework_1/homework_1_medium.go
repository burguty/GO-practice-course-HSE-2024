package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

// 11
func from_c_to_f(celsius float64) float64 {
	if celsius < -273.15 {
		panic("The absolute zero has been reached!")
	}
	return celsius*9/5 + 32
}

// 12
func countdown() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n < 1 {
		panic("Expected natural number!")
	}
	for i := n; i > 0; i-- {
		fmt.Println(i)
	}
}

// 13
func len_string(s string) int {
	l := 0
	for range s {
		l++
	}
	return l
}

// 14
func contains[T comparable](arr []T, element T) bool {
	for _, a := range arr {
		if a == element {
			return true
		}
	}
	return false
}

func enter_for_contains() {
	var (
		n     int
		arr   []int
		query int
	)

	_, err := fmt.Scan(&n)
	if err != nil || n < 0 {
		panic("Expected integer number >= 0!")
	}
	for i := 0; i < n; i++ {
		var a int
		_, err := fmt.Scan(&a)
		if err != nil {
			panic("Expected number!")
		}
		arr = append(arr, a)
	}

	_, err = fmt.Scan(&query)
	if err != nil {
		panic("Expected number!")
	}

	if contains(arr, query) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// 15
func average(arr []int) float64 {
	if len(arr) == 0 {
		panic("Empty array!")
	}
	sum := 0.0
	for _, a := range arr {
		sum += float64(a)
	}
	return sum / float64(len(arr))
}

func enter_for_average() {
	var (
		arr []int
		n   int
	)
	_, err := fmt.Scan(&n)
	if err != nil || n < 0 {
		panic("Expected integer number >= 0!")
	}
	for i := 0; i < n; i++ {
		var a int
		_, err := fmt.Scan(&a)
		if err != nil {
			panic("Expected number!")
		}
		arr = append(arr, a)
	}
	fmt.Println(average(arr))
}

// 16
func multiply_table() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic("Expected integer number!")
	}
	for i := 1; i < 11; i++ {
		fmt.Println(n, "x", i, "=", n*i)
	}
}

// 17
func palindrome() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = s[0 : len(s)-1]
	for i := 0; i < len(s)-i-1; i++ {
		if s[i] != s[len(s)-i-1] {
			fmt.Println("Not Palindrome!")
			return
		}
	}
	fmt.Println("Palindrome!")
}

// 18
func min_max(arr []int) (int, int) {
	max_value := math.MinInt
	min_value := math.MaxInt
	for _, a := range arr {
		if a < min_value {
			min_value = a
		}
		if a > max_value {
			max_value = a
		}
	}
	return min_value, max_value
}

// 19
func erase_element(arr *[]int, index int) {
	if index < 0 || index >= len(*arr) {
		panic("Index out of range!")
	}
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
}

// 20
func linear_search(arr []int, query int) int {
	for ind, a := range arr {
		if a == query {
			return ind
		}
	}
	return -1
}

func main() {
	// fmt.Println(from_cels_to_far(1))
	// countdown()
	//var s string
	//_, err := fmt.Scan(&s)
	//if err != nil {
	//	fmt.Println("mdmd")
	//	return
	//}
	//fmt.Println(len_string(s))
	// enter_for_contains()
	// enter_for_average()
	// multiply_table()
	// palindrome()
	// fmt.Println(min_max([]int{1, 100, 3}))
	//arr := []int{1, 2, 3, 4, 5}
	//erase_element(&arr, 2)
	//fmt.Println(arr)
	// fmt.Println(linear_search([]int{1, 2, 3, 4, 5}, 1))
}
