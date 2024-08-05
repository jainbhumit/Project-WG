package util

import (
	"encoding/json"
	"file/models"
	"io"
	"log"
	"os"
)

func LoadCourses(obj *models.Course) {
	// Open the JSON file

	f, err := os.Open("course.json")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer f.Close()

	// Read the file content
	byteValue, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	// Unmarshal the JSON content into courseObj
	err = json.Unmarshal(byteValue, &obj)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}
}
