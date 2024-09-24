package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		fmt.Printf("you have not caught %s \n", name)
	}

	if ok {
		fmt.Printf("Name: %s \n", pokemon.Name)
		fmt.Printf("Height: %d \n", pokemon.Height)
		fmt.Printf("Weight: %d \n", pokemon.Weight)
	}

	return nil
}
