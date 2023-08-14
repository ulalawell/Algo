package main

import "fmt"

func main() {
	var data = []int{5, 3, 6, 1}

	fmt.Println(bubbleSort(data))
}

func bubbleSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}
