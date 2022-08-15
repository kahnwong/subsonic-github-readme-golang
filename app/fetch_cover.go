package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

type CoverEnv struct {
	ID   string `url:"id"`
	Size int64  `url:"size"`
}

func GetCover(coverID string) string {
	// set params
	authParams, _ := query.Values(authEnv)

	coverEnv := CoverEnv{
		ID:   coverID,
		Size: 48,
	}
	coverParams, _ := query.Values(coverEnv)

	// fetch cover
	url := fmt.Sprintf("%s/rest/getCoverArt?%s&%s", baseURL, authParams.Encode(), coverParams.Encode())
	resp, err := http.Get(url)
	if err != nil {
		log.Println("No response from request")
	}
	defer resp.Body.Close()

	// convert to base64
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body")
	}

	return base64.StdEncoding.EncodeToString(body)
}
