package main

import (
	"fmt"
	"net/http"
	// "time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	
	// fmt.Printf("\nhost: %s\nurl: %s", r.Host, r.URL)
	// 读取表单前先执行
	r.ParseForm()

	// 遍历cookie
	for key, cookie := range r.Cookies() {
		fmt.Printf("%s:%s\n", key, cookie)
	}

	// 读取表单
	name := r.FormValue("name")

	//写头,Add, Del, get
	w.Header().Set("Server", "nginx") 
	http.SetCookie(w, &http.Cookie{
		Name:   "name",  //名字
		Value:  name,      //值
		Path:   "/",          //路径
		Domain: "localhost", //域名
		MaxAge: 120,          //存活时间
	})
	// 写状态码，参数为statuscode int
	w.WriteHeader(200)
	w.Write([]byte(name))
	// fmt.Fprintf(w, "Hello World! %s", name)
	
}

func main() {
	listen := "0.0.0.0:8080"
	fmt.Printf("server is runing at: %s\n", listen)
	http.HandleFunc("/", greet)
	http.ListenAndServe(listen, nil)
	
}