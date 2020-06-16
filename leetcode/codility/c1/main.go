package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solution(13124545))
}

func solution(N int) int {
	if N < 5 {
		return 0
	}
	bin := strconv.FormatInt(int64(N), 2)
	fmt.Println(bin)
	spBin := strings.Split(bin, "")
	maxBin := 0
	lastCount := 0
	lastBit := "1"
	for _, bit := range spBin {
		if bit == "0" {
			lastCount += 1
			lastBit = bit
		}
		if bit == "1" && lastBit == "0" {
			if lastCount > maxBin {
				maxBin = lastCount
			}
			lastCount = 0
		}
	}
	return maxBin
}