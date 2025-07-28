package main

import "fmt"

func Find(array []int, num int) *[]int {

	maps := make(map[int]int)

	for i := range array {
		if index, seen := maps[array[i]]; seen {
			temp := []int{index, i}
			return &temp
		}
		maps[array[i]-num] = i
	}
	return nil
}

func main() {

	answer := Find([]int{5, 2, 3, 5}, -2)
	if answer == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println(*answer)
}
