package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(list1 []int, list2 []int) {

	result := 0

	slices.Sort(list1)
	slices.Sort(list2)

	for i := 0; i < len(list1); i++ {
		curr := list1[i] - list2[i]

		if curr < 0 {
			curr = curr * -1
		}
		result += curr

	}

	fmt.Println("%v", result)
}

func part2(list1 []int, list2 []int) {

	slices.Sort(list1)
	slices.Sort(list2)
	result := 0

	freq := make(map[int]int)

	// calculate how many times a num in list1 appears in list2

	for i := 0; i < len(list1); i++ {
		_, ok := freq[list1[i]]
		if !ok {
			count := 0
			for j := 0; j < len(list1); j++ {
				if list2[j] == list1[i] {
					count++
				}

			}
			freq[list1[i]] = count
		}

	}

	for k, v := range freq {

		result += k * v
	}
	fmt.Println("%v", result)

}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		lineToList := strings.Split(line, "   ")
		element1, _ := strconv.Atoi(lineToList[0])
		element2, _ := strconv.Atoi(lineToList[1])
		list1 = append(list1, element1)
		list2 = append(list2, element2)

	}

	// part1(list1, list2)
	part2(list1, list2)

}
