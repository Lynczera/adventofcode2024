package main

import (
	"bufio"
	"os"
	"slices"
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

	println(part1(rules, updates))
	println(part2(rules, updates))
}
func part1(rules []string, updates []string) int {
	count := 0
	rulesmap := buildRules(rules)

	for _, u := range updates {
		rulesarr := strings.Split(u, ",")
		valid := true
		for j := 0; j < len(rulesarr); j++ {
			curr_rules, _ := rulesmap[rulesarr[j]]
			for k := j + 1; k < len(rulesarr); k++ {
				if slices.Contains(curr_rules, rulesarr[k]) {
					valid = false
					break

				} else {
					continue

				}

			}
			if !valid {
				break
			}

		}

		if valid {
			parsedVal, _ := strconv.Atoi(rulesarr[len(rulesarr)/2])
			count += parsedVal
		}

	}
	return count

}

func part2(rules []string, updates []string) int {
	count := 0
	countfixed := 0
	rulesmap := buildRules(rules)

	for _, u := range updates {
		// curr_rules := rulesmap[u]
		rulesarr := strings.Split(u, ",")
		valid := true
		for j := 0; j < len(rulesarr); j++ {
			curr_rules, _ := rulesmap[rulesarr[j]]
			for k := j + 1; k < len(rulesarr); k++ {
				if slices.Contains(curr_rules, rulesarr[k]) {
					valid = false
					break

				} else {
					continue

				}

			}
			if !valid {
				break
			}

		}

		if valid {
			parsedVal, _ := strconv.Atoi(rulesarr[len(rulesarr)/2])
			count += parsedVal
		} else {
			fixedUpdate := fixUpdate(rulesarr, rulesmap)
			parsedfixed, _ := strconv.Atoi(fixedUpdate[len(fixedUpdate)/2])
			countfixed += parsedfixed

		}

	}
	return countfixed

}

// 5397,
func fixUpdate(ups []string, rulesmap map[string][]string) []string {

	// curr_rules := rulesmap[u]
	reset := true
	for reset {
		valid := true
		for j := 0; j < len(ups); j++ {
			curr_rules, _ := rulesmap[ups[j]]
			for k := j + 1; k < len(ups); k++ {
				if slices.Contains(curr_rules, ups[k]) {
					a := ups[j]
					b := ups[k]
					ups[j] = b
					ups[k] = a
					valid = false
					break

				} else {
					continue

				}

			}

		}
		if valid {
			reset = false
		}
	}
	return ups

}
func buildRules(rules []string) map[string][]string {

	res := make(map[string][]string)

	//build a reverse rule -> 75 = [23, 24] -> all the numbers that 75 has to go after / all the numbers that must come before 75
	for _, rule := range rules {
		splitrule := strings.Split(rule, "|")
		key := splitrule[1]
		currRules, found := res[key]
		if found {
			currRules = append(currRules, splitrule[0])
			res[key] = currRules

		} else {
			res[key] = []string{splitrule[0]}
		}

	}
	return res

}
