/*
- Leap year -

Write a function that returns true or false depending on whether its input integer is a leap year or not.
A leap year is defined as one that is divisible by 4, but is not otherwise divisible by 100 unless it is also
divisible by 400.

For example, 2001 is a typical common year and 1996 is a typical leap year, whereas 1900 is an atypical common
year and 2000 is an atypical leap year.
*/

package main

import "errors"

func isLeapYear(year int) (bool, error) {
	if year < 1 {
		return false, errors.New("years less than 1 are unsupported")
	}
	return year%4 == 0 && (year%100 != 0 || year%400 == 0), nil
}

func main() {

}
