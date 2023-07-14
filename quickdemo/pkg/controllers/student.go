package controllers

import (
	"net/http"
	"strconv"

	"github.com/HsiaoCz/gino/quickdemo/pkg/models"
	"github.com/HsiaoCz/gino/quickdemo/pkg/utils"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("请输入数字学号"))
		return
	}
	chinese, err := strconv.ParseFloat(r.FormValue("chinese"), 32)
	if err != nil {
		utils.ResponseString(w, http.StatusOK, err.Error())
		return
	}
	math, err := strconv.ParseFloat(r.FormValue("math"), 32)
	if err != nil {
		utils.ResponseString(w, http.StatusOK, err.Error())
		return
	}
	english, err := strconv.ParseFloat(r.FormValue("english"), 32)
	if err != nil {
		utils.ResponseString(w, http.StatusOK, err.Error())
	}
	student := &models.Student{
		ID:      id,
		Name:    r.FormValue("name"),
		Class:   r.FormValue("class"),
		Chinese: float32(chinese),
		Math:    float32(math),
		English: float32(english),
	}
	utils.ResponseJSON(w, http.StatusOK, student)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request)    {}
func UpdateStudent(w http.ResponseWriter, r *http.Request)    {}
func GetSingleStudent(w http.ResponseWriter, r *http.Request) {}
func RetrieveStudents(w http.ResponseWriter, r *http.Request) {}
