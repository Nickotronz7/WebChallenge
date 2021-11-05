package main

import (
	"fmt"
	"strings"
)





func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	gameOn := true
	currentPlayer := "X"

	for gameOn {
		printBoard(board)
		fmt.Printf("It's %s's turn\n", currentPlayer)
		move := getMove(currentPlayer)
		makeMove(board, move, currentPlayer)

		if win(board, currentPlayer) {
			printBoard(board)
			fmt.Printf("%s wins!\n", currentPlayer)
			gameOn = false
		} else if fullBoard(board) {
			printBoard(board)
			fmt.Println("It's a tie!")
			gameOn = false
		} else {
			switchPlayer(&currentPlayer)
		}
	}
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func getMove(player string) (move int) {
	fmt.Printf("%s, what is your move? ", player)
	fmt.Scanf("%d", &move)
	return
}

func makeMove(board [][]string, move int, player string) {
	if move < 1 || move > 9 {
		fmt.Println("Invalid move")
		return
	}
	move--
	row := move / 3
	column := move % 3
	board[row][column] = player
}

func win(board [][]string, player string) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		} 
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}




func fullBoard(board [][]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

func switchPlayer(player *string) {
	if *player == "X" {
		*player = "O"
	} else {
		*player = "X"
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}