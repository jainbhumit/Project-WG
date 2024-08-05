package functionality

import (
	"bufio"
	"encoding/json"
	"file/models"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const todoFileName = "todos.json"

func ReadTodos(username string) (models.Todo, error) {
	todos := models.Todo{Username: username}
	file, err := os.Open(todoFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil // File doesn't exist, return empty list of todo
		}
		return todos, err
	}
	defer file.Close()

	var todoList []models.Todo
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todoList); err != nil {
		return todos, err
	}

	for _, todo := range todoList {
		if todo.Username == username {
			return todo, nil // Return the todo list for the current user
		}
	}

	return todos, nil // Return empty todo list if user not found
}

func WriteTodos(todo models.Todo) error {
	existingTodos, err := ReadAllTodos()
	if err != nil {
		return err
	}

	updated := false
	for i, t := range existingTodos {
		if t.Username == todo.Username {
			existingTodos[i] = todo
			updated = true
			break
		}
	}

	if !updated {
		existingTodos = append(existingTodos, todo)
	}

	data, err := json.MarshalIndent(existingTodos, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(todoFileName, data, 0644)
}

func ReadAllTodos() ([]models.Todo, error) {
	var todoList []models.Todo
	file, err := os.Open(todoFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return todoList, nil // File doesn't exist, return empty slice
		}
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&todoList); err != nil {
		return nil, err
	}

	return todoList, nil
}
func AddToDo(username string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the new task:")
	newTask, _ := reader.ReadString('\n')
	newTask = strings.TrimSpace(newTask) // Trim whitespace

	todo, err := ReadTodos(username)
	if err != nil {
		fmt.Println("Error reading to-do list:", err)
		return
	}

	todo.Tasks = append(todo.Tasks, newTask)

	err = WriteTodos(todo)
	if err != nil {
		fmt.Println("Error writing to-do list:", err)
		return
	}

	fmt.Println("Task added successfully.")
}

func DeleteToDo(username string) {
	taskId := make(map[int]string)

	var taskToDelete int
	fmt.Println("Enter the task to delete:")
	fmt.Scanln(&taskToDelete)

	todo, err := ReadTodos(username)
	if err != nil {
		fmt.Println("Error reading to-do list:", err)
		return
	}
	for i, task := range todo.Tasks {
		taskId[i+1] = task
	}
	updatedTasks := []string{}
	for _, task := range todo.Tasks {
		if task != taskId[taskToDelete] {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if len(updatedTasks) == len(todo.Tasks) {
		fmt.Println("Task not found.")
		return
	}

	todo.Tasks = updatedTasks

	err = WriteTodos(todo)
	if err != nil {
		fmt.Println("Error writing to-do list:", err)
		return
	}

	fmt.Println("Task deleted successfully.")
}
func ShowToDo(username string) {
	todo, err := ReadTodos(username)
	if err != nil {
		fmt.Println("Error reading to-do list:", err)
		return
	}

	if len(todo.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("To-Do List:")
	for i, task := range todo.Tasks {
		fmt.Printf("%d - %s\n", i+1, task)
	}
}
