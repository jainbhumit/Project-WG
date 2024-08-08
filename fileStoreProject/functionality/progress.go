package functionality

import (
	"encoding/json"
	"file/models"
	"fmt"
	"io/ioutil"
	"os"
)

var Progressfile = "progress.json"

func ReadProgress() (map[string]models.Progress, error) {
	progress := make(map[string]models.Progress)
	file, err := os.Open(Progressfile)
	if err != nil {
		if os.IsNotExist(err) {
			return progress, nil // File doesn't exist, return empty map
		}
		return nil, err
	}
	defer file.Close()

	var progressList []models.Progress
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&progressList); err != nil {
		return nil, err
	}

	for _, prog := range progressList {
		progress[prog.Username] = prog
	}

	return progress, nil
}
func WriteProgress(progress models.Progress) error {
	progressList, err := ReadProgress()
	if err != nil {
		return err
	}

	progressList[progress.Username] = progress

	var progressArray []models.Progress
	for _, p := range progressList {
		progressArray = append(progressArray, p)
	}

	data, err := json.MarshalIndent(progressArray, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(Progressfile, data, 0644)
}
func ShowUserProgress(username string) {
	progressList, err := ReadProgress()
	if err != nil {
		fmt.Println("Error reading progress:", err)
		return
	}

	progress, ok := progressList[username]
	if !ok {
		fmt.Println("No progress found for user:", username)
		return
	}

	fmt.Printf("Progress for %s:\n", username)
	completed, total := func() (int, int) {
		var totalCompleted int
		var totalLessons int
		for _, course := range progress.Courses {
			if totalLessons == 0 {
				totalLessons = course.TotalLessons
			}
			totalCompleted += len(course.CompletedLessons)
		}
		return totalCompleted, totalLessons
	}()
	percentage := float64(completed) / float64(total) * 100
	for _, course := range progress.Courses {
		fmt.Printf("Course ID: %d\n", course.CourseID)
		fmt.Printf("Completed Lessons: %v\n", course.CompletedLessons)
	}
	fmt.Printf("Progress: %.2f%%\n", percentage)

}
func UpdateUserProgress(username string) error {
	var courseID int
	var lessonID float32
	fmt.Println("Enter course ID: ")
	fmt.Scan(&courseID)
	fmt.Println("Enter lesson ID: ")
	fmt.Scan(&lessonID)

	progressList, err := ReadProgress()
	if err != nil {
		return err
	}

	progress, exists := progressList[username]
	if !exists {
		// Initialize progress if it doesn't exist
		progress = models.Progress{
			Username: username,
			Courses:  []models.CourseProgress{},
		}
		progressList[username] = progress
	}

	// Find the course in the user's progress
	courseFound := false
	for i, course := range progress.Courses {
		if course.CourseID == courseID {
			// Add the completed lesson if it doesn't already exist
			for _, completedLesson := range course.CompletedLessons {
				if completedLesson == lessonID {
					// Lesson already completed
					return nil
				}
			}
			progress.Courses[i].CompletedLessons = append(progress.Courses[i].CompletedLessons, lessonID)
			courseFound = true
			break
		}
	}

	// If the course was not found in the user's progress, add it
	if !courseFound {
		newCourse := models.CourseProgress{
			CourseID:         courseID,
			CompletedLessons: []float32{lessonID},
			TotalLessons:     12,
		}
		progress.Courses = append(progress.Courses, newCourse)
	}

	// Write the updated progress back to the file
	return WriteProgress(progress)
}
