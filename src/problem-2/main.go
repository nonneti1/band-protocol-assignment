/*
* Explanation:
* This problem_2 function act as a calculator to find how many chickens can superman protected by using a roof.

	Comments:
	- I've spent too much time on thinking the solution but in the end I'm satisfied by using only one loop.

	Time Complexity: O(n)
	Memory Complexity: O(1)
*/
package main

import "fmt"

func problem_2(roofLength int, chickenPositions []int) int {
	// Declare a tracker variable for current amount of protectable chicken and current roof position index
	var currentMaxProtected, currentRoofPosition int

	// Get chicken amount from the length of array
	chickenAmount := len(chickenPositions)

	// Start by iterating through each chicken
	for i := 0; i < chickenAmount; i++ {
		fmt.Println("Index ", i)
		fmt.Printf("current chicken position : %d ,current roof position %d\n", chickenPositions[i], currentRoofPosition)

		// Check if current chicken can be protected under the roof from current roof position index
		if chickenPositions[i]-chickenPositions[currentRoofPosition] < roofLength {
			// Update tracker maximum number of protected chickens by finding the maximum value between currentMaxProtected and i-currentRoofPosition to find how far away the chicken is from the roof position and add 1 because the chicken at currentRoofPosition is also being protected
			currentMaxProtected = max(currentMaxProtected, i-currentRoofPosition+1)
			fmt.Println("Update max protected", currentMaxProtected)
		} else {
			// Move the roof to the next index position
			currentRoofPosition++
			fmt.Println("update roof position", currentRoofPosition)
		}
	}

	return currentMaxProtected
}

// Find the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var chickenAmount, roofLength int

	fmt.Print("Enter chicken amount : ")
	fmt.Scanf("%d\n", &chickenAmount)

	fmt.Print("Enter roof length : ")
	fmt.Scanf("%d\n", &roofLength)

	chickenPosition := make([]int, chickenAmount)

	for i := 0; i < chickenAmount; i++ {
		fmt.Printf("Enter chicken No.%d position : ", i+1)
		fmt.Scanf("%d\n", &chickenPosition[i])
	}

	result := problem_2(roofLength, chickenPosition)

	fmt.Println(result)
}
