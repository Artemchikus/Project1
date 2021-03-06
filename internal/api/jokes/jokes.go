package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Artemchikus/internal/api"
)

const getJockePath = "/api?format=json"

type JokeClient struct {
	url string
}

func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url : baseUrl,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url+getJockePath
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("API request status: %s", http.StatusText(resp.StatusCode))
	}

	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}