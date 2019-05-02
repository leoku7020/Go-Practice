package main

import "fmt"
//不定型別宣告
var anything
//一般宣告
var c, d int //同時宣告c d 為int
var e int, f string //宣告 e 為 int，f 為 string
var (
    x int  //宣告 int
    y int
    z int
    c bool //宣告 boolen
    python bool
    php bool
)
//宣告附初始值
var a, b int = 1, 2
var java, vue, react = true, false, "no!"

func main() {
    //短變數宣告，一定要用在函數內 可以省略var
    short_c, short_python, short_java := false, true, "yes!"

    fmt.Println("一般宣告")
    fmt.Println(x, y, z, c, python, php)
    fmt.Println("宣告初始值")
    fmt.Println(a, b, java, vue, react)
    fmt.Println("短變數宣告")
    fmt.Println(short_c, short_python, short_java)
}
