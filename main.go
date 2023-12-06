package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type FlashCard struct {
	Front      string `json:"front"`
	Back       string `json:"back"`
	LastReview int    `json:"lastReview"`
	Interval   int    `json:"interval"`
}

var flashcards []FlashCard
var dueCards []FlashCard
var TotalCards int

func loadFlashCards(filename string) error {
	// Read the JSON file
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into a slice of FlashCard structs
	err = json.Unmarshal(fileContent, &flashcards)
	if err != nil {
		return err
	}

	return nil
}

func selectDueCards(flashcards []FlashCard) []FlashCard {
	var dueCards []FlashCard
	currentTime := int(time.Now().Unix())

	for _, card := range flashcards {
		dueTime := card.LastReview + card.Interval
		if currentTime >= dueTime {
			dueCards = append(dueCards, card)
		}
	}

	return dueCards
}

func saveFlashCards(filename string) error {
	// Marshal flashcards to JSON
	fileContent, err := json.MarshalIndent(flashcards, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling flashcards:", err)
		return err
	}

	// Write JSON data to the file
	err = os.WriteFile(filename, fileContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("Flashcards saved successfully.")
	return nil
}

func main() {
	err := loadFlashCards("./config.json")
	if err != nil {
		fmt.Println("Error loading flash cards:", err)
		return
	}

	dueCards = selectDueCards(flashcards)
	TotalCards = len(dueCards)

	// Create an instance of the app structure
	app := NewApp()
	cards := NewCards()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "what",
		Width:  800,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app, cards,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
