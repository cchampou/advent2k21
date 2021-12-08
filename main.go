package main

import (
	"./day1"
	"./day2"
	"./day3"
	"./day4"
	"./day5"
	"./day6"
	"./day7"
	"./day8"
	"./utils"
	"log"
	"os"
)

func main() {
	utils.CheckArgs()

	switch os.Args[1] {
	case "1":
		day1.Run()
		return
	case "2":
		day2.Run()
		return
	case "3":
		day3.Run()
		return
	case "4":
		day4.Run()
		return
	case "5":
		day5.Run()
		return
	case "6":
		day6.Run()
		return
	case "7":
		day7.Run()
		return
	case "8":
		day8.Run()
		return
	default:
		log.Fatalln("Unknown day number")
	}
}
