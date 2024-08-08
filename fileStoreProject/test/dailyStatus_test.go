package test

import (
	"encoding/json"
	"file/functionality"
	"file/models"
	"fmt"
	"os"
	"testing"
	"time"
)

// Helper function to create a test daily status file
func createDailyStatusFile(t *testing.T, statuses []struct {
	Username      string               `json:"username"`
	DailyStatuses []models.DailyStatus `json:"daily_statuses"`
}) {
	file, err := os.Create("daily_status_test.json")
	if err != nil {
		t.Fatalf("Failed to create test daily status file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(statuses); err != nil {
		t.Fatalf("Failed to encode daily statuses to test file: %v", err)
	}
}

func TestReadDailyStatusValid(t *testing.T) {
	defer os.Remove("daily_status_test.json")
	createDailyStatusFile(t, []struct {
		Username      string               `json:"username"`
		DailyStatuses []models.DailyStatus `json:"daily_statuses"`
	}{
		{
			Username: "testuser",
			DailyStatuses: []models.DailyStatus{
				{Date: time.Now().Truncate(24 * time.Hour), Status: "Test Status"},
			},
		},
	})
	originalFileName := functionality.DailyStatusFile
	functionality.DailyStatusFile = "daily_status_test.json"
	defer func() { functionality.DailyStatusFile = originalFileName }()

	statusMap, err := functionality.ReadDailyStatus()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(statusMap) == 0 {
		t.Fatalf("Expected 1 user status, got %d", len(statusMap))
	}
}

func TestReadDailyStatusFileNotFound(t *testing.T) {
	defer os.Remove("daily_status_test.json")

	originalFileName := functionality.DailyStatusFile
	functionality.DailyStatusFile = "daily_status_test.json"
	defer func() { functionality.DailyStatusFile = originalFileName }()

	statusMap, err := functionality.ReadDailyStatus()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(statusMap) != 0 {
		t.Fatalf("Expected 0 user status, got %d", len(statusMap))
	}
}

func TestWriteDailyStatusValid(t *testing.T) {
	defer os.Remove("daily_status_test.json")
	createDailyStatusFile(t, []struct {
		Username      string               `json:"username"`
		DailyStatuses []models.DailyStatus `json:"daily_statuses"`
	}{})

	originalFileName := functionality.DailyStatusFile
	functionality.DailyStatusFile = "daily_status_test.json"
	defer func() { functionality.DailyStatusFile = originalFileName }()

	statusMap := map[string][]models.DailyStatus{
		"testuser": {
			{Date: time.Now().Truncate(24 * time.Hour), Status: "Test Status"},
		},
	}
	err := functionality.WriteDailyStatus(statusMap)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	storedStatusMap, err := functionality.ReadDailyStatus()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(storedStatusMap) == 0 {
		t.Fatalf("Expected 1 user status, got %d", len(storedStatusMap))
	}
}

func TestWriteDailyStatusInvalid(t *testing.T) {
	defer os.Remove("daily_status_test.json")
	createDailyStatusFile(t, []struct {
		Username      string               `json:"username"`
		DailyStatuses []models.DailyStatus `json:"daily_statuses"`
	}{})

	originalFileName := functionality.DailyStatusFile
	functionality.DailyStatusFile = "daily_status_test.json"
	defer func() { functionality.DailyStatusFile = originalFileName }()

	statusMap := map[string][]models.DailyStatus{
		"testuser": {
			{Date: time.Now().Truncate(24 * time.Hour), Status: "Test Status"},
		},
	}
	err := functionality.WriteDailyStatus(statusMap)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Introduce an error by creating a file with invalid JSON
	invalidJSON := []byte(`{"invalid": "json"}`)
	err = os.WriteFile("daily_status_test.json", invalidJSON, 0644)
	if err != nil {
		t.Fatalf("Failed to write invalid JSON to test file: %v", err)
	}

	storedStatusMap, err := functionality.ReadDailyStatus()
	if err == nil {
		t.Fatalf("Expected error, got none")
	}
	if len(storedStatusMap) != 0 {
		t.Fatalf("Expected 0 user status, got %d", len(storedStatusMap))
	}

}

func TestShowDailyStatusValid(t *testing.T) {
	defer os.Remove("daily_status_test.json")
	createDailyStatusFile(t, []struct {
		Username      string               `json:"username"`
		DailyStatuses []models.DailyStatus `json:"daily_statuses"`
	}{
		{
			Username: "testuser",
			DailyStatuses: []models.DailyStatus{
				{Date: time.Now().Truncate(24 * time.Hour), Status: "Test Status"},
			},
		},
	})
	originalFileName := functionality.DailyStatusFile
	functionality.DailyStatusFile = "daily_status_test.json"
	defer func() { functionality.DailyStatusFile = originalFileName }()

	// Capture the output
	output := captureOutput(func() {
		functionality.ShowDailyStatus("testuser")
	})

	expectedOutput := fmt.Sprintf("Daily status for testuser:\nDate: %s\nStatus: Test Status\n", time.Now().Truncate(24*time.Hour).Format("2006-01-02"))

	if output != expectedOutput {
		t.Fatalf("Expected output:\n%s\nGot:\n%s", expectedOutput, output)
	}
}
