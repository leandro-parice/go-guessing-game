# Go Guessing Game 🎯

My first Go project as part of my studies with Rocketseat's Go course.

## About

This is a simple number guessing game built in Go where the player tries to guess a randomly generated number between 0 and 100. It's my first hands-on project to learn Go fundamentals including:

- Basic syntax and data types
- User input handling with `bufio.Scanner`
- String manipulation and conversion
- Control structures (loops, conditionals)
- Arrays and slices
- Error handling
- Random number generation

## How to Play

1. The game generates a random number between 0 and 100
2. You have up to 10 attempts to guess the correct number
3. After each guess, the game tells you if your number is too high or too low
4. If you guess correctly, the game shows all your attempts
5. If you use all 10 attempts without guessing correctly, the game reveals the answer

## Features

- Input validation (only accepts integer numbers)
- Smart attempt tracking (only shows actual guesses made, not empty slots)
- Clear feedback for each guess
- Portuguese language interface

## How to Run

Make sure you have Go installed on your system, then:

```bash
go run main.go
```

## Learning Journey

This project helped me understand:
- Go's type system and error handling patterns
- The difference between arrays and slices
- How to work with user input in Go
- Go's approach to random number generation
- String manipulation and conversion functions

---

**Course**: Rocketseat Go  
**Status**: Learning Project #1  
**Language**: Go 1.21+