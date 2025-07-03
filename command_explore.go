package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("command 'explore' requires one location argument")
	}
	pokemonsResp, err := cfg.pokeapiClient.ExploreLocation(args[0])
	if err != nil {
		return err
	}

	for _, enc := range pokemonsResp.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}
	return nil
}
