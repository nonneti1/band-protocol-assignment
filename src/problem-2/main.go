package main

import "fmt"

func problem_2(positions []int, k int) int {
	n := len(positions)

	// Variable to track the maximum chickens protected
	maxProtected := 0 //1,2,
	// Variable to track the current starting position of the roof
	currentStart := 0 //1,2

	for i := 0; i < n; i++ {
		if positions[i]-positions[currentStart] > k {
			continue
		}
		fmt.Printf("%d ", currentStart)
		fmt.Println(maxProtected)
		// Check if the current chicken is within the roof's reach
		fmt.Println(positions[i], " ", positions[currentStart])
		if positions[i]-positions[currentStart] <= k {
			// Update maxProtected if more chickens are covered
			maxProtected = max(maxProtected, i-currentStart+1)
		} else {
			// Move the roof to the next chicken if it cannot cover the current one
			currentStart++
		}
	}

	return maxProtected
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	res := problem_2([]int{2, 5, 10, 12, 15}, 5)
	fmt.Println(res)
}
