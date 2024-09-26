package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("your pokedex is empty")
	}

	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.pokedex {
		fmt.Println("- %s", key)
	}
	return nil
}
