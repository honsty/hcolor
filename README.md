# Linux 终端文字及背景颜色 v1.2

```go
package main
import(
        "fmt"
        "github.com/honsty/hcolor"
)

func main(){
        fmt.Println(hcolor.FormatColor("测试文本:黑",hcolor.Black))
        fmt.Println(hcolor.FormatColor("测试文本:红",hcolor.Red))
        fmt.Println(hcolor.FormatColor("测试文本:绿",hcolor.Green))
        fmt.Println(hcolor.FormatColor("测试文本:黄",hcolor.Yellow))
        fmt.Println(hcolor.FormatColor("测试文本:蓝",hcolor.Blue))
        fmt.Println(hcolor.FormatColor("测试文本:洋红",hcolor.Magenta))
        fmt.Println(hcolor.FormatColor("测试文本:蓝绿",hcolor.Cyan))
        fmt.Println(hcolor.FormatColor("测试文本:白",hcolor.White))

        fmt.Println(hcolor.FormatBGColor("测试文本:黑背景",hcolor.BBlack,hcolor.White))
        fmt.Println(hcolor.FormatBGColor("测试文本:红背景",hcolor.BRed,hcolor.Cyan))
        fmt.Println(hcolor.FormatBGColor("测试文本:绿背景",hcolor.BGreen,hcolor.Magenta))
        fmt.Println(hcolor.FormatBGColor("测试文本:黄背景",hcolor.BYellow,hcolor.Blue))
        fmt.Println(hcolor.FormatBGColor("测试文本:蓝背景",hcolor.BBlue,hcolor.Yellow))
        fmt.Println(hcolor.FormatBGColor("测试文本:洋红背景",hcolor.BMagenta,hcolor.Green))
        fmt.Println(hcolor.FormatBGColor("测试文本:蓝绿背景",hcolor.BCyan,hcolor.Red))
        fmt.Println(hcolor.FormatBGColor("测试文本:白背景",hcolor.BWhite,hcolor.Black))

}
```
