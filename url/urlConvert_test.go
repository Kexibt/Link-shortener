package url

import (
	"math"
	"strconv"
	"testing"
)

func TestConvertToShort(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1, "00000000001"},
		{10, "0000000000A"},
		{2567, "000000000dk"},
	}

	for _, test := range tests {
		id, str := test.input, test.want
		str1 := ConvertToShort(id)
		if str != str1 {
			t.Logf("Test 'Convert to short' failed successfully. \n\tExpected: %s. Actual: %s\n", str, str1)
			t.Fail()
		}
	}
}

func TestConvertToID(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"00000000001", 1},
		{"0000000000A", 10},
		{"000000000dk", 2567},
	}

	for _, test := range tests {
		input, want := test.input, test.want
		got := ConvertToID(input)
		if got != want {
			t.Logf("Test 'Convert to short' failed successfully. \n\tExpected: %d. Actual: %d\n", got, want)
			t.Fail()
		}
	}

	// bonus
	b := "zzzzzzzzzz"
	c := ConvertToID(b)
	if c < 9*int(math.Pow10(17)) {
		t.Log("something wrong.. it was sleeping test. expected: > "+strconv.Itoa(9*int(math.Pow10(17))),
			", but received: "+strconv.Itoa(c))
		t.Fail()
	}
}
