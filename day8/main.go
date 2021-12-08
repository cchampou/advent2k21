package day8

import (
	"../utils"
	"fmt"
	"github.com/thoas/go-funk"
	"strconv"
	"strings"
)

func Run() {
	score := 0
	entries := parser()
	for _, current := range entries {
		number, _ := strconv.Atoi(current.result)
		score += number
	}
	fmt.Println(score)
}

func parser() []Entry {
	var output []Entry
	lines := utils.ReadFileInput("day8/input")
	for i := range lines {
		splitLine := strings.Split(lines[i], "|")
		newEntry := Entry{config: make([]string, 10)}
		newEntry.patterns = strings.Split(splitLine[0], " ")
		newEntry.outputs = strings.Split(splitLine[1], " ")[1:]
		configured := buildConfig(newEntry, 1000)
		resolved := resolve(configured)
		output = append(output, resolved)
	}
	return output
}

func countUniq(entries []Entry) int {
	score := 0
	for _, entry := range entries {
		for _, output := range entry.outputs {
			outputLen := len(output)
			if outputLen == 2 || outputLen == 3 || outputLen == 4 || outputLen == 7 {
				score++
			}
		}
	}
	return score
}

func buildConfig(entry Entry, max int) Entry {
	if len(entry.patterns) == 0 || max == 0 {
		return entry
	}
	current := entry.patterns[0]
	switch len(current) {
	case 2:
		entry.config[1] = current
		return buildConfig(rotate(entry), max - 1)
	case 3:
		entry.config[7] = current
		return buildConfig(rotate(entry), max - 1)
	case 4:
		entry.config[4] = current
		return buildConfig(rotate(entry), max - 1)
	case 7:
		entry.config[8] = current
		return buildConfig(rotate(entry), max - 1)
	}

	if len(entry.config[4]) > 0 {
		diff4, _ := funk.DifferenceInt(toInts(entry.config[4]), toInts(current))
		if len(diff4) == 0 && len(current) == 6 {
			entry.config[9] = current
			return buildConfig(rotate(entry), max - 1)
		}
	}
	if len(entry.config[9]) > 0 && len(entry.config[7]) > 0 {
		diff7, _ := funk.DifferenceInt(toInts(entry.config[7]), toInts(current))
		if len(diff7) == 0 && len(current) == 6 {
			entry.config[0] = current
			return buildConfig(rotate(entry), max - 1)
		}
	}
	if len(entry.config[9]) > 0 && len(entry.config[0]) > 0 {
		if len(current) == 6 {
			entry.config[6] = current
			return buildConfig(rotate(entry), max - 1)
		}
	}
	if len(entry.config[7]) > 0 {
		diff7new, _ := funk.DifferenceInt(toInts(entry.config[7]), toInts(current))
		if len(current) == 5 && len(diff7new) == 0 {
			entry.config[3] = current
			return buildConfig(rotate(entry), max - 1)
		}
	}
	if len(entry.config[4]) > 0 && len(entry.config[3]) > 0 {
		diff4new, _ := funk.DifferenceInt(toInts(entry.config[4]), toInts(current))
		if len(current) == 5 && len(diff4new) == 2 {
			entry.config[2] = current
			return buildConfig(rotate(entry), max - 1)
		}
		if len(current) == 5 && len(diff4new) == 1 {
			entry.config[5] = current
			return buildConfig(rotate(entry), max - 1)
		}
	}
	return buildConfig(next(entry), max - 1)
}

func resolve(entry Entry) Entry {
	entry.result = ""
	for _, current := range entry.outputs {
		for j, config := range entry.config {
			if match(current, config) && len(current) == len(config) {
				entry.result = entry.result + strconv.Itoa(j)
			}
		}
	}
	return entry
}

func match(a string, b string) bool {
	result := true
	for _, current := range a {
		if strings.Index(b, string(current)) == -1 {
			result = false
		}
	}
	return result
}

func next(entry Entry) Entry {
	return Entry{config: entry.config, patterns: append(entry.patterns[1:], entry.patterns[0]), outputs: entry.outputs}
}

func rotate(entry Entry) Entry {
	return Entry{config: entry.config, patterns: entry.patterns[1:], outputs: entry.outputs}
}

func toInts(input string) []int {
	var output []int
	for _, char := range input {
		output = append(output, int(char))
	}
	return output
}