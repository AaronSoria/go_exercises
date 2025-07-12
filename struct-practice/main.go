package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

func main() {
	note, err := note.New(
		getStringInput("Get note's Title"),
		getStringInput("Get note's Content"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = note.SaveToJson()
	if err != nil {
		fmt.Print(err)
	}
	println("Done. :D")
}

func getStringInput(message string) (input string) {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	return input
}
