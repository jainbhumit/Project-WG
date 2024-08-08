package test

import (
	"bytes"
	"encoding/json"
	"file/functionality"
	"file/models"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// To capture the std output
// Function to capture output
func captureOutput(f func()) string {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	f()
	w.Close()
	os.Stdout = originalStdout
	return <-outC
}

// Helper function to create a test progress file
func createProgressFile(t *testing.T, progress []models.Progress) {
	file, err := os.Create("test_progress.json")
	if err != nil {
		t.Fatalf("Failed to create test progress file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(progress); err != nil {
		t.Fatalf("Failed to encode progress to test file: %v", err)
	}
}

func TestReadProgressValid(t *testing.T) {
	defer os.Remove("test_progress.json")
	createProgressFile(t, []models.Progress{
		{Username: "testuser", Courses: []models.CourseProgress{
			{CourseID: 1, CompletedLessons: []float32{1, 2, 3}, TotalLessons: 10},
		}},
	})

	originalFileName := functionality.Progressfile
	functionality.Progressfile = "test_progress.json"
	defer func() { functionality.Progressfile = originalFileName }()

	progress, err := functionality.ReadProgress()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(progress) == 0 {
		t.Fatalf("Expected 1 user progress, got %d", len(progress))
	}
}

func TestReadProgressFileNotFound(t *testing.T) {
	defer os.Remove("test_progress.json")

	originalFileName := functionality.Progressfile
	functionality.Progressfile = "test_progress.json"
	defer func() { functionality.Progressfile = originalFileName }()

	progress, err := functionality.ReadProgress()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(progress) != 0 {
		t.Fatalf("Expected 0 user progress, got %d", len(progress))
	}
}

func TestWriteProgressValid(t *testing.T) {
	defer os.Remove("test_progress.json")
	createProgressFile(t, []models.Progress{})

	originalFileName := functionality.Progressfile
	functionality.Progressfile = "test_progress.json"
	defer func() { functionality.Progressfile = originalFileName }()

	progress := models.Progress{
		Username: "testuser",
		Courses: []models.CourseProgress{
			{CourseID: 1, CompletedLessons: []float32{1, 2}, TotalLessons: 10},
		},
	}
	err := functionality.WriteProgress(progress)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	storedProgress, err := functionality.ReadProgress()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(storedProgress) == 0 {
		t.Fatalf("Expected 1 user progress, got %d", len(storedProgress))
	}
}

func TestWriteProgressInvalid(t *testing.T) {
	defer os.Remove("test_progress.json")
	createProgressFile(t, []models.Progress{})

	originalFileName := functionality.Progressfile
	functionality.Progressfile = "test_progress.json"
	defer func() { functionality.Progressfile = originalFileName }()

	progress := models.Progress{
		Username: "testuser",
		Courses: []models.CourseProgress{
			{CourseID: 1, CompletedLessons: []float32{1, 2}, TotalLessons: 10},
		},
	}
	err := functionality.WriteProgress(progress)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Introduce an error by creating a file with invalid JSON
	invalidJSON := []byte(`{"invalid": "json"}`)
	err = ioutil.WriteFile("test_progress.json", invalidJSON, 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid JSON to test file: %v", err)
	}

	storedProgress, err := functionality.ReadProgress()
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if len(storedProgress) != 0 {
		t.Fatalf("Expected 0 user progress, got %d", len(storedProgress))
	}
}

func TestShowUserProgressValid(t *testing.T) {
	defer os.Remove("test_progress.json")
	createProgressFile(t, []models.Progress{
		{Username: "testuser", Courses: []models.CourseProgress{
			{CourseID: 1, CompletedLessons: []float32{1, 2, 3}, TotalLessons: 10},
		}},
	})

	originalFileName := functionality.Progressfile
	functionality.Progressfile = "test_progress.json"
	defer func() { functionality.Progressfile = originalFileName }()

	// Capture the output
	output := captureOutput(func() {
		functionality.ShowUserProgress("testuser")
	})

	expectedOutput := `Progress for testuser:
Course ID: 1
Completed Lessons: [1 2 3]
Progress: 30.00%
`

	if output != expectedOutput {
		t.Fatalf("Expected output:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

func TestShowUserProgressUserNotFound(t *testing.T) {
	defer os.Remove("test_progress.json")
	createProgressFile(t, []models.Progress{})

	originalFileName := functionality.Progressfile
	functionality.Progressfile = "test_progress.json"
	defer func() { functionality.Progressfile = originalFileName }()

	// Capture the output
	output := captureOutput(func() {
		functionality.ShowUserProgress("nonexistentuser")
	})

	expectedOutput := "No progress found for user: nonexistentuser\n"

	if output != expectedOutput {
		t.Fatalf("Expected output:\n%s\nGot:\n%s", expectedOutput, output)
	}
}
