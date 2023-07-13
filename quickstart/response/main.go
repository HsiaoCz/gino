package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 响应
// ResponseWriter 本身为一个interface
// Header()Header
// WriteHeader(int) 发送Http回复的头域和状态码
// Write([]byte)(int,error)
// write方法向连接中写入作为HTTP的一部分回复的数据

func main() {
	http.HandleFunc("/response", responseStatusCode)
	http.HandleFunc("/responseHTML", responseHTML)
	http.HandleFunc("/responseStringHTML", responseStringHTML)
	http.HandleFunc("/responseJSON", responseJSON)
	http.ListenAndServe(":9901", nil)
}

func responseStatusCode(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	// 返回的状态码可以是错误的
	// 数据可以是正确的
	w.WriteHeader(http.StatusOK)
}

// 返回的HTML的数据
// 如果先设置响应状态码，后面再设置响应头里面的数据就不生效了
func responseHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1 style='color:red'>Hello</h1>"))
}

// 返回以字符串显示的HTML
func responseStringHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1 style='color:red'> Hello</h1>"))
}

// 返回JSON格式的数据
type URLData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func responseJSON(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	username := r.FormValue("username")
	password := r.FormValue("password")
	urldata := &URLData{Username: username, Password: password}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(urldata)
}
