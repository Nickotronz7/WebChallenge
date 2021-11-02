package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type Game_s struct {
	Turn  bool       `json:"turn"`
	Board [][]string `json:"board"`
}

func createGame() ([]byte, error) {
	j := Game_s{Turn: true, Board: [][]string{
		{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"},
	}}

	b, err := json.Marshal(j)

	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return b, nil
}


func main_server() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
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
	
	encodedJson, _ := createGame()

	fmt.Println("An opponent has been found...!"+c.RemoteAddr().String())
	fmt.Fprintf(c, string(encodedJson)+"\n")

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}

}
