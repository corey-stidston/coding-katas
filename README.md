# Coding Katas

## What are Coding Katas?

A coding kata is a programming exercise that helps developers hone their skills through practice and repetition. The term "kata" comes from martial arts, where practitioners perform detailed choreographed patterns of movements to perfect their technique.

In software development, katas help you:
- Practice Test-Driven Development (TDD)
- Practice building in small steps
- Improve problem-solving skills
- Learn new programming patterns
- Experiment with different approaches to the same problem
- Build muscle memory for good coding practices

## Katas in this Project

Currently implemented katas:

1. **Hello World** - A simple starter kata that prints "Hello, World!"
2. **Leap Year** - Determine if a given year is a leap year
3. **FizzBuzz** - Print numbers 1 to 100, replacing multiples of 3 with "Fizz", multiples of 5 with "Buzz", and multiples of both with "FizzBuzz"

## Running Tests

To run tests for all katas:

```bash
./run-tests.sh
```

## Creating a New Kata

To create a new kata project structure:

```bash
./create-kata.sh kata-name
```

This will create a new directory with the basic Go project structure including:
- `go.mod`
- `main.go`
- `main_test.go`

## References

For more kata descriptions and ideas, visit [Samman Coaching Katas](https://sammancoaching.org/kata_descriptions/index.html)
