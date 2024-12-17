package main

import (
	"bufio"
	"os"
)

// var directions = [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}, []int{1, -1}, []int{1, 1}, []int{-1, -1}, []int{-1, 1}}

const target = "XMAS"
const revtarget = "SAMX"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)

	}
	// println(part1(words))
	println(part2(words))
}
func part1(words []string) int {
	graph := graphIt(words)
	count := 0

	for r, _ := range graph {
		rlimit := len(graph)
		for c, _ := range graph[r] {
			climit := len(graph[r])
			//up
			if r >= 3 {
				w := ""
				for i := range 4 {
					w += graph[r+(-1*i)][c]
				}
				if w == target {
					count++
				}

			}
			//down
			if r <= rlimit-4 {
				w := ""
				for i := range 4 {
					w += graph[r+i][c]
				}
				if w == target {
					count++
				}
			}
			//left
			if c >= 3 {
				w := ""
				for i := 0; i >= -3; i-- {
					w += graph[r][c+i]
				}
				if w == target {
					count++
				}
			}
			//right
			if c <= climit-4 {
				w := ""
				for i := range 4 {
					w += graph[r][c+i]
				}
				if w == target {
					count++
				}
			}
			// ul
			if r >= 3 && c >= 3 {
				w := ""
				for t := range 4 {
					w += graph[r+(-1*t)][c+(-1*t)]
				}
				if w == target {
					count++
				}
			}
			// ur
			if r >= 3 && c <= climit-4 {
				w := ""
				for t := range 4 {
					w += graph[r+(-1*t)][c+t]
				}
				if w == target {
					count++
				}
			}
			// dl
			if r <= rlimit-4 && c >= 3 {
				w := ""
				for t := range 4 {
					w += graph[r+t][c+(-1*t)]
				}
				if w == target {
					count++
				}
			}
			// dr
			if r <= rlimit-4 && c <= climit-4 {
				w := ""
				for t := range 4 {
					w += graph[r+t][c+t]
				}
				if w == target {
					count++
				}
			}

		}

	}
	return count
}

func part2(words []string) int {
	graph := graphIt(words)
	count := 0

	for r, _ := range graph {
		rlimit := len(graph)
		for c, _ := range graph[r] {
			climit := len(graph[r])
			if graph[r][c] == "A" {
				//check bounds
				if r >= 1 && r <= rlimit-2 && c >= 1 && c <= climit-2 {
					if ((graph[r-1][c-1] == "M" && graph[r+1][c+1] == "S") || (graph[r-1][c-1] == "S" && graph[r+1][c+1] == "M")) && ((graph[r-1][c+1] == "S" && graph[r+1][c-1] == "M") || (graph[r-1][c+1] == "M" && graph[r+1][c-1] == "S")) {
						count++
					}

				}
			}
		}

	}
	return count
}

func graphIt(words []string) [][]string {
	graph := [][]string{}

	for _, row := range words {
		currRow := []string{}
		for _, letter := range row {
			currRow = append(currRow, string(letter))

		}
		graph = append(graph, currRow)
	}
	return graph
}
package main

import (
	"bufio"
	"os"
)

// var directions = [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}, []int{1, -1}, []int{1, 1}, []int{-1, -1}, []int{-1, 1}}

const target = "XMAS"
const revtarget = "SAMX"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)

	}
	// println(part1(words))
	println(part2(words))
}
func part1(words []string) int {
	graph := graphIt(words)
	count := 0

	for r, _ := range graph {
		rlimit := len(graph)
		for c, _ := range graph[r] {
			climit := len(graph[r])
			//up
			if r >= 3 {
				w := ""
				for i := range 4 {
					w += graph[r+(-1*i)][c]
				}
				if w == target {
					count++
				}

			}
			//down
			if r <= rlimit-4 {
				w := ""
				for i := range 4 {
					w += graph[r+i][c]
				}
				if w == target {
					count++
				}
			}
			//left
			if c >= 3 {
				w := ""
				for i := 0; i >= -3; i-- {
					w += graph[r][c+i]
				}
				if w == target {
					count++
				}
			}
			//right
			if c <= climit-4 {
				w := ""
				for i := range 4 {
					w += graph[r][c+i]
				}
				if w == target {
					count++
				}
			}
			// ul
			if r >= 3 && c >= 3 {
				w := ""
				for t := range 4 {
					w += graph[r+(-1*t)][c+(-1*t)]
				}
				if w == target {
					count++
				}
			}
			// ur
			if r >= 3 && c <= climit-4 {
				w := ""
				for t := range 4 {
					w += graph[r+(-1*t)][c+t]
				}
				if w == target {
					count++
				}
			}
			// dl
			if r <= rlimit-4 && c >= 3 {
				w := ""
				for t := range 4 {
					w += graph[r+t][c+(-1*t)]
				}
				if w == target {
					count++
				}
			}
			// dr
			if r <= rlimit-4 && c <= climit-4 {
				w := ""
				for t := range 4 {
					w += graph[r+t][c+t]
				}
				if w == target {
					count++
				}
			}

		}

	}
	return count
}

func part2(words []string) int {
	graph := graphIt(words)
	count := 0

	for r, _ := range graph {
		rlimit := len(graph)
		for c, _ := range graph[r] {
			climit := len(graph[r])
			if graph[r][c] == "A" {
				//check bounds
				if r >= 1 && r <= rlimit-2 && c >= 1 && c <= climit-2 {
					if ((graph[r-1][c-1] == "M" && graph[r+1][c+1] == "S") || (graph[r-1][c-1] == "S" && graph[r+1][c+1] == "M")) && ((graph[r-1][c+1] == "S" && graph[r+1][c-1] == "M") || (graph[r-1][c+1] == "M" && graph[r+1][c-1] == "S")) {
						count++
					}

				}
			}
		}

	}
	return count
}

func graphIt(words []string) [][]string {
	graph := [][]string{}

	for _, row := range words {
		currRow := []string{}
		for _, letter := range row {
			currRow = append(currRow, string(letter))

		}
		graph = append(graph, currRow)
	}
	return graph
}
