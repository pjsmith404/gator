package main

import (
	"fmt"
	"os/user"
	"github.com/pjsmith404/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}

	err = cfg.SetUser(user.Username)
	if err != nil {
		fmt.Println(err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cfg)
}
