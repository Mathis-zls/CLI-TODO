package utils

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

var filename = ""

type Todo struct {
	Done        bool
	Id          int
	Name        string
	Description string
	Due_time    time.Time
}

func GetTodos(filename string) ([]Todo, error) {
	return parseJsonToTodo(filename)
}

func parseJsonToTodo(file string) ([]Todo, error) {
	tasks := []Todo{}
	data, err := os.ReadFile(file)
	if err != nil {
		return tasks, errors.New("Cant read the json file prob filename wrong")
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return []Todo{}, errors.New("Coudn't parse Json file")
	}
	return tasks, nil
}

type TodoParams struct {
	Done        bool      `json:"done"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Due_time    time.Time `json:"due_too"`
}

func AddTodo(arg TodoParams) error {
	todos, err := parseJsonToTodo(filename)
	todo := Todo{
		Id:          todos[len(todos)-1].Id + 1,
		Name:        arg.Name,
		Description: arg.Description,
		Due_time:    arg.Due_time,
	}
	todos = append(todos, todo)
	parseTodoToJson(todos)
}

func parseTodoToJson(todos []Todo) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return errors.New("Coudn't Marshal json")
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return errors.New("Coudn't write to the file")
	}
	return nil
}

func DeleteTodo(id int) {

}
