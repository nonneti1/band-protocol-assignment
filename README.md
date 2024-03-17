# band-protocol-problem_1
Assignment submission for Band Protocol

Problem 1: Boss Baby's Revenge
Description:
Boss Baby, known for his cool and clever ways, deals with teasing from the neighborhood kids who shoot
water guns at his house. In response, Boss Baby seeks revenge by shooting at least one shot back, but only
if the kids have already shot at him first. As Boss Baby's assistant, your task is to check if Boss Baby has
sought revenge for every shot aimed at him at least once and hasn't initiated any shooting himself.

<p>Input: A string (S, 1 <= len(S) <= 1,000,000) containing characters 'S' and 'R', where 'S' represents a shot and 'R' represents revenge. The string represents the sequence of events for one day.</p>
<p>Output: Return "Good boy" if all shots have been revenged at least once and Boss Baby doesn’t initiate any shots to the neighborhood kids. Return "Bad boy" if these conditions are not satisfied.</p>

## Manually run the program
<p>After cloning the project, Launch the terminal in current directory and use command</p>
    
    go run problem_1.go <string input here>
    
Where string input contains "S" and "R"
eg. "SRSR", "SSRRR", "SSSR"

## Run unit test
I've also added unit test for other cases beside given in the example

    go test -v    
<p>All test cases should return passed</p>
