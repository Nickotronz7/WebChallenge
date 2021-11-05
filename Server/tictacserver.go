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
	GameOn bool       `json:"gameOn"`
	Board  [][]string `json:"board"`
}

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Print("Please provide a port number")
	}

	PORT := ":" + arguments[1]
	// PORT := ":65000"

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
		makeMove(board, currentPlayer)
		printBoard(board)
		checkGameState(board, &gameOn, "X")

		if gameOn {
			gameState, _ := createJson(board, gameOn)
			fmt.Fprintf(c, string(gameState)+"\n")
			fmt.Println("Waiting for other player...")
			// se lee la respuesta del otro jugador
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			
			var game Game
			netData = netData[:len(netData)-1]
			err = json.Unmarshal([]byte(netData), &game)
			
			if err != nil {
				fmt.Println(err)
			}
			
			board = game.Board
			gameOn = game.GameOn
			checkGameState(board, &gameOn, "O")

		} else {
			printBoard(board)
			gameState, _ := createJson(board, gameOn)
			fmt.Fprintf(c, string(gameState)+"\n")
			return
		}
	}
}

func checkGameState(board [][]string, gameState *bool, currentPlayer string) {
	if win(board, currentPlayer) {
		// printBoard(board)
		fmt.Printf("%s wins!\n", currentPlayer)
		*gameState = false
	} else if fullBoard(board) {
		printBoard(board)
		fmt.Println("It's a tie!")
		*gameState = false
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

func getMove(player string) (move int) {
	fmt.Printf("%s, what is your move? \n", player)
	fmt.Scanf("%d", &move)
	return move
}

func makeMove(board [][]string, player string) {
	move := 0
	move = getMove("X")
	for {
		if (move <= 0) || (move >= 9) {
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
