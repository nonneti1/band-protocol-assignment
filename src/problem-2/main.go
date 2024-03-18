package main

import "fmt"

func problem_2(roofLength int, chickenPositions []int) int {
	var currentMaxProtected, currentRoofPosition int

	chickenAmount := len(chickenPositions)

	for i := 0; i < chickenAmount; i++ {
		fmt.Println("Index ", i)
		fmt.Printf("current chicken position : %d ,current roof position %d\n", chickenPositions[i], currentRoofPosition)

		if chickenPositions[i]-chickenPositions[currentRoofPosition] < roofLength {
			currentMaxProtected = max(currentMaxProtected, i-currentRoofPosition+1)
			fmt.Println("Update max protected", currentMaxProtected)
		} else {
			currentRoofPosition++
			fmt.Println("update roof position", currentRoofPosition)
		}
	}

	return currentMaxProtected
}

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
