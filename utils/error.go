package utils

import "fmt"

func IfError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}
