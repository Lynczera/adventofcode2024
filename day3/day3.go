package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	var memory []string

	for scanner.Scan() {
		line := scanner.Text()
		memory = append(memory, line)

	}
	fmt.Println(part1(memory))
	fmt.Println(part2(memory))
}
func part1(mem []string) int {
	res := 0
	for _, element := range mem {
		res += compute(element)
	}
	return res
}

func compute(mem string) int {
	res := 0
	r, _ := regexp.Compile("mul\\(\\d+\\,\\d+\\)")

	r1, _ := regexp.Compile("\\d+\\,\\d+")

	operations := r.FindAllString(mem, -1)
	for _, memop := range operations {
		curr_op := r1.FindString(memop)
		nums := strings.Split(curr_op, ",")

		parsed_num1, _ := strconv.Atoi(nums[0])
		parsed_num2, _ := strconv.Atoi(nums[1])

		res += parsed_num1 * parsed_num2
	}

	return res
}

// 28208682 85698778
func part2(mem []string) int {
	res := 0
	for _, element := range mem {
		// println(sanitize(element))
		res += compute(sanitize(element))
	}
	// println(mem[0])
	// println()
	// println(sanitize(mem[0]))
	return res
}

func sanitize(unSaniMem string) string {
	// separate do and donts

	// r, _ := regexp.Compile("x?don't\\(\\).*x?do\\(\\)")

	r, _ := regexp.Compile("don't\\(\\)")
	dont_conditional := r.FindStringIndex(unSaniMem)
	if dont_conditional == nil {
		return unSaniMem
	}
	r2, _ := regexp.Compile("do\\(\\)")

	do_conditional := r2.FindStringIndex(unSaniMem[dont_conditional[1]:])

	if do_conditional == nil {
		return unSaniMem[:dont_conditional[0]]
	}
	// xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))
	return unSaniMem[:dont_conditional[0]] + sanitize(unSaniMem[dont_conditional[1]:][do_conditional[0]:])

}
