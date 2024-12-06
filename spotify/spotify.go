package spotify

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var token string
var tokenExpiry time.Time

type AlbumsData struct {
	Items []struct {
		TotalTracks int `json:"total_tracks"`
		Images      []struct {
			URL string `json:"url"`
		} `json:"images"`
		Name        string `json:"name"`
		ReleaseDate string `json:"release_date"`
	} `json:"items"`
}

type TracksData struct {
	Album struct {
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
		Name        string `json:"name"`
		ReleaseDate string `json:"release_date"`
	} `json:"album"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Name string `json:"name"`
}

type SpotifyToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func getToken() (string, error) {
	if time.Now().Before(tokenExpiry) {
		return token, nil
	}

	clientID := "104aba1aa6914956a3aca6d153e468aa"
	clientSecret := "7ece27b239954dadbe95942b46d70b88"

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = "grant_type=client_credentials"

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get token")
	}

	var result SpotifyToken
	json.NewDecoder(resp.Body).Decode(&result)

	token = result.AccessToken
	tokenExpiry = time.Now().Add(time.Duration(result.ExpiresIn) * time.Second)

	return token, nil
}

func GetAlbums() (AlbumsData, error) {
	token, err := getToken()
	if err != nil {
		return AlbumsData{}, err
	}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/artists/2kXKa3aAFngGz2P4GjG5w2/albums", nil)
	if err != nil {
		return AlbumsData{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return AlbumsData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AlbumsData{}, errors.New("failed to get albums")
	}

	var result AlbumsData
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

func GetTrackDetails() (TracksData, error) {
	token, err := getToken()
	if err != nil {
		return TracksData{}, err
	}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/tracks/1Mzg6bu3hkCwJKEf7v49MN", nil)
	if err != nil {
		return TracksData{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TracksData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return TracksData{}, errors.New("failed to get track details")
	}

	var result TracksData
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}
