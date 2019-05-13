package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Change_Input_Cost_215_And_Amount_300_Should_Be_50_20_10_5(t *testing.T) {
	expected := []int{50, 20, 10, 5}
	cashProfile := []int{1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}
	actual := calculateChange(215, 300, cashProfile)
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("Expected %d should be Actual %d", expected, actual)
	}
}

func Test_Change_Input_Cost_486_And_Amount_600_Should_Be_100_10_2_2(t *testing.T) {
	expected := []int{100, 10, 2, 2}
	cashProfile := []int{1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}
	actual := calculateChange(486, 600, cashProfile)
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("Expected %d should be Actual %d", expected, actual)
	}
}

func Test_Change_Input_Cost_12_And_Amount_400_Should_Be_200_100_50_20_10_5_2_1(t *testing.T) {
	expected := []int{200, 100, 50, 20, 10, 5, 2, 1}
	cashProfile := []int{1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}
	actual := calculateChange(12, 400, cashProfile)
	if diff := cmp.Diff(expected, actual); diff != "" {
		t.Errorf("Expected %d should be Actual %d", expected, actual)
	}
}
