package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"taquion.com/interfaces/note"
	"taquion.com/interfaces/outputtable"
	"taquion.com/interfaces/saver"
	"taquion.com/interfaces/todo"
)

func main() {
	result := Add(1, 1) // metodo que usa generics
	fmt.Println("la suma es: ", result)
	doSomethingElse("Hola prro")
	note, err := note.New(
		getStringInput("Get note's Title"),
		getStringInput("Get note's Content"))
	if err != nil {
		fmt.Print(err)
	}

	doSomething(note)

	err = Display(note)
	if err != nil {
		fmt.Print(err)
	}

	todo, err := todo.New(
		getStringInput("Get todo's Content"))

	if err != nil {
		fmt.Println(err)
		return
	}

	err = Display(todo)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Add[T int | float64 | string](a, b T) T {
	return a + b
}

func doSomething(value interface{}) { //cualquier valor esta permitido
	switch value.(type) {
	case int:
		fmt.Println("es un numero")
	case string:
		fmt.Println("es un texto")
	case float64:
		fmt.Println("es un punto flotante")
	default:
		fmt.Println("no conozco este valor")
	}
}

func doSomethingElse(value interface{}) { //cualquier valor esta permitido
	intValue, ok := value.(int)
	if ok {
		fmt.Println("es un numero")
		fmt.Println("es un numero + 1:", intValue)
	}

	stringValue, ok := value.(string)
	if ok {
		fmt.Println("es un texto")
		fmt.Println("dice esto:", stringValue)
	}

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

func Display(data outputtable.Outputtable) error {
	data.Display()
	return Save(data)
}

func Save(data saver.Saver) error {
	return data.SaveToJson()
}
