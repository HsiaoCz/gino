# gino

## 1、web 工作原理

C/S 架构

client-server

client 发送 request
server 返回 response

server 端有一个多路复用器，负责路由请求

go http

```go
func main() {
	http.HandleFunc("/user", userHandle)
	http.ListenAndServe(":9090", nil)
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

```
