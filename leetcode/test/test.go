package main

import "fmt"

func main() {
	var str []string
	var dat = []string{"hi"}
	dat = append(dat, str...)
	fmt.Println(dat)
	str = []string{"add"}
	dat = append(dat, str...)
	fmt.Println(dat)
}
