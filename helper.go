package main

import (
	"fmt"
	"os"
)

// This is only used for quick dev extract external api response
func writeToLogFile(filePath, log string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(log)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
