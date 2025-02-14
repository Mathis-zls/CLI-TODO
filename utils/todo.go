package utils

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
	todo := Todo{
		Id:          todos[len(todos)-1].Id + 1,
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
		if index+1 != index {
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
