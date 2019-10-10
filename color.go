package hcolor

import (
	"fmt"
	"runtime"
)

type Color int

const (
	Black   Color = 30 + iota //黑
	Red                       //红
	Green                     //绿
	Yellow                    //黄
	Blue                      //蓝
	Magenta                   //洋红
	Cyan                      //蓝绿
	White                     //白
)

const (
	BBlack   Color = 40 + iota //黑背景
	BRed                       //红背景
	BGreen                     //绿背景
	BYellow                    //黄背景
	BBlue                      //蓝背景
	BMagenta                   //洋红背景
	BCyan                      //蓝绿背景
	BWhite                     //白背景
)

//格式化字体颜色
func FormatColor(msg string, color Color) string {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", color, msg)
	}
	return msg
}

//格式化背景颜色
func FormatBGColor(msg string, bgColor, bfColor Color) string {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		return fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m", bgColor, bfColor, msg)
	}
	return msg
}
