package main

import (
	"unsafe"
	"fmt"
)

func main()  {
	array := []int{1,2,3}
	// 取地址 物理的
	fmt.Println("sss",unsafe.Pointer(&array[0]))
	//将地址转换为10进制数
	base := uintptr(unsafe.Pointer(&array[0]))
	fmt.Println("base=",base)
	//获取 单位数据大小
	size := unsafe.Sizeof(array[0])
	fmt.Println("size=",size)
	//移动 指向
	ptr := unsafe.Pointer(base + 2 * size)
	fmt.Println("ptr=",ptr)
	element := *(*int)(ptr)
	fmt.Println("element=",element)
	fmt.Println(element,array[2])

	
}
