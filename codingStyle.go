//強制性編譯風格規範
//左右括弧Error
/*
package main

import "fmt"

func main()
{
    i:=1
    fmt.Println("Hello World", i)
}
syntax error: unexpected semicolon or newline before {
non-declaration statement outside function body
syntax error: unexpected }
*/
//左右括弧Correct
/*
package main

import "fmt"

func main() {
    i:= 1
    fmt.Println("Hello World", i)
}

*/

//宣告變數卻未使用Error
/*
package main

import "fmt"

func main() {
    i:= 1
    fmt.Println("Hello World")
}
i declared and not used
*/

//非強制性編譯風格建議
package main
import "fmt"
func main(){i:=1
fmt.Println("Hello World",i)}
//可編譯只是很醜
//可用cmd go fmt *.go 自動格式化格式
//格式化做了以下
//調整每一條語句位置
//重新擺放括弧位置
//以tab幫你縮排
//添加空格
