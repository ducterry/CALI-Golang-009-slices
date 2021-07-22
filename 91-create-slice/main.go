package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {

	// 01. Khai báo 1 Slice
	var (
		mySlice = []int{3, 5, 7, 9, 11, 13, 17}
	)

	// 02. In màn hình
	mySlice = append(mySlice, 29)
	mySlice = append(mySlice, 02)
	mySlice = append(mySlice, 1991)
	fmt.Println("mySlice: ", mySlice)
}
