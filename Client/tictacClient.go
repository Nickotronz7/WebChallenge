package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

type Game struct {
	GameOn bool       `json:"gameOn"`
	Board  [][]string `json:"board"`
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	var game Game
	gameOn := true
	fmt.Println("Waiting for rival movement")
	for gameOn {
		message, _ := bufio.NewReader(c).ReadString('\n')
		err = json.Unmarshal([]byte(message[:len(message)-1]), &game)

		printBoard(game.Board)
		gameOn = game.GameOn

		checkGameState(game.Board, &gameOn, "X")

		if !gameOn {
			break
		}

		makeMove(game.Board, "O")

		printBoard(game.Board)
		checkGameState(game.Board, &gameOn, "O")

		if !gameOn {
			gameState, _ := createJson(game.Board, gameOn)
			fmt.Fprintf(c, string(gameState)+"\n")
			break
		}

		gameState, _ := createJson(game.Board, gameOn)
		fmt.Fprintf(c, string(gameState)+"\n")

		fmt.Println("Waiting for other player...")
	}
	return
}

func getMove(player string) (move int) {
	fmt.Printf("%s, what is your move? \n", player)
	fmt.Scanf("%d", &move)
	return move
}

func makeMove(board [][]string, player string) {
	move := 0
	move = getMove("O")
	for {
		if (move <= 0) || (move > 9) {
			fmt.Printf(">>>Move %d out of range<<<\n", move)
			move = getMove("X")
		} else {
			move--
			row := move / 3
			column := move % 3
			if board[row][column] == "_" {
				board[row][column] = player
				break
			} else {
				fmt.Printf(">>Move %d, not allowed<<\n", move)
				move = getMove("X")
			}
		}
	}
}

func createJson(board [][]string, gameOn bool) ([]byte, error) {
	jStructure := Game{GameOn: gameOn, Board: board}
	bjStructure, err := json.Marshal(jStructure)

	if err != nil {
		return nil, err
	}

	return bjStructure, nil
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func checkGameState(board [][]string, gameState *bool, currentPlayer string) {
	if win(board, currentPlayer) {
		fmt.Printf("%s wins!\n", currentPlayer)
		*gameState = false
	} else if fullBoard(board) {
		printBoard(board)
		fmt.Println("It's a tie!")
		*gameState = false
	}
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
