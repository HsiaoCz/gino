package main

import "net/http"

// cookie 实际上就是服务器保存再浏览器上的一段信息
// 浏览器有了cookie之后，每次向服务器发送请求给服务器，服务器收到请求后，就可以根据该信息处理请求
// cookie运行原理
// 第一次向服务器发送请求时在服务端创建cookie
// 将在服务端创建的Cookie以响应头的方式发送给浏览器
// 下次再发送请求浏览器自动携带该cookie
// 服务器得到cookie之后根据cookie的信息区分不同的用户

// 创建cookie发送给客户端

func main() {
	http.HandleFunc("/ping", setCookie)
	http.HandleFunc("/getcookie", getCookie)
	http.HandleFunc("/expire-cookie", expireCookie)
	http.ListenAndServe(":9099", nil)
}

// 设置cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	// 1、创建cookie对象
	cookie := http.Cookie{
		Name:  "username",
		Value: "jason",
	}
	w.Header().Set("Set-Cookie", cookie.String())

	// 第二种方式设置cookie
	cookie2 := http.Cookie{
		Name:  "hello",
		Value: "12222",
	}
	w.Header().Add("Set-Cookie", cookie2.String())

	// 第三种设置cookie的方式
	http.SetCookie(w, &http.Cookie{Name: "lisi", Value: "1222"})
}

// 获取cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Get("Cookie")
	w.Write([]byte(cookie))
}

// 设置cookie的有效期
func expireCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "usernname",
		Value:  "jason",
		MaxAge: 60,
	}
	http.SetCookie(w, &cookie)
}

// session
// 第一次向服务器发送请求时创建session，给他设置一个全球唯一的ID
// 创建一个cookie，将cookie的value设置成session的ID值，并将cookie发送给浏览器
// 下次发送请求浏览器会携带(浏览器会处理)该cookie(其实就是session信息)
// 服务器获取cookie(其实就是session值)并根据他的value值找到服务器中对应的session，也就知道请求时哪个用户发送的
