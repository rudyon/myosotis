package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type FlashCard struct {
	Back  string `json:"back"`
	Front string `json:"front"`
}

var flashcards []FlashCard
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

func main() {
	err := loadFlashCards("./config.json")
	if err != nil {
		fmt.Println("Error loading flash cards:", err)
		return
	}

	// Create an instance of the app structure
	app := NewApp()
	cards := NewCards()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "what",
		Width:  1024,
		Height: 768,
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
