package main

import (
	"./day1"
	"./day2"
	"./day3"
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
	default:
		log.Fatalln("Unknown day number")
	}
}
