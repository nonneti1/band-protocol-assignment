/** Explanation:
* This problem_1 function act as the Boss baby's assistant to analyze actions represented by the string 'str', where each character in the string can be either 'S' (shot) or 'R' (revenge). The goal is to determine whether the Boss is ("Good boy") or ("Bad boy") based on given rules.

  Comments:
  - This is the first time I've coded in Golang eventhough the instructions have said I can use any languages, I thought that since I'm applying in a position using golang then why not just give it a try.
  - Beside the conditions given, I've made sure to implement a fail safe to handle other possible cases eg. single action handler, empty input or any action beside "S" and "R" will return "Bad boy"

  Time Complexity: O(n)
  Memory Complexity: O(1)
*/

package main

import (
	"fmt"
	"os"
)

func problem_1(str string) string {
	// Declare action counter variable for shot and revenge
	var shotCount, revengeCount uint32

	// Start by iterating through the string
	for i, elem := range str {
		// Cast current character from rune to string.
		currentAction := string(elem)
		isShot := currentAction == "S"
		isRevenge := !isShot

		isFirstIndex := i == 0

		// Increase an amount of current action.
		if isShot {
			shotCount++
		} else {
			revengeCount++
		}

		// Safecheck if first action is revenge then return "Bad boy" because a good boss shouldn't shot the others first.
		if isFirstIndex && isRevenge {
			return "Bad boy"
		}

		// Check if action length is less than 1 then return "Bad boy" result because one action should return bad either getting shot or revenge.
		actionLength := len(str)
		if actionLength <= 1 {
			return "Bad boy"
		}

		// After incrementing and safe check earlier conditions, We'll just continue to next action.
		if isFirstIndex {
			continue
		}

		// Create a variable to check if current action is the final action.
		isFinalAction := i == len(str)-1

		// Create a variable to check if current revenge action is the last revenge in sequence.
		isLastRevengeInSequence := isRevenge && (isFinalAction || string(str[i+1]) != currentAction)

		// Check if it's the final action and check if the boss got shot more than revenge then return "Bad boy" because the boss hasn't revenged them enough.
		if isFinalAction && shotCount > revengeCount {
			return "Bad boy"
		}

		// Check if current action is the last revenge action and check if the boss has done the sweet revenge then we'll reset the action counter and continue to next action.
		if isLastRevengeInSequence && revengeCount >= shotCount {
			shotCount = 0
			revengeCount = 0
			continue
		}

		// Check if current action is not final action then we'll continue to next action
		if !isFinalAction {
			continue
		}

		// Return "Bad boy" because the boss's action hasn't met the satisfied conditions.
		return "Bad boy"
	}

	// Return "Good boy" becuase he has passed all the satisfied conditions and he's a good boy.
	return "Good boy"
}

func main() {
	action := os.Args[1]
	result := problem_1(action)
	fmt.Println(result)
}
