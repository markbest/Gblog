package main

import (
	"flag"
	"fmt"
	"github.com/markbest/Gblog/migrate"
)

var Usage = func() {
	fmt.Println("USAGE: migrate command [arguments] ...")
	fmt.Println("\nThe commands are:\n\taction\tmigrate [create|up|down|status]")
	fmt.Println("\tfile\tmigrate create file")
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
		migrate.HandleMigrateUp()
	case "down":
		migrate.HandleMigrateDown()
	case "status":
		migrate.HandleMigrateStatus()
	default:
		Usage()
	}
}