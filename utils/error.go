package utils

import "fmt"

// IfError function to handle error
func IfError(err error) {

	// If error return error
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
}
