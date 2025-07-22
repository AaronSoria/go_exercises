package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"createAt"`
}

func New(title, content string) (*Note, error) {
	if isEmpty(title) {
		return nil, errors.New("title can be empty")
	}
	if isEmpty(content) {
		return nil, errors.New("content can be empty")
	}

	return &Note{
		Title:    title,
		Content:  content,
		CreateAt: time.Now(),
	}, nil
}

func (note Note) SaveToJson() error {
	fmt.Println("saving...")
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func (note Note) Display() {
	fmt.Println(note.Title)
	fmt.Println(note.Content)
}

func isEmpty(value string) bool {
	return value == ""
}
