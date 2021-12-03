package day1

import (
	"../utils"
	"fmt"
)

func getSum(a []int, b int) int {
	return a[b-2] + a[b-1] + a[b]
}

func Run() {
	measures := utils.ConvertLinesToInt(utils.ReadFileInput("day1/input"))

	total := 0
	for index, _ := range measures {
		if index > 2 && getSum(measures, index) > getSum(measures, index-1) {
			total++
		}
	}

	fmt.Println(total)
}