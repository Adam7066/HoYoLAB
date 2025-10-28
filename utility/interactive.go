package utility

import (
	"fmt"
	"os"

	"hoyolab/db"
)

// AddUserInteractive 互動式詢問所有欄位並寫入資料庫
func AddUserInteractive(dbInstance *db.DB) {
	var user db.User
	for {
		fmt.Print("Name: ")
		if _, err := fmt.Scanln(&user.Name); err == nil && user.Name != "" {
			break
		} else {
			fmt.Println("Invalid input, please enter a valid name.")
		}
	}
	for {
		fmt.Print("LtmidV2: ")
		if _, err := fmt.Scanln(&user.LtmidV2); err == nil && user.LtmidV2 != "" {
			break
		} else {
			fmt.Println("Invalid input, please enter a valid LtmidV2.")
		}
	}
	for {
		fmt.Print("LtuidV2: ")
		if _, err := fmt.Scanln(&user.LtuidV2); err == nil && user.LtuidV2 != "" {
			break
		} else {
			fmt.Println("Invalid input, please enter a valid LtuidV2.")
		}
	}
	for {
		fmt.Print("LtokenV2: ")
		if _, err := fmt.Scanln(&user.LtokenV2); err == nil && user.LtokenV2 != "" {
			break
		} else {
			fmt.Println("Invalid input, please enter a valid LtokenV2.")
		}
	}

	user.Genshin = AskBool("Genshin (y/n): ")
	user.Starrail = AskBool("Starrail (y/n): ")
	user.ZZZ = AskBool("ZZZ (y/n): ")

	err := dbInstance.InsertUser(&user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to add user: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("User added successfully.")
}
