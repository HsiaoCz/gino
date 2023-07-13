package models

type Student struct {
	ID      int     `json:"id" form:"id"`
	Name    string  `json:"name" form:"name"`
	Class   string  `json:"class" form:"class"`
	Chinese float32 `json:"chinese" form:"chinese"`
	Math    float32 `json:"math" form:"math"`
	English float32 `json:"english" form:"english"`
}
