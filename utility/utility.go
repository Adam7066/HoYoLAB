package utility

import (
	"fmt"
)

// askBool 互動式詢問 y/n，回傳 bool
func AskBool(prompt string) bool {
	var response string
	for {
		fmt.Print(prompt)
		if _, err := fmt.Scanln(&response); err != nil {
			fmt.Println("Invalid input, please enter 'y' or 'n'.")
			continue
		}
		switch response {
		case "y", "Y":
			return true
		case "n", "N":
			return false
		default:
			fmt.Println("Invalid input, please enter 'y' or 'n'.")
		}
	}
}
