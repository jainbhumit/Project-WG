package models

//	type CourseProgress struct {
//		Course   string `json:"course_name"`
//		Progress int    `json:"progress"` //percentage of how much complete
//	}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//CoursesProgress []CourseProgress `json:"courses_progress"`
}
