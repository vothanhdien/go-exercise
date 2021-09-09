package main

import (
	"baitapbuoi2/findMaxLengthElement"
	"baitapbuoi2/max2Numbers"
	"baitapbuoi2/removeDuplicates"
	"baitapbuoi2/staff"
	"fmt"
)

func main() {
	smax2Numbers := []int{1, 9, 5, 6, 6}
	sfindMaxLengthElement := []string{"a", "avc", "aadf", "axxx"}
	sremoveDuplicates := []int{1, 1, 2, 3, 6, 6, 7, 9}
	staffDic := []staff.Staff{
		{"Bli", 1.2, 100000},
		{"Cli", 1.3, 500000},
		{"Zli", 1.3, 500000},
		{"Ali", 1.4, 600000},
		{"Xli", 1.4, 600000},
		{"Dli", 1.5, 700000},
	}
	if result, err := max2Numbers.Get(smax2Numbers); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bai1 max2Numbers: ", result)
	}
	if result, err := findMaxLengthElement.Find(sfindMaxLengthElement); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bai2 findMaxLengthElement:", result)
	}
	if result, err := removeDuplicates.Remove(sremoveDuplicates); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bai3 removeDuplicates:", result)
	}
	if result, err := staff.SortName(staffDic); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bai4 SortName Ascending:", result)
	}
	if result, err := staff.Salary(staffDic); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bai4 Salary decrease:", result)
	}
	if result, err := staff.Max2Salary(staffDic); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Bai4 Max2Salary:", result)
	}
}
