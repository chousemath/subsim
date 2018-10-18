package utils

import (
	"fmt"
	"log"
)

func errStr(err error) string {
	return fmt.Sprintf("Error: %v\n", err)
}

// CheckErr xxx
func CheckErr(err error) {
	if err != nil {
		log.Fatal(errStr(err))
	}
}

// LogErr xxx
func LogErr(err error) {
	if err != nil {
		fmt.Println(errStr(err))
	}
}
