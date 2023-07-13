package routers

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/gino/quickdemo/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/student/create", controllers.CreateStudent).Methods("POST")
	r.HandleFunc("/student/update", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/delete", controllers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/student/getsignel", controllers.GetSingleStudent).Methods("GET")
	r.HandleFunc("/student/retrieve", controllers.RetrieveStudents).Methods("GET")

	srv := http.Server{
		Handler:      r,
		Addr:         ":9191",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	log.Fatal(srv.ListenAndServe())
}
