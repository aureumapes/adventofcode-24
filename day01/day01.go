package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputB, _ := os.ReadFile("input.txt")
	input := string(inputB)
	inputLines := strings.Split(input, "\n")
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	for _, inputLine := range inputLines {
		lists := strings.Split(inputLine, "   ")
		one, _ := strconv.Atoi(lists[0])
		two, _ := strconv.Atoi(lists[1])
		list1 = append(list1, one)
		list2 = append(list2, two)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	result := 0
	for i, _ := range list1 {
		i1 := list1[i]
		i2 := list2[i]
		if i1 > i2 {
			result += i1 - i2
		} else {
			result += i2 - i1
		}
	}
	fmt.Println("Part One: ", result)

	// part Two
	result = 0
	for _, numb1 := range list1 {
		similarity := 0
		for _, numb2 := range list2 {
			if numb1 == numb2 {
				similarity++
			}
		}
		result += similarity * numb1
	}
	fmt.Println("Part Two: ", result)
}
