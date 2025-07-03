package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(pokemon string) (RespPokemonDetail, error) {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	url := baseURL + "/pokemon/" + pokemon + "/"

	if val, ok := c.cache.Get(url); ok {
		cachedExploreResp := RespPokemonDetail{}
		json.Unmarshal(val, &cachedExploreResp)
		return cachedExploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonDetail{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemonDetail{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonDetail{}, err
	}

	pokemonResp := RespPokemonDetail{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemonDetail{}, err
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}
