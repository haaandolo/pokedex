package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	caught := fmt.Sprintf("%s escaped!", pokemon.Name)
	if pokemon.BaseExperience > 100 {
		caught = fmt.Sprintf("%s was caught!", pokemon.Name)
		cfg.pokedex[name] = pokemon
	}

	fmt.Printf("Throwing a Pokeball at %s... \n", pokemon.Name)
	fmt.Println(caught)

	return nil
}
