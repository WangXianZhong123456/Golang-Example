package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("An errer occurred:%s", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("Heool,%s!What can I do for you?\n", name)
	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred:%s\n", err)
			continue
		}
		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye", "q":
			fmt.Println("Bye Bye")
			time.Sleep(time.Duration(time.Second * 2))
			os.Exit(0)
		default:
			fmt.Println("Sorry I dont know what you said")
		}
	}
}
