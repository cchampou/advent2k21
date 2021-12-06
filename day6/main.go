package day6

import (
	"../utils"
	"fmt"
	"strings"
)

func Run() {
	swap := make([]int, 9)
	data := utils.ConvertLinesToInt(strings.Split(utils.ReadFileInput("day6/input")[0], ","))
	for _, value := range data {
		swap[value] = swap[value] + 1
	}
	for day := 0; day < 256; day++ {
		reset := 0
		for i := 0; i < 9; i++ {
			if i == 0 {
				reset = swap[0]
			} else {
				swap[i - 1] = swap[i]
			}
		}
		swap[6] = swap[6] + reset
		swap[8] = reset
	}
	fmt.Println(count(swap))
}

func count(swap []int) int {
	score := 0
	for _, value := range swap {
		score += value
	}
	return score
}