package main

import "fmt"

//全局变量
var aa = 55
var ss = "def"

//常量,无需所有字母大写
const bb, cc = "bbbbb", true

//统一声明
var (
	aaa = 555
	sss = "global"
)

//变量默认都睡有初始值
func variableInit() {
	var a int
	var s string
	fmt.Println(a, s)
}

/**
多个变量同时赋值，
*/
func multiVariableInit() {
	var a, b, c int           //有初始值
	var d, e, f int = 1, 2, 3 //没有初始值
	fmt.Println(a, b, c, d, e, f)
}

/**
编译器自动识别类型
*/
func variableTypeDeduction() {
	var a, b, c, d = 1, 2, true, "def"
	fmt.Println(a, b, c, d)
}

/**
省略var
*/
func variableShorter() {
	a, b, c, d := 1, 2, true, "def"
	b = 5
	fmt.Println(a, b, c, d)
}
func main() {
	variableInit()
	multiVariableInit()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss)
	fmt.Println(aaa, sss)
	fmt.Println(bb, cc)
}
