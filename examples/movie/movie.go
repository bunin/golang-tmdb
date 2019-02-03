package main

import (
	"fmt"
	"os"

	"github.com/cyruzin/golang-tmdb"
)

func main() {
	var tmdbClient tmdb.Client
	tmdbClient.APIKey = os.Getenv("APIKey")

	movie, err := tmdbClient.GetMovieDetails(299536, nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(movie.Title)

	fmt.Println("------")

	// With options
	options := make(map[string]string)
	options["append_to_response"] = "credits"
	options["language"] = "pt-BR"

	movie, err = tmdbClient.GetMovieDetails(299536, options)

	if err != nil {
		fmt.Println(err)
	}

	// Credits - Iterate cast from append to response.
	for _, v := range movie.MovieCreditsShort.Credits.Cast {
		fmt.Println(v.Name)
	}
}
