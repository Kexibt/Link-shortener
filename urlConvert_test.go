package main

import (
	"math"
	"strconv"
	"testing"
)

func TestConvertToShort(t *testing.T) {
	a := 1
	b := convertToShort(a)
	if b != "00000000001" {
		t.Log("10 should be '00000000001', but received " + b)
		t.Fail()
	}

	a = 10
	b = convertToShort(a)
	if b != "0000000000A" {
		t.Log("10 should be '0000000000A', but received " + b)
		t.Fail()
	}

	a = 2567
	b = convertToShort(a)
	if b != "000000000dk" {
		t.Log("10 should be '000000000dk', but received " + b)
		t.Fail()
	}
}

func TestConvertToID(t *testing.T) {
	a := 0
	b := "00000000"
	c := convertToID(b)
	if a != c {
		t.Log("00000000 should be 0, but received " + strconv.Itoa(c))
		t.Fail()
	}

	a = 1
	b = "00000001"
	c = convertToID(b)
	if a != c {
		t.Log("00000001 should be 1, but received " + strconv.Itoa(c))
		t.Fail()
	}

	a = 2567
	b = convertToShort(a)
	c = convertToID(b)
	if a != c {
		t.Log("2567 should be 2567, but received " + strconv.Itoa(c))
		t.Fail()
	}

	b = "zzzzzzzzzz"
	c = convertToID(b)
	if c < 9*int(math.Pow10(17)) {
		t.Log("something wrong.. with sleeping test. expected: > "+strconv.Itoa(9*int(math.Pow10(17))),
			", but received: "+strconv.Itoa(c))
		t.Fail()
	}
}
