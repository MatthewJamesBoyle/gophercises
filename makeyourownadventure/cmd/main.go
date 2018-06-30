package main

import (
	"flag"
	"fmt"
	"gophercises/makeyourownadventure"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "port")
	file := flag.String("file", "story.json", "The json file for the choose your own adventure story.")
	flag.Parse()
	fmt.Printf("Using the story in %s \n", *file)

	f, err := os.Open(*file)

	story, err := makeyourownadventure.JsonStory(f)

	if err != nil {
		panic(err)
	}

	h := makeyourownadventure.NewHandler(story, nil)

	fmt.Printf("Starting the server on %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
