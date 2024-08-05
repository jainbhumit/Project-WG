package models

import "time"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Mobile   string `json:"mobile"`
}

type Course struct {
	Courses []struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Lessons []struct {
			ID    float32 `json:"id"`
			Title string  `json:"title"`
		} `json:"lessons"`
	} `json:"courses"`
}

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
}
type DailyStatus struct {
	Date   time.Time `json:"date"`
	Status string    `json:"status"`
}
