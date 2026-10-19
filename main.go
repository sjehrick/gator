package main

import (
	"fmt"
	"github.com/sjehrick/gator/internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading file: %w", err)
		os.Exit(1)
	}
	
	err = cfg.SetUser("sam")
	if err != nil {
		fmt.Printf("Error reading file: %w", err)
		os.Exit(1)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading file: %w", err)
		os.Exit(1)
	}

	fmt.Println(cfg)

}
