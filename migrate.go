package main

import (
	"blog/migrate"
	"flag"
	"fmt"
)

var Usage = func() {
	fmt.Println("USAGE: migrate command [arguments] ...")
	fmt.Println("\nThe commands are:\n\taction\tmigrate action [create|up|down|status]")
	fmt.Println("\tfile\tmigrate file name")
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		Usage()
		return
	}

	if args[0] == "help" || args[0] == "h" {
		Usage()
		return
	}

	switch args[0] {
	case "create":
		if len(args) != 2 {
			fmt.Println("USAGE: migrate create <filename>")
			return
		}
		migrate.CreateMigration(args[1])
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
