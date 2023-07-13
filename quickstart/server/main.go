package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/user", userHandle)
	http.HandleFunc("/query", getQueryString)
	http.HandleFunc("/header", headerThing)
	http.HandleFunc("/body", requestBody)
	http.HandleFunc("/queryform", queryForm)
	http.ListenAndServe(":9090", nil)
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprintln(w, "hello")
}

// 获取URL的查询参数
func getQueryString(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery // 原生的，未经处理的
	fmt.Fprintf(w, "query string:%s", query)

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	fmt.Println(w, username)
	fmt.Fprintln(w, password)
}

// 获取请求头中的数据
func headerThing(w http.ResponseWriter, r *http.Request) {
	// header 是一个map[string][]string结构
	for key, value := range r.Header {
		fmt.Fprintf(w, "%s-%s", key, value)
	}
}

// 获取请求体中的数据
func requestBody(w http.ResponseWriter, r *http.Request) {
	// 方法一：获取body的数据
	length := r.ContentLength
	data := make([]byte, length)
	_, err := r.Body.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, string(data))
	// 这种方式获取到请求体的数据之后，还需要手动对数据进行分割

	// 方式二
	res, _ := io.ReadAll(r.Body)
	fmt.Fprintln(w, string(res))
	// 这种方式可能一次读到很大的数据量

	// 这两种方法都是手动处理
	// 还可以自动处理
	// 方式一:
	// 对原生的数据进行解析，解析到form里面
	// 然后通过r.Form.Get()获取到数据
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	fmt.Fprintln(w, username)
	fmt.Fprintln(w, password)

	// 方法二
	// 获取到body中的数据
	hello := r.FormValue("username")
	hi := r.FormValue("password")
	fmt.Fprintln(w, hello)
	fmt.Fprintln(w, hi)
}

// 获取请求参数
// form字段是http.url.Values类型，Form是解析好的表单数据，包括URL字段的query参数和Post或put的表单数据
// 本质为map[string][]string类型
// 这里通过form获取查询参数
// 如果是get请求，会将url里的参数解析到form里面
// 如果是post请求，会将body里面的数据也解析到form里面
func queryForm(w http.ResponseWriter, r *http.Request) {
	// 通过form获取查询参数和之前的一样
	// 需要先解析form
	r.ParseForm()
	hello := r.Form.Get("hello")
	hi := r.Form.Get("hi")
	fmt.Fprintln(w, hello)
	fmt.Fprintln(w, hi)
}
