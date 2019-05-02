package main

import "fmt"
//整數
var a int = 1;

//浮點數 相當於Ｃ的Double
//float32 float64 不能一起計算需轉型別
//float32
var floatValue float32 = 7.0
var floatValue2 float32 = 3.0
//float64 
var floatValue3 float64 = 4.0

//複數
var complexValue complex64 = 1.2 + 12i
var complexValue2 complex64 =1.2 + 10i
//complex128
var complexValue3 = complex(3.2, 12) 

//字串
var str string = "Hello"
var str2 = "World"

//布林 只能為true false
var boolean bool = true
var boolean2 bool = false

func main() {
    //Int
    fmt.Println("#Int")
    fmt.Println("a + 1 =", a + 1)
    //Float32
    fmt.Println("#Float")
    fmt.Println("7.0 / 3.0 =", floatValue/floatValue2)
    //Float32 & Float64
    fmt.Println("7.0 + 4.0 =", floatValue + float32(floatValue3))
    //Complex
    fmt.Println("#Complex")
    fmt.Println("complexValue3 + complexValue2 實數", real(complex64(complexValue3) + complexValue2))
    fmt.Println("complexValue3 + complexValue2 虛數", imag(complex64(complexValue3) + complexValue2))
    //String
    fmt.Println("#String")
    fmt.Println("1" + "1")
    fmt.Println(len(str + str2))
    fmt.Println(str, str2[1])
    fmt.Println(str +" "+ str2)
    fmt.Println("%c", str[1])
    //Boolean
    fmt.Println("#Boolean")
    fmt.Println("boolean=", boolean )
    fmt.Println("boolen2=", boolean2)
    fmt.Println("boolean || boolean2=", boolean || boolean2)
    fmt.Println("boolean && boolean2=", boolean && boolean2)
}
