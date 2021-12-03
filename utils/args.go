package utils

import (
	"log"
	"os"
)

func CheckArgs() {
	if len(os.Args[1:]) < 1 {
		log.Fatalln("To few arguments were provided. There should be one.")
	}
}
