package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// method to the Client Struct
// make get request to /location-area
// handles all possible errors
// return locationAreasRes, error
func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasRes, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageUrl != nil {
		fullURL = *pageUrl
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasRes{}, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasRes{}, err
	}

	// A defer statement defers the execution of a function until the surrounding function returns.
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasRes{}, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return LocationAreasRes{}, err
	}

	locationAreasRes := LocationAreasRes{}

	err = json.Unmarshal(data, &locationAreasRes)

	if err != nil {
		return LocationAreasRes{}, err
	}

	return locationAreasRes, nil
}
