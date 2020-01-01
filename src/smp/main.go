package main

import (
	"bufio"
	"fmt"
	"os"
	"smp/lib"
	"smp/mp"
	"strconv"
	"strings"
)

var manager *lib.MusicManager
var id int = 1

func handleLibCommands(token []string) {
	if len(token) < 2 {
		fmt.Println("Parameter length too short")
		return
	}
	switch token[1] {
	case "list":
		for i := 0; i < manager.Len(); i++ {
			music, _ := manager.Get(i)
			fmt.Println(i+1, ":", music.Name, music.Artist, music.Source, music.Type)
		}
	case "add":
		if len(token) == 6 {
			id++
			manager.Add(&lib.Music{Id: strconv.Itoa(id), Name: token[2], Artist: token[3], Source: token[4], Type: token[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(token) == 3 {
			manager.RemoveByName(token[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", token[1])
	}
}

func handlePlayCommand(token []string) {
	if len(token) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	music := manager.Find(token[1])
	if music == nil {
		fmt.Println("The music ", token[1], " does not exist.")
		return
	}
	mp.Play(music.Name, music.Type)
}

func main() {
	fmt.Println("Enter following commands to control the player:\n",
		"lib list -- View the existing music lib\n",
		"lib remove <name> -- Remove the specified music from the lib\n",
		"play <name> -- Play the specified music\n",
		"lib add <name><artist><source><type> -- Add a music to the music lib")

	manager = lib.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command-> ")
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
