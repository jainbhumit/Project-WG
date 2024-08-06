package models

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
