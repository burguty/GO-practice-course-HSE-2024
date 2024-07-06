package main

import (
	"fmt"
)

// 21
func delete_dublicates[T comparable](arr []T) (result []T) {
	used := make(map[T]bool)
	for _, a := range arr {
		if !used[a] {
			result = append(result, a)
		}
		used[a] = true
	}
	return result
}

// 22
func bubble_sort(arr *[]int) {
	for k := 0; k < len(*arr)-1; k++ {
		is_change := false
		for i := 0; i < len(*arr)-k-1; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				(*arr)[i], (*arr)[i+1] = (*arr)[i+1], (*arr)[i]
				is_change = true
			}
		}
		if !is_change {
			break
		}
	}
}

// 23
func fibonacci(n int) {
	if n < 0 {
		panic("n must be >= 0")
	}
	f0 := 1
	f1 := 1
	f2 := 2
	for k := 0; k < n; k++ {
		fmt.Print(f0, " ")
		f2 = f0 + f1
		f0 = f1
		f1 = f2
	}
}

// 24
func count[T comparable](arr []T, query T) int {
	cnt := 0
	for _, a := range arr {
		if a == query {
			cnt++
		}
	}
	return cnt
}

// 25
func intersection(arr1 []int, arr2 []int) (result []int) {
	cnt := make(map[int]int)
	if len(arr1) > len(arr2) {
		arr1, arr2 = arr2, arr1
	}
	for _, a := range arr1 {
		cnt[a]++
	}
	for _, a := range arr2 {
		if cnt[a] > 0 {
			result = append(result, a)
			cnt[a]--
		}
	}
	return result
}

// 26
func anagram(str1 string, str2 string) bool {
	cnt1 := make(map[rune]int)
	cnt2 := make(map[rune]int)
	for _, ch := range str1 {
		cnt1[ch]++
	}
	for _, ch := range str2 {
		cnt2[ch]++
	}
	for ch, k := range cnt1 {
		if cnt2[ch] != k {
			return false
		}
	}
	for ch, k := range cnt2 {
		if cnt1[ch] != k {
			return false
		}
	}
	return true
}

// 27
func merge(arr1 []int, arr2 []int) (m []int) {
	p1 := 0
	p2 := 0
	for p1 < len(arr1) || p2 < len(arr2) {
		if p1 == len(arr1) {
			m = append(m, arr2[p2])
			p2++
		} else if p2 == len(arr2) {
			m = append(m, arr1[p1])
			p1++
		} else if arr1[p1] < arr2[p2] {
			m = append(m, arr1[p1])
			p1++
		} else {
			m = append(m, arr2[p2])
			p2++
		}
	}
	return m
}

// 28
func hash_table() {

}

// 29
func binary_search(arr []int, query int) int {
	left := -1
	right := len(arr)
	for right-left > 1 {
		mid := (left + right) / 2
		if arr[mid] >= query {
			right = mid
		} else {
			left = mid
		}
	}
	if right == len(arr) || arr[right] != query {
		return -1
	}
	return right
}

// 30

func main() {
	arr := []int{6, 2, 10, 2, 1, 6, 1, 2}
	// arr2 := []int{5, 1, 2, 2, 8, 10, 6}
	// arr_str := []string{"aba", "a", "b", "aba", "rarara"}
	// fmt.Println(delete_dublicates(arr)
	bubble_sort(&arr)
	//fmt.Println(arr)
	// fibonacci(10)
	// fmt.Println(count(arr_str, "aba"))
	// fmt.Println(intersection(arr, arr2))
	//bubble_sort(&arr)
	//bubble_sort(&arr2)
	//fmt.Println(merge(arr, arr2))
	//fmt.Println(anagram("rardara", "ararar"))
	//fmt.Println(arr)
	//fmt.Println(binary_search(arr, 7))
}
