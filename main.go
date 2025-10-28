package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"hoyolab/checkin"
	"hoyolab/db"
	"hoyolab/utility"
)

func main() {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		os.Exit(1)
	}
	dbPath := filepath.Join(filepath.Dir(execPath), "hoyolab.sqlite3")

	dbInstance, err := db.InitDB(dbPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open DB: %v\n", err)
		os.Exit(1)
	}

	addUser := flag.Bool("add", false, "add a new user to the database")
	flag.Parse()
	if *addUser {
		utility.AddUserInteractive(dbInstance)
		return
	}

	users, err := dbInstance.GetAllUsers()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get users: %v\n", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for _, u := range users {
		if u.LastCheckinAt != nil {
			now := time.Now().In(time.Local)
			if u.LastCheckinAt.In(time.Local).Year() == now.Year() &&
				u.LastCheckinAt.In(time.Local).YearDay() == now.YearDay() {
				continue
			}
		}
		wg.Add(1)
		go func(user db.User) {
			defer wg.Done()
			checkin.CheckinUser(user)
			dbInstance.UpdateLastCheckinAt(user.ID, time.Now())
		}(u)
	}
	wg.Wait()
}
