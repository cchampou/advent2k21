package day4

import "testing"

func TestGenerateCalledNumbers(t *testing.T) {
	input := "1,2,3,4,5"
	want := []int{1,2,3,4,5}
	result := GenerateCalledNumbers(input)
	for index, value := range result {
		if want[index] != value {
			t.Fatalf("GenerateCalledNumbers(\"%v\") should give %v", input, want)
		}
	}
}

func TestGenerateCalledNumbersWithEmpty(t *testing.T) {
	input := ""
	var want []int
	result := GenerateCalledNumbers(input)
	for index, value := range result {
		if want[index] != value {
			t.Fatalf("GenerateCalledNumbers(\"%v\") should give %v", input, want)
		}
	}
}