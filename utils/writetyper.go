package utils

import (
	"fmt"
	"time"
)

func Echo(text string, dropline bool) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Print()
	if dropline {
		fmt.Println()
	}
}
