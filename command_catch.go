package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("command 'catch' requires one pokemon argument")
	}
	catchResp, err := cfg.pokeapiClient.CatchPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("%s was caught!\n", catchResp.Name)
	return nil
}
