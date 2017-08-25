package main

import (
	"blog/migrate"
	"fmt"
	"flag"
)

var Usage = func() {
	fmt.Println("USAGE: migrate command [arguments] ...")
	fmt.Println("\nThe commands are:\n\taction\tmigrate action [create|up|down|status]")
	fmt.Println("\tfile\tmigrate file name")
}

func main() {
	flag.Parse()
	ch := flag.Args()

	if ch == nil || len(ch) == 0 {
		Usage()
		return
	}

	if ch[0] == "help" || ch[0] == "h" {
		Usage()
		return
	}

	switch ch[0] {
	case "create":
		if len(ch) != 2 {
			fmt.Println("USAGE: migrate create <filename>")
			return
		}
		migrate.CreateMigration(ch[1])
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
