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

## Running Tests

To run tests for all katas:

```bash
./test.sh
```


To run tests for a specific kata:

```bash
./test.sh tennis
```

## Creating a New Kata

To create a new kata project structure:

```bash
./create-kata.sh kata-name
```

This will create a new directory with the basic Go project structure including:
- `go.mod`
- `kata-name.go`
- `kata-name_test.go`

## References

For more kata descriptions and ideas, visit [Samman Coaching Katas](https://sammancoaching.org/kata_descriptions/index.html)
