package main

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/Manoj-Reddy-Vadde/task/cmd"
	"github.com/Manoj-Reddy-Vadde/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main()  {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	// err := db.Init(dbPath)
	// if err != nil {
	// 	panic(err)
	// }
	must(cmd.RootCmd.Execute())
}

func must(err error)  {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}