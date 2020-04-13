// main 可执行程序的入口
package main

// 引入包
import (
	"fmt"
	"net/http"      // HTTP网络库
	"text/template" // 页面渲染
)

// 结构体
// type 结构体名称 struct
type Handler struct {
}

// 定义一个结构体的方法
// 首页的Handler
func (h *Handler) Index(writer http.ResponseWriter, request *http.Request) {
	var data struct {
		Title string // 成员变量
		Body  string
	}

	data.Title = "Go Web"
	data.Body = "Welcome!"

	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintln(writer, err)
		return
	}

	t.Execute(writer, data)
}

// 登录的Handler
func (h *Handler) Login(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintln(writer, err)
		return
	}

	username := request.FormValue("username")
	password := request.FormValue("password")

	// 类型转换
	if username == "admin" && password == "admin" {
		writer.Write([]byte("登录成功"))
	} else {
		writer.Write([]byte("用户名或密码错误"))
	}
}

func main() {
	handler := &Handler{} // 实例化一个结构体类型
	// handler = new(Handler)

	// 定义路由规则
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/login", handler.Login)

	// 启动HTTP服务
	http.ListenAndServe(":8888", nil)
}
