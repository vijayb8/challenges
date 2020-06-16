package main

import "fmt"

func main()  {
	fmt.Println(solution([]int{2,8,9,7,6}, 3))
}


func solution(A []int, K int) []int {
	if len(A) == K || K == 0{
		return A
	}
	fmt.Println(len(A)-1)
	fmt.Println(A[len(A)-1])
	// write your code in Go 1.4
	for i:=0; i<K; i++ {
		arr := make([]int, len(A))
		for j := 0; j<len(A) ; j++ {
			if j == 0 {
				val := A[len(A)-1]
				arr[j] = val
				continue
			}
			arr[j] = A[j-1]
		}
		A = arr
	}
	return A
}

