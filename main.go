package main

import (
	"fmt"
	"log"
	"os/user"
	"github.com/pjsmith404/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	user, err := user.Current()
	if err != nil {
		log.Fatalf("Couldn't get current user: %v", err)
	}

	err = cfg.SetUser(user.Username)
	if err != nil {
		log.Fatalf("Coudln't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	fmt.Println(cfg)
}
