package main

import "fmt"

func main() {
	fmt.Println(decode1([]int{6, 4, 5, 6}))
}

func decode1(encoded []int) []int {
	all := 0

	for i := 1; i <= len(encoded)+1; i++ {
		all = all ^ i
	}

	t1 := 0

	for i, v := range encoded {
		if i%2 == 1 {
			t1 = t1 ^ v
			//fmt.Println(all)
		}
	}

	t := all ^ t1
	ans := make([]int, 0)
	ans = append(ans, t)

	temp := t
	for i := 0; i < len(encoded); i++ {
		temp = temp ^ encoded[i]
		ans = append(ans, temp)
	}

	return ans
}
