package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Track struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Artist     string `json:"artist"`
	AlbumImg   string `json:"image"`
	Preview    string `json:"preview"`
	Popularity int    `json:"popularity"`
}

// mon Token Spotify
func GetToken() string {
	clientID := "34dccfbabfd5445682bfdd25c737f1a0"
	clientSecret := "317313b38c92415ebd0a18b12ebb71e0"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, _ := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	return res["access_token"].(string)
}

// mes 100 sons
func GetMyPlaylist() []Track {
	token := GetToken()
	playlistID := "5OKXoZW86C0inbf0dQqSzp"
	apiUrl := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", playlistID)

	req, _ := http.NewRequest("GET", apiUrl, nil)
	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	var data struct {
		Items []struct {
			Track struct {
				ID         string `json:"id"`
				Name       string `json:"name"`
				Preview    string `json:"preview_url"`
				Popularity int    `json:"popularity"`
				Album      struct {
					Images []struct {
						URL string `json:"url"`
					} `json:"images"`
				} `json:"album"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
			} `json:"track"`
		} `json:"items"`
	}

	json.NewDecoder(resp.Body).Decode(&data)

	var tracks []Track
	for _, item := range data.Items {
		tracks = append(tracks, Track{
			ID:         item.Track.ID,
			Name:       item.Track.Name,
			Artist:     item.Track.Artists[0].Name,
			AlbumImg:   item.Track.Album.Images[0].URL,
			Preview:    item.Track.Preview,
			Popularity: item.Track.Popularity,
		})
	}
	return tracks
}

func LoadFavorites() []Track {
	content, err := os.ReadFile("../favorites.json")
	if err != nil {
		return []Track{}
	}
	var favs []Track
	json.Unmarshal(content, &favs)
	return favs
}

func SaveToFavorites(id string) {
	track := GetTrackByID(id)
	favs := LoadFavorites()

	for _, f := range favs {
		if f.ID == id {
			return
		}
	}

	favs = append(favs, track)
	updatedData, _ := json.MarshalIndent(favs, "", "  ")

	os.WriteFile("../favorites.json", updatedData, 0644)
}

func GetTrackByID(id string) Track {
	all := GetMyPlaylist()
	for _, t := range all {
		if t.ID == id {
			return t
		}
	}
	return Track{}
}
