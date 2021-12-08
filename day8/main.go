package day8

import (
	"../utils"
	"fmt"
	"strings"
)

func Run() {
	entries := parser()
	fmt.Println(countUniq(entries))
}

func parser() []Entry {
	var output []Entry
	lines := utils.ReadFileInput("day8/input")
	for i := range lines {
		splitLine := strings.Split(lines[i], "|")
		newEntry := Entry{}
		newEntry.patterns = strings.Split(splitLine[0], " ")
		newEntry.outputs = strings.Split(splitLine[1], " ")
		output = append(output, buildConfig(newEntry))
		fmt.Println(newEntry)
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

func buildConfig(entry Entry) Entry {
	entry.config = make([]string, 10)
	for _, pattern := range entry.patterns {
		switch len(pattern) {
		case 2:
			entry.config[1] = pattern
		case 3:
			entry.config[7] = pattern
		case 4:
			entry.config[4] = pattern
		case 7:
			entry.config[8] = pattern
		}
	}
	return entry
}