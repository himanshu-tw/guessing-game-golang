# Go Number Guessing Game

A simple, interactive command-line number guessing game built using Go (Golang). 

This project is a solution to the [roadmap.sh Number Guessing Game challenge](https://roadmap.sh/projects/number-guessing-game).

## Features

- **Multiple Difficulty Levels:** Choose between Easy (10 chances), Medium (5 chances), or Hard (3 chances).
- **Dynamic Feedback:** Tells you if your guess is too high or too low.
- **Performance Tracking:** Displays the number of attempts it took to guess the correct number.

## Getting Started

### Prerequisites

Make sure you have [Go](https://go.dev) installed on your system (version 1.16 or higher recommended).

### Installation & Running

1. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/himanshu-tw/guessing-game-golang.git
   ```

2. Navigate into the project directory:
   ```bash
   cd guessing-game-golang
   ```

3. Run the game directly:
   ```bash
   go run main.go
   ```

## How to Play

1. Run the application to start the game.
2. Select your desired difficulty level by typing `1`, `2`, or `3`.
3. The system will think of a random number between 1 and 100.
4. Input your guesses. The game will guide you by telling you if the secret number is higher or lower.
5. Win by guessing the number before running out of chances!
