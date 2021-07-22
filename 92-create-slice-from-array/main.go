package main

import "fmt"

/*
	- Ngày: 22.07.2021
	- By: Nguyễn Đăng Đức
*/
func main() {
	// 01. Khai bao bien
	var (
		myArray = [5]string{"Alpha", "Beta", "Gamma", "Delta", "Epsilon"}
	)

	// 02. Create Slice
	var mySlice []string = myArray[1:4]

	//03. Print
	fmt.Println("Array a = ", myArray)
	fmt.Println("Slice s = ", mySlice)
}
