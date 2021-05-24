package main

import "fmt"

var trees = make([]string, 0)

func main() {

	frequencySort("haha12h")
}

func frequencySort(s string) string {
	//bucket := make(map[string]int)
	for _, v := range s {
		buildTree(string(v))

		//v1 := string(v)
		//bucket[v1]++
	}

	//for k, v := range bucket {
	//
	//	for i := 0; i < v; i++ {
	//		buildTree(k)
	//	}
	//}

	fmt.Println(trees)

	return ""
}

func buildTree(v string) {
	lenTree := len(trees)
	trees = append(trees, v)
	currentNum := (lenTree - 1) / 2

	for true {
		if currentNum < 0 {
			break
		}

		currVal := trees[currentNum]

		if lenTree > (currentNum * 2) {
			leftChild := trees[currentNum*2]
			if currVal < leftChild {
				trees[currentNum], trees[currentNum*2] = trees[currentNum*2], trees[currentNum]
			}
		}

		if lenTree > (currentNum*2 + 1) {
			rightChild := trees[currentNum*2+1]
			if currVal < rightChild {
				trees[currentNum], trees[currentNum*2+1] = trees[currentNum*2+1], trees[currentNum]
			}
		}

		currentNum = currentNum - 1
	}
}
