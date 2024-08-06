package models

import "time"

type Todo struct {
	Username string   `json:"username"`
	Tasks    []string `json:"tasks"`
}

type Progress struct {
	Username string           `json:"username"`
	Courses  []CourseProgress `json:"courses"`
}

type CourseProgress struct {
	CourseID         int       `json:"course_id"`
	CompletedLessons []float32 `json:"completed_lessons"`
	TotalLessons     int       `json:"total_lessons"`
}
type DailyStatus struct {
	Date   time.Time `json:"date"`
	Status string    `json:"status"`
}
