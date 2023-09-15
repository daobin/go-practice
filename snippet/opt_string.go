package main

import (
	"fmt"
)

func main() {
	// STD 标准库字符串操作
	str := "abc\tadd hello y.o,u"
	fmt.Printf("\nstr: %#v\n\n", str)

	// 以空白符分隔字符串
	//fmt.Printf("strings.Fields: %#v\n\n", strings.Fields(str))

	// 指定分隔字符分隔原字符串，SplitAfter 会包含分隔符
	//fmt.Printf("strings.Split: %#v\n\n", strings.Split(str, "o"))
	//fmt.Printf("strings.SplitN: %#v\n\n", strings.SplitN(str, "o", 2))
	//fmt.Printf("strings.SplitAfter: %#v\n\n", strings.SplitAfter(str, "o"))

	// 指定切割符，返回前后两部分，如果没有找到切割符，返回的前部分为原字符串，后部分为空
	//before, after, found := strings.Cut(str, "o")
	//fmt.Printf("strings.Cut: %#v-%#v-%#v\n\n", before, after, found)

}
