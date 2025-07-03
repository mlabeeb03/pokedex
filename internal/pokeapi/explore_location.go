package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location string) (RespLocationDetail, error) {
	url := baseURL + "/location-area/" + location + "/"

	if val, ok := c.cache.Get(url); ok {
		cachedExploreResp := RespLocationDetail{}
		json.Unmarshal(val, &cachedExploreResp)
		return cachedExploreResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationDetail{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationDetail{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationDetail{}, err
	}

	exploreResp := RespLocationDetail{}
	err = json.Unmarshal(dat, &exploreResp)
	if err != nil {
		return RespLocationDetail{}, err
	}

	c.cache.Add(url, dat)

	return exploreResp, nil
}
