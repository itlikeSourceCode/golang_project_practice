// 包的申明，关键字package
// 申明包的格式： package 包的名称
// 包的名称：main
// main包是执行程序的入口
package main

// 包的引入，关键字import
// 引入的包的格式： import "包的名称"
// 包的名称fmt
import "fmt"

// 全局常量的定义，关键字const
// 定义全局常量的格式：const 常量名 变量类型 = 值
// 变量名：字母，数字，下划线(只能是字母或下划线开头)
const max int = 10

// 全局变量的定义，关键字var
var a int = 1

// var a = 1
// a := 1

// 函数，关键字func
// 函数定义的格式：func 函数名(参数) 返回值
// main 主函数的入口
func main() {

	// 内部变量
	b := 2

	// 选择语句，关键字：if
	// 表达式：a+b > max
	if a+b > max {
		// 调用外部包函数的方式： 包的名称.函数的名称
		// 规范：首字母大写，才能在外部包访问
		fmt.Printf("a + b > MAX\n")
	} else {
		fmt.Printf("a + b <= MAX\n")
	}
}
