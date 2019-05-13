package main_test

import (
	"testing"
)

func Test_Add_2_and_3_should_be_5(t *testing.T) {
	expected := 5

	num1 := 2
	num2 := 3

	actual := num1 + num2
	if expected != actual {
		t.Errorf("Expected %d should be Actual %d", expected, actual)
	}
}
