package main

import (
	"fmt"
	"log"
	"os"
)

var baseURL = os.Getenv("BASE_URL")

func main() {
	err := os.MkdirAll("output", os.ModePerm)
	if err != nil {
		log.Println("Error creating output directory")
	}

	if authEnv.Token == "" {
		log.Panic("Token is not set")
	}

	randomSongs := FetchRandomSongs(authEnv)
	for i, song := range randomSongs {
		track := ParseTrack(song)

		filename := fmt.Sprintf("random-song-%d", i)
		GenerateTrackInfo(track, filename)
	}

	nowPlaying := FetchNowPlaying(authEnv)
	if len(nowPlaying) == 0 { // no recently played tracks
		log.Println("No recently played tracks")

		src := "template/now-playing-null.svg"
		dest := "output/now-playing.svg"

		bytesRead, err := os.ReadFile(src)

		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(dest, bytesRead, 0644)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		nowPlayingTrack := ParseNowPlaying(nowPlaying[0]) // returns one item only
		GenerateTrackInfo(nowPlayingTrack, "now-playing")
	}
}
