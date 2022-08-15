package main

import (
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type TrackInfo struct {
	Title             string
	Album             string
	Artist            string
	Base64ImageString string
}

func TrackInfoTemplate() *template.Template {
	tmpl, err := template.ParseFiles("template/track-info.svg")
	if err != nil {
		log.Println("Template file doesn't exist")
	}

	return tmpl
}

func GenerateTrackInfo(track TrackInfo, filename string) {
	log.Printf("Generating %s.svg...", filename)

	outPath := path.Join("output", filename+".svg")
	file, err := os.Create(outPath)
	if err != nil {
		log.Println("Error creating file")
	}
	defer file.Close()

	err = TrackInfoTemplate().Execute(file, track)
	if err != nil {
		log.Println("Error executing template")
	}
}

// prob need to find a way to reduce code duplicate
func ParseTrack(song Song) TrackInfo {
	track := TrackInfo{}

	track.Title = strings.Replace(song.Title, "&", "&amp;", -1)
	track.Album = strings.Replace(song.Album, "&", "&amp;", -1)
	track.Artist = strings.Replace(song.Artist, "&", "&amp;", -1)
	track.Base64ImageString = GetCover(song.CoverArt)

	return track
}

func ParseNowPlaying(song NowPlayingSong) TrackInfo {
	track := TrackInfo{}

	track.Title = strings.Replace(song.Title, "&", "&amp;", -1)
	track.Album = strings.Replace(song.Album, "&", "&amp;", -1)
	track.Artist = strings.Replace(song.Artist, "&", "&amp;", -1)
	track.Base64ImageString = GetCover(song.CoverArt)

	return track
}
