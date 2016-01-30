package hcolor

import (
	"testing"
)

func TestColor(t *testing.T) {
	t.Log(FormatColor("测试文本:黑", Black))
	t.Log(FormatColor("测试文本:红", Red))
	t.Log(FormatColor("测试文本:绿", Green))
	t.Log(FormatColor("测试文本:黄", Yellow))
	t.Log(FormatColor("测试文本:蓝", Blue))
	t.Log(FormatColor("测试文本:洋红", Magenta))
	t.Log(FormatColor("测试文本:蓝绿", Cyan))
	t.Log(FormatColor("测试文本:白", White))

}

func TestBGColor(t *testing.T) {
	t.Log(FormatBGColor("测试文本:黑背景", BBlack, White))
	t.Log(FormatBGColor("测试文本:红背景", BRed, Cyan))
	t.Log(FormatBGColor("测试文本:绿背景", BGreen, Magenta))
	t.Log(FormatBGColor("测试文本:黄背景", BYellow, Blue))
	t.Log(FormatBGColor("测试文本:蓝背景", BBlue, Yellow))
	t.Log(FormatBGColor("测试文本:洋红背景", BMagenta, Green))
	t.Log(FormatBGColor("测试文本:蓝绿背景", BCyan, Red))
	t.Log(FormatBGColor("测试文本:白背景", BWhite, Black))
}
