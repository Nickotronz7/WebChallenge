// create ticktactoe game with a board of 3x3 and a player to play with X or O with comand line input

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
	Move  int        `json:"move"`
	Board [][]string `json:"board"`
}

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Print("Please provide a port number")
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	fmt.Println("Game Started waiting for opponent...")

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("An opponent has been found...!" + c.RemoteAddr().String())

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	gameOn := true
	currentPlayer := "X"

	for gameOn {
		printBoard(board)
		fmt.Printf("It's %s's turn\n", currentPlayer)
		move := getMove(currentPlayer)
		makeMove(board, move, currentPlayer)

		checkGameState(board, &gameOn, &currentPlayer)

		if gameOn {
			gameState, _ := createJson(board)
			fmt.Fprintf(c, string(gameState)+"\n")

			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			var response Game
			netData = netData[:len(netData)-1]
			err = json.Unmarshal([]byte(netData), &response)

			if err != nil {
				fmt.Println(err)
			}

			board = response.Board
			makeMove(board, response.Move, currentPlayer)

			checkGameState(board, &gameOn, &currentPlayer)
		}

	}
}

func checkGameState(board [][]string, gameState *bool, currentPlayer *string) {
	if win(board, *currentPlayer) {
		printBoard(board)
		fmt.Printf("%s wins!\n", *currentPlayer)
		*gameState = false
	} else if fullBoard(board) {
		printBoard(board)
		fmt.Println("It's a tie!")
		*gameState = false
	} else {
		switchPlayer(currentPlayer)
	}
}

func createJson(board [][]string) ([]byte, error) {

	jStructure := Game{Move: 0, Board: board}
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

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }
