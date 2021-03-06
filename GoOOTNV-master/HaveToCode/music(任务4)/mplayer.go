package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./library"
	"./mp"
)

var lib *library.MusicManager = library.NewMusicManager()
var id int = 1
var play mp.Player

func main() {
	fmt.Println(`
                Enter following commands to control the player:
                lib list -- View the existing music lib
                lib add <name><artist><source><type> -- Add a music to the music lib
                lib remove <id> -- Remove the specified music from the lib
                play <name> -- Play the specified music
        `)

	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}

	}
}
func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	mp.Paly(e.Name, "MP3")

}
func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		lib.List()
	case "add":
		//fmt.Println("len(tokens)", len(tokens))
		if len(tokens) == 6 {

			lib.Add(&library.Music{strconv.Itoa(id),
				tokens[2], tokens[3], tokens[4], tokens[5]})
			id++
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")

		}
	case "remove":
		if len(tokens) == 3 {

			lib.Remove(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <id>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}
