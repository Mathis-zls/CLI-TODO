package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var filename = ""
var directory = "todocli"

func parseJsonToTodo() ([]Todo, error) {
	filepath := fmt.Sprintf("%s/%s", directory, filename)

	tasks := []Todo{}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return tasks, errors.New("cant read the json file prob filename wrong")
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return []Todo{}, errors.New("coudn't parse Json file")
	}
	return tasks, nil
}

func parseTodoToJson(todos []Todo) error {
	filepath := fmt.Sprintf("%s/%s", directory, filename)

	data, err := json.Marshal(todos)
	if err != nil {
		return errors.New("coudn't Marshal json")
	}
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		return errors.New("coudn't write to the file")
	}
	return nil
}

func loadDir() ([]os.DirEntry, error) {
	if _, err := os.Stat(directory); err != nil {
		err := os.Mkdir(directory, 0755)
		if err != nil {
			return nil, errors.New("cant create Directory")
		}
	}
	filenames, err := os.ReadDir(directory)
	if err != nil {
		return nil, errors.New("cant read Directory")
	}
	return filenames, nil
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

	filename = name
	return nil
}

func LoadSave(name string) error {
	names, err := getValidFileNames()
	if err != nil {
		return err
	}
	found := false
	for _, file := range names {
		if file == name {
			found = true
			break
		}
	}
	if !found {
		return errors.New("no File with that name saved")
	}
	return nil

}

func getValidFileNames() ([]string, error) {
	filenames, err := loadDir()
	if err != nil {
		return nil, err
	}
	validFileNames := []string{}
	for _, file := range filenames {
		validFileNames = append(validFileNames, file.Name())
	}
	return validFileNames, nil
}
