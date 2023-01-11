package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s string = "hello"
	fmt.Println("original string:", s)
	modifyString(&s)
	fmt.Println(s)

}

func modifyString(s *string) {
	p := (*uintptr)(unsafe.Pointer(s))
	var arr *[5]byte = (*[5]byte)(unsafe.Pointer(*p))
	var len *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Sizeof((*uintptr)(nil))))

	for i := 0; i < (*len); i++ {
		fmt.Printf("%p => %c\n", &((*arr)[i]), (*arr)[i])
		p1 := &((*arr)[i])
		v := (*p1)
		(*p1) = v + 1
	}
}
