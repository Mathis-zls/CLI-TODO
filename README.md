
CLI-TODO is a Command Line Interface for creating Todolists.

# Overview

CLI-TODO provides:
* Adding Todos
* Updating Todos
* Deleting Todos
* Creating new Todolists
* Loading Todolists
* Viewing Todolists
* Loading Todolists


## Commands

Following Commands are available:
Shows all commands that are avalible:
```bash
./CLI-TODO 
```
Creating a new Todolist:
you can use a name only once
```bash
./CLI-TODO NewTodoList -n "exampleName" 
```
Loading a new Todolist:
-n/-name is required
```bash
./CLI-TODO LoadTodoList -n "exampleName" 
```
Adding a new Todo:
Only -n/-name is required and an existing Todolist
```bash
./CLI-TODO AddTodo -n "exampleName" -d "exampleDiscription" -t "exampleDeadline"
```
Deleting a Todo:
Only -i/-id of the Todo is required.
```bash
./CLI-TODO DeleteTodo -i 1 
```
Updating a Todo:
Only -i/-id is required and at least one other parameter
```bash
./CLI-TODO UpdateTodo -id 1 -n "exampleName" -d "exampleDiscription" -t "exampleDeadline"
```
Complete a Todo:
Only -i/-id is required
```bash
./CLI-TODO MarkTodoAsComplete -id 1
```
Viewing the Todos:
```bash
./CLI-TODO ViewTodos
```

# Installing
Clone the Repository or Download CLI-TODO / CLI-TODO.exe

# Usage
First you need to be in the directory of the CLI-TODO / CLI-TODO.exe file.
Then you can run all the Commands.

#Info 
This project is build with:
cobra [https://github.com/spf13/cobra] 
simpletable[https://github.com/alexeyco/simpletable]
