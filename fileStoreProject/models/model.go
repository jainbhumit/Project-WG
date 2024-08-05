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

//type Course struct {
//	Courses []struct {
//		ID      string   `json:"id"`
//		Title   string   `json:"title"`
//		Lessons []string `json:"lessons"`
//	} `json:"courses"`
//}
