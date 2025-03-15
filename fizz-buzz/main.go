/*
Write a program that prints the numbers from 1 to 100.
But for multiples of three print “Fizz” instead of the number and for the multiples of five print “Buzz”.
For numbers which are multiples of both three and five print “FizzBuzz”.

Sample output:

1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
... etc up to 100
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fizzBuzz() []string {
	var output []string

	for i := 1; i <= 100; i++ {
		output = append(output, strconv.Itoa(i))
	}

	return output
}

func main() {
	result := fizzBuzz()
	fmt.Print(strings.Join(result, "\n"))
}
