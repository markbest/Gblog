package main

import (
	"blog/migrate"
	"fmt"
	"os"
)

var Usage = func() {
	fmt.Println("USAGE: migrate command [arguments] ...")
	fmt.Println("\nThe commands are:\n\taction\tmigrate action [create|up|down|status]")
	fmt.Println("\tfile\tmigrate file name")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	if args[1] == "help" || args[1] == "h" {
		Usage()
		return
	}

	switch args[1] {
	case "create":
		if len(args) != 3 {
			fmt.Println("USAGE: migrate create <filename>")
			return
		}
		migrate.CreateMigration(args[2])
	case "up":
		migrate.MigrateUp()
	case "down":
		migrate.MigrateDown()
	case "status":
		migrate.MigrateStatus()
	default:
		Usage()
	}
}
