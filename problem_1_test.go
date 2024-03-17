package main

import (
	"testing"
)

func TestProblem_1(t *testing.T) {
	testCase := []struct {
		input  string
		output string
	}{
		{"SRSSRRR", "Good boy"},
		{"RSSRR", "Bad boy"},
		{"SSSRRRRS", "Bad boy"},
		{"SSRR", "Good boy"},
		{"SRRSSR", "Bad boy"},
		{"S", "Bad boy"},
		{"R", "Bad boy"},
		{"SS", "Bad boy"},
		{"RR", "Bad boy"},
		{"ABCDEFG", "Bad boy"},
		{"2077CP", "Bad boy"},
		{` `, "Bad boy"},
	}
	for _, tc := range testCase {
		t.Run(tc.input, func(t *testing.T) {
			actual := problem_1(tc.input)
			if actual != tc.output {
				t.Errorf("for %s: got %s, expected %s", tc.input, actual, tc.output)
			}
		})
	}
}
