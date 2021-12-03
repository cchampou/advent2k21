package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadFileInput(filepath string) []string {
	var data []string
	f, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Open file errorL %v", err)
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("Failed to close the file: %v", err)
		}
	}()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		data = append(data, sc.Text())
	}
	return data
}
