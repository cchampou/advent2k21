package utils

import "testing"

func TestConvertLinesToIntWithEasyData(t *testing.T) {
	input := []string{"-1", "0", "1"}
	want := []int{-1,0,1}
	res := ConvertLinesToInt(input)
	for i, value := range res {
		if want[i] != value {
			t.Fatalf("Got %d instead of %d", value, want[i])
		}
	}
}

func TestConvertLinesToIntWithEmptyArray(t *testing.T) {
	var input []string
	res := ConvertLinesToInt(input)
	if len(res) > 0 {
		t.Fatalf("Output array should have zero length")
	}
}
