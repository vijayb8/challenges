package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution("ccaaffddecee"))
}


func Solution(S string) int {
	if len(S) == 0 {
		return 0
	}
	// create a frequency map of all alphabets
	freqMap := map[string]int{}

	for _, val := range strings.Split(S, "") {
		if _, ok := freqMap[val]; ok {
			freqMap[val] = freqMap[val]+1
			continue
		}
		freqMap[val] = 1
	}

	var deletionsRequired = 0

	iterCount := map[int]int{}

	for _, v := range freqMap {
		if _, ok:= iterCount[v]; ok {
			iterCount[v] = iterCount[v] + 1
			continue
		}
		iterCount[v] = 1
	}

	keys := make([]int, 0, len(iterCount))

	for k := range iterCount {
		keys = append(keys, k)
	}
	for i, v := range keys {
		if v > 1 {
			deletionsRequired += 1
			iterCount[v] = iterCount[v] - 1
			if i < len(iterCount) {
				iterCount[keys[i+1]] = iterCount[keys[i+1]] + 1
			}
		}
	}
	return deletionsRequired

	//for k, v := range iterCount {
	//	if v > 1 {
	//		deletionsRequired+=1
	//		iterCount[k] = iterCount[k] - 1
	//
	//	}
	//}
	//
	//for _, v := range freqMap {
	//	if v > max {
	//		max = v
	//		if min == 0 {
	//			min = v
	//		}
	//	} else {
	//		if v < min {
	//			min = v
	//			if max==0 {
	//				max = v
	//			}
	//		} else {
	//			if v >= min && v <= max {
	//				if min == 1{
	//					deletionsRequired+=1
	//				} else {
	//					min-=1
	//					deletionsRequired+=v-min
	//				}
	//			}
	//		}
	//	}
	//}
	//return deletionsRequired
}

