package utils

import (
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Done        bool
	Id          int
	Name        string
	Description string
	Due_time    string
}

func GetTodos() ([]Todo, error) {
	return parseJsonToTodo()
}

type TodoParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Due_time    string `json:"due_too"`
}

func AddTodo(arg TodoParams) error {
	todos, err := parseJsonToTodo()
	if err != nil {
		return err
	}
	newID := len(todos) + 1
	todo := Todo{
		Id:          newID,
		Name:        arg.Name,
		Description: arg.Description,
		Due_time:    arg.Due_time,
		Done:        false,
	}
	todos = append(todos, todo)
	err = parseTodoToJson(todos)
	if err != nil {
		return err
	}
	return nil

}
func DeleteTodo(id int) error {
	todos, err := parseJsonToTodo()
	if err != nil {
		return err
	}
	newTodos := []Todo{}
	for index, todo := range todos {
		if index+1 != id {
			if todo.Id > id {
				todo.Id--
			}
			newTodos = append(newTodos, todo)
		}
	}
	err = parseTodoToJson(newTodos)
	if err != nil {
		return err
	}
	return nil
}

func GetMaxID() (int, error) {
	todos, err := parseJsonToTodo()
	if err != nil {
		return 0, err
	}
	return todos[len(todos)-1].Id, nil
}

func UpdateTodo(id int, arg TodoParams) error {
	todos, err := parseJsonToTodo()
	if err != nil {
		return err
	}
	var todo *Todo = &todos[id-1]
	if arg.Description != "" {
		todo.Description = arg.Description
	}
	if arg.Name != "" {
		todo.Name = arg.Name
	}
	if arg.Due_time != "" {
		todo.Due_time = arg.Due_time
	}
	err = parseTodoToJson(todos)
	if err != nil {
		return err
	}
	return nil

}

func MarkTodoAsComplete(id int) error {
	todos, err := parseJsonToTodo()
	if err != nil {
		return err
	}
	todos[id-1].Done = true
	err = parseTodoToJson(todos)
	if err != nil {
		return err
	}
	return nil
}
func NewSave(name string) error {
	filenames, err := loadDir()
	if err != nil {
		return err
	}
	for _, file := range filenames {
		if name == file.Name() {
			return errors.New("file already exists")
		}
	}
	err = setConfig(name)
	if err != nil {
		return err
	}

	err = parseTodoToJson([]Todo{})
	if err != nil {
		return errors.New("coudnt create File")
	}
	return nil
}

func LoadSave(name string) error {
	err := isNameValid(name)
	if err != nil {
		return err
	}

	err = setConfig(name)
	if err != nil {
		return err
	}
	return nil

}

func DeleteSave(name string) error {
	loadedList, err := loadConfig()
	if err != nil {
		return err
	}
	if loadedList == name {
		if err = setConfig(""); err != nil {
			return err
		}
	}
	err = isNameValid(name)
	if err != nil {
		return err
	}
	err = os.Remove(fmt.Sprintf("%s/%s.json", directory, name))
	if err != nil {
		return errors.New("can't delete File")
	}

	return nil
}
