package test

import (
	"encoding/json"
	"file/functionality"
	"file/models"
	"os"

	"testing"
)

// Helper function to create a test to-do file
func createToDoFile(t *testing.T, todos []models.Todo) {
	file, err := os.Create("todos_test.json")
	if err != nil {
		t.Fatalf("Failed to create test todo file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(todos); err != nil {
		t.Fatalf("Failed to encode todos to test file: %v", err)
	}
}

func TestReadTodos_Valid(t *testing.T) {
	defer os.Remove("todos_test.json")
	createToDoFile(t, []models.Todo{
		{
			Username: "testuser",
			Tasks:    []string{"Task 1", "Task 2"},
		},
	})
	originalFileName := functionality.TodoFileName
	functionality.TodoFileName = "todos_test.json"
	defer func() { functionality.TodoFileName = originalFileName }()

	todo, err := functionality.ReadTodos("testuser")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(todo.Tasks) != 2 {
		t.Fatalf("Expected 2 tasks, got %d", len(todo.Tasks))
	}
}

func TestReadTodos_UserNotFound(t *testing.T) {
	defer os.Remove("todos_test.json")
	createToDoFile(t, []models.Todo{})
	originalFileName := functionality.TodoFileName
	functionality.TodoFileName = "todos_test.json"
	defer func() { functionality.TodoFileName = originalFileName }()
	todo, err := functionality.ReadTodos("testuser")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(todo.Tasks) != 0 {
		t.Fatalf("Expected 0 tasks, got %d", len(todo.Tasks))
	}
}

func TestWriteTodos_Valid(t *testing.T) {
	defer os.Remove("todos_test.json")
	createToDoFile(t, []models.Todo{})

	todo := models.Todo{
		Username: "testuser",
		Tasks:    []string{"Task 1"},
	}
	originalFileName := functionality.TodoFileName
	functionality.TodoFileName = "todos_test.json"
	defer func() { functionality.TodoFileName = originalFileName }()
	err := functionality.WriteTodos(todo)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	storedTodos, err := functionality.ReadAllTodos()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(storedTodos) != 1 {
		t.Fatalf("Expected 1 todo, got %d", len(storedTodos))
	}
	if storedTodos[0].Username != "testuser" {
		t.Fatalf("Expected username 'testuser', got '%s'", storedTodos[0].Username)
	}
}
