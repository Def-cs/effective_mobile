package external

import (
	"effective_mobile/pkg/dto"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrBadRequest  = errors.New("incorrect request")
	ErrNoResponse  = errors.New("no response from API")
	ErrDecodeError = errors.New("error decoding response")
)

type ApiClient struct {
	baseURL string
	client  *http.Client
}

func NewApiClient(baseURL string) *ApiClient {
	return &ApiClient{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (ac *ApiClient) GetSongInfo(group, song string) (*dto.SongDetail, error) {
	url := fmt.Sprintf("%s/info?group=%s&song=%s", ac.baseURL, group, song)

	resp, err := ac.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusBadRequest:
			return nil, ErrBadRequest
		default:
			return nil, ErrNoResponse
		}
	}

	var songDetail dto.SongDetail
	if err = json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, err
	}

	return &songDetail, nil
}
