package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(solution([]int{9, 3, 9, 3, 9, 7, 9, 7, 4}))
}

func solution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	sort.Ints(A)
	for i:=0; i<len(A)-1; i+=2 {
		if A[i] == A[i+1] {
			continue
		}
		return A[i]
	}
	return A[len(A)-1]
}
