package main

import "fmt"

//接口本身是 调用方 和 实现方 均 需要遵守的一种协议， 大家按照统一的方法 命名 参数类型和数量 来协调 逻辑处理的 功能
//
//go语言中  使用 组合 实现 对象特性描述， 对象的内部 使用 结构体内嵌 组合 对象 应该具有的特性，对外通过 接口 暴露 能使用的特性
//
//go语言的每个接口中的方法数量不会很多，但是通过多个接口的嵌入和组合的方式，可以将简单的接口扩展为复杂的接口
//
// 1、接口的方法与实现接口的类型方法格式一致
// 2、接口中所以方法都被实现

//几种常见的接口无法实现的错误
//	1、函数名不一致，导致报错
//	2、实现接口的方法签名不一致导致的报错

//go语言 类型 断言 简述
// go语言中有四种 接口相关的类型 转换情形

//i.T
func main()  {
	// 编译器将123的类型推断为  内置类型int
	var x interface{} = 123
	// 断言为非接口类型
	//
	n, ok := x.(int)
	fmt.Println(n, ok)
	n = x.(int)
	fmt.Println(n)

	//
	a, ok := x.(float32)
	fmt.Println(a, ok)
	//
	b := x.(float32)
	fmt.Println(b)

}




