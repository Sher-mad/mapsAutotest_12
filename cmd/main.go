package main

import (
	"fmt"
)

func main() {
	// arr := []types.Payment{}
	// fmt.Println(arr)
	categories := []string{"auto", "food", "beaty", "mobile", "fun"}
	top3:= categories[0:3]
	fmt.Println(categories)
	fmt.Println(len(categories))
	fmt.Println(cap(categories))
	fmt.Println(top3)
	fmt.Println(len(top3))
	fmt.Println(cap(top3))

}
