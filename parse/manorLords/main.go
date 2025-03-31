package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL to parse
	url := "https://store.steampowered.com/app/1363080/Manor_Lords/"

	// Fetch the webpage
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching the URL %v", err)
	}
	defer resp.Body.Close()

	// Load the HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Error loading the URL %v", err)
	}

	// Get the title using Find
	title := doc.Find("title").Text()
	fmt.Printf("\nTitle: %s\n", title)

	// Find the div with the class name 'game_area_purchase'
}
