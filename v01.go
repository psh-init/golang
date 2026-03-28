package main

import "fmt"

func userInfo() {
	userId := 1001
	userName := "Priyanshu"
	isActive := false
	targetFound := true
	totalAsset := 12500000000
	fmt.Printf(
		"UserID: %d\nUserName: %s\nIsActive: %t\nTargetFound: %t\nTotalAsset: %d\n",
		userId,
		userName,
		isActive,
		targetFound,
		totalAsset,
	)
}

func main() {
	userInfo()
}
