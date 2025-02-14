package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var directory = "todocli"

func parseJsonToTodo() ([]Todo, error) {
	filename, err := loadConfig()
	if err != nil {
		return []Todo{}, err
	}
	filepath := fmt.Sprintf("%s/%s.json", directory, filename)
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
	filename, err := loadConfig()
	if err != nil {
		return err
	}
	filepath := fmt.Sprintf("%s/%s.json", directory, filename)
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

type Config struct {
	Filename string `json:"filname"`
}

func setConfig(name string) error {
	config := Config{
		Filename: name,
	}
	data, err := json.Marshal(config)
	if err != nil {
		return errors.New("marshal Json Config failed")
	}
	err = os.WriteFile(directory+"/config.json", data, 0644)
	if err != nil {
		return errors.New("write Json Config failed")
	}
	return nil
}

func loadConfig() (filename string, loadError error) {
	config := Config{}
	data, err := os.ReadFile(directory + "/config.json")
	if err != nil {
		return "", errors.New("failed to read Config file")
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return "", errors.New("failed to Unmarshal config file")
	}
	return config.Filename, nil
}
