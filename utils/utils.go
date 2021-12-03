package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileInput() []string {
	var data []string
	f, err := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Open file errorL %v", err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		data = append(data, sc.Text())
	}
	return data
}