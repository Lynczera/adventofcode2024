package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	target := "rules"

	var updates []string
	var rules []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			target = "updates"
			continue
		}
		if target == "rules" {
			rules = append(rules, line)
		}
		if target == "updates" {
			updates = append(updates, line)
		}

	}
}
func part1(words []string) {

}
