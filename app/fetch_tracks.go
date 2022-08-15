package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"log"

	"github.com/google/go-querystring/query"
)

func FetchRandomSongs(authEnv subsonicAuth) [5]Song {
	log.Println("Fetching random songs")

	// fetch response
	authParams, _ := query.Values(authEnv)
	url := fmt.Sprintf("%s/rest/getRandomSongs?%s", baseURL, authParams.Encode())
	resp, err := http.Get(url)
	if err != nil {
		log.Println("No response from request")
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
	}

	var result RandomSongs
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	// parse response
	tracksRaw := result.SubsonicResponse.RandomSongs.Song[:5]

	tracks := [5]Song{}
	for i, rec := range tracksRaw {
		tracks[i] = rec
	}

	return tracks
}

func FetchNowPlaying(authEnv subsonicAuth) []NowPlayingSong {
	log.Println("Fetching now playing")

	// fetch response
	authParams, _ := query.Values(authEnv)
	url := fmt.Sprintf("%s/rest/getNowPlaying?%s", baseURL, authParams.Encode())
	resp, err := http.Get(url)
	if err != nil {
		log.Println("No response from request")
	}
	defer resp.Body.Close()

	// read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
	}

	var result NowPlaying
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Can not unmarshal JSON")
	}

	// parse response
	tracksRaw := result.SubsonicResponse.NowPlaying.Entry

	if len(tracksRaw) == 1 { // if has recently played tracks
		tracks := []NowPlayingSong{}

		for _, rec := range tracksRaw {
			tracks = append(tracks, rec)
		}

		return tracks
	} else {
		return []NowPlayingSong{}
	}

}
