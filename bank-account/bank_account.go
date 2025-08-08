/*

Bank Account
Write a class named BankAccount that implements the following public interface:

public interface BankAccount
{
    void deposit(int amount)
    void withdraw(int amount)
    void printStatement()
}

Example statement
When you call the ‘printStatement’ method, something like the following is printed on standard output:

Date       || Amount || Balance
2012-01-14 || -500   || 2500
2012-01-13 || 2000   || 3000
2012-01-10 || 1000   || 1000
This example statement shows one withdrawal on 14th January 2012, and two deposits on 13th and 10th January respectively.

Notes
You cannot change the public interface of the BankAccount
We’re using ints to represent money, which in general may not be the best idea. In a real system, we would always use a datatype with guaranteed arbitrary precision, but doing so here would distract from the main purpose of the exercise.
Don’t worry about matching the exact formatting of the bank statement, the important thing is to print a table that has column headings and which orders transactions by date.

*/

package main

import "fmt"

type bankAccount struct {
    balance float32
}

func BankAccount() *bankAccount {
	return &bankAccount{}
}

func (bankAccount *bankAccount) deposit(amount int) {
 //
}

func (bankAccount *bankAccount) withdraw(amount int) {
 //
}

func (bankAccount *bankAccount) printStatement() {
	fmt.Print("Date       || Amount || Balance")
}
