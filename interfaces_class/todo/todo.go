package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (Todo, error) {
	if isEmpty(content) {
		return Todo{}, errors.New("content can be empty")
	}

	return Todo{
		Text: content,
	}, nil
}

func (note Todo) SaveToJson() error {
	fmt.Println("saving...")
	fileName := "todo.json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func isEmpty(value string) bool {
	return value == ""
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}
