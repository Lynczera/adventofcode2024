package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		line := scanner.Text()
		lineToList := strings.Split(line, " ")
		currReport := []int{}
		for i := 0; i < len(lineToList); i++ {

			element, parse_err := strconv.Atoi(lineToList[i])
			if parse_err != nil {
				panic(parse_err)
			}
			currReport = append(currReport, element)
		}
		reports = append(reports, currReport)

	}
	fmt.Println(part1(reports))
	fmt.Println(part2(reports))
}

func part1(reports [][]int) int {
	safe_reports := 0

	for i := 0; i < len(reports); i++ {
		curr_report := reports[i]
		check := curr_report[1] - curr_report[0]
		if (check > 0 && check < 4) || (check*-1 > 0 && check*-1 < 4) {
			safe := true
			for j := 1; j < len(curr_report)-1; j++ {
				curr_check := curr_report[j+1] - curr_report[j]
				if (curr_check > 0 && curr_check < 4) || (curr_check*-1 > 0 && curr_check*-1 < 4) {
					if check < 0 {
						if curr_check > 0 {
							safe = false
							break
						}
						check = curr_check
					} else {
						if curr_check < 0 {
							safe = false
							break
						}
						check = curr_check
					}
				} else {
					safe = false
					break
				}

			}
			if safe {
				safe_reports++
			}
		}
	}
	return safe_reports

}

func remove(list []int, index int) []int {
	return append(list[:index], list[index+1:]...)
}

func safe_check(report []int, removed int) bool {

	if removed > 1 {
		return false
	} else {
		check := report[1] - report[0]
		if (check > 0 && check < 4) || (check*-1 > 0 && check*-1 < 4) {
			for j := 1; j < len(report)-1; j++ {
				curr_check := report[j+1] - report[j]
				if (curr_check > 0 && curr_check < 4) || (curr_check*-1 > 0 && curr_check*-1 < 4) {
					if check < 0 {
						if curr_check > 0 {

							unsafe_list1 := make([]int, len(report))
							copy(unsafe_list1, report)

							unsafe_list2 := make([]int, len(report))
							copy(unsafe_list2, report)

							unsafe_list3 := make([]int, len(report))
							copy(unsafe_list3, report)
							return safe_check(remove(unsafe_list1, j), removed+1) || safe_check(remove(unsafe_list2, j+1), removed+1) || safe_check(remove(unsafe_list3, j-1), removed+1)
						}
						check = curr_check
						continue
					} else {
						if curr_check < 0 {
							unsafe_list1 := make([]int, len(report))
							copy(unsafe_list1, report)

							unsafe_list2 := make([]int, len(report))
							copy(unsafe_list2, report)

							unsafe_list3 := make([]int, len(report))
							copy(unsafe_list3, report)
							return safe_check(remove(unsafe_list1, j), removed+1) || safe_check(remove(unsafe_list2, j+1), removed+1) || safe_check(remove(unsafe_list3, j-1), removed+1)
						}
						check = curr_check
						continue
					}
				} else {
					unsafe_list1 := make([]int, len(report))
					copy(unsafe_list1, report)

					unsafe_list2 := make([]int, len(report))
					copy(unsafe_list2, report)

					unsafe_list3 := make([]int, len(report))
					copy(unsafe_list3, report)
					return safe_check(remove(unsafe_list1, j), removed+1) || safe_check(remove(unsafe_list2, j+1), removed+1) || safe_check(remove(unsafe_list3, j-1), removed+1)
				}

			}
			//safe
			return true
		} else {
			//unsafe
			unsafe_list1 := make([]int, len(report))
			copy(unsafe_list1, report)

			unsafe_list2 := make([]int, len(report))
			copy(unsafe_list2, report)

			return safe_check(remove(unsafe_list1, 0), removed+1) || safe_check(remove(unsafe_list2, 1), removed+1)
		}
	}

}

// 344 381 637(not tested) 317 349
func part2(reports [][]int) int {

	safes := 0

	for i := 0; i < len(reports); i++ {
		curr_report := reports[i]
		is_safe := safe_check(curr_report, 0)
		if is_safe {
			safes++
		}

	}
	return safes
}
