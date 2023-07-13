package main

import (
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
	http.ListenAndServe(":9901", nil)
}

func responseStatusCode(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.WriteHeader(http.StatusOK)
}
