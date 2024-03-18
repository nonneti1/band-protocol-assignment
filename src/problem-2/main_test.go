package main

import (
	"fmt"
	"testing"
)

func TestProblem_2(t *testing.T) {
	testCase := []struct {
		inputRoofLength       int
		inputChickenPositions []int
		output                int
	}{
		{1, []int{1, 10}, 1},
		{4, []int{1, 3, 11, 20}, 2},
		{5, []int{2, 5, 10, 12, 15}, 2},
		{10, []int{1, 11, 30, 34, 35, 37}, 4},
		{20, []int{1, 21, 30, 31, 45, 55, 61, 67, 99}, 3},
		{1000000000, []int{1, 2, 3, 6, 19, 21, 25, 39, 69, 99, 101, 1001, 2002, 30099, 40005, 99965, 100536, 999999, 9302945}, 19},
	}
	for _, tc := range testCase {
		testName := fmt.Sprintf("Roof length : %d and Chicken amount : %d and Chicken position : %v", tc.inputRoofLength, len(tc.inputChickenPositions), tc.inputChickenPositions)
		t.Run(testName, func(t *testing.T) {
			actual := problem_2(tc.inputRoofLength, tc.inputChickenPositions)
			if actual != tc.output {
				t.Errorf("for %s: got %d, expected %d", testName, actual, tc.output)
			}
		})
	}
}
