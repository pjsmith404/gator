package main

import (
	"database/sql"
	"github.com/pjsmith404/gator/internal/config"
	"github.com/pjsmith404/gator/internal/database"
	"log"
	"os"
)

import _ "github.com/lib/pq"

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	dbQueries := database.New(db)

	programState := state{
		db:     dbQueries,
		config: &cfg,
	}

	cmds := commands{
		commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatalf("No command provided")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatalf("Command failed: %v", err)
	}
}
