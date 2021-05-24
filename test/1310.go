package main

import "fmt"

func main() {
	fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
}

func xorQueries(arr []int, queries [][]int) []int {
	res := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		temp := 0

		start, end := queries[i][0], queries[i][1]

		for j := start; j <= end; j++ {
			temp = temp ^ arr[j]
			fmt.Println(temp)
		}
		fmt.Println("temp is ", temp)

		res = append(res, temp)
	}

	return res
}
