package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width, height         int32   = 600, 600
	fontSize              int32   = 32
	cardWidth, cardHeight float32 = float32(width - 200), float32(height - 200)
)

var (
	flipped        bool
	flipping       bool
	flipPercentage float32 = 1
	cards          []Card
	card           Card
	delta_time     = rl.GetFrameTime() * float32(rl.GetFPS())
)

type Card struct {
	Front      string
	Back       string
	LastReview int64
	Interval   int
}

func main() {
	rl.InitWindow(width, height, "memory")

	loadCards()

	cardIndex := -1

	for !rl.WindowShouldClose() {
		delta_time = rl.GetFrameTime() * float32(rl.GetFPS())

		handleInput(&cardIndex)

		updateCardIndex(&cardIndex)

		draw()

		saveCards()
	}
}

func loadCards() {
	fileData, err := os.ReadFile("cards.json")
	if err != nil {
		fmt.Println("Error loading cards:", err)
		return
	}

	err = json.Unmarshal(fileData, &cards)
	if err != nil {
		fmt.Println("Error unmarshalling cards:", err)
		return
	}
}

func saveCards() {
	fileData, err := json.MarshalIndent(cards, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling cards:", err)
		return
	}

	err = os.WriteFile("cards.json", fileData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func handleInput(cardIndex *int) {

	if rl.IsKeyPressed(rl.KeySpace) {
		flipping = true
	}

	if rl.IsKeyPressed(rl.KeyG) {
		recallGood(cardIndex)
	}

	if rl.IsKeyPressed(rl.KeyB) {
		recallBad(cardIndex)
	}
}

func recallGood(cardIndex *int) {
	if cards[*cardIndex].Interval == -1 {
		cards[*cardIndex].Interval = 2
	} else {
		cards[*cardIndex].Interval *= 2
	}

	cards[*cardIndex].LastReview = time.Now().Unix()

	if *cardIndex < len(cards) {
		handleCardIndex(cardIndex)
	} else {
		*cardIndex = -1
	}
}

func recallBad(cardIndex *int) {
	cards[*cardIndex].Interval = 1

	cards[*cardIndex].LastReview = time.Now().Unix()

	if *cardIndex < len(cards) {
		handleCardIndex(cardIndex)
	} else {
		*cardIndex = -1
	}
}

func handleCardIndex(cardIndex *int) {
	for *cardIndex < len(cards) {
		if cards[*cardIndex].Interval == -1 || time.Since(time.Unix(cards[*cardIndex].LastReview, 0)).Minutes() > float64(cards[*cardIndex].Interval) {
			// If the condition is met, break out of the loop
			break
		}

		// If the condition is not met, increment the card index and print the card
		fmt.Println(cards[*cardIndex])
		*cardIndex++
	}

	if *cardIndex < len(cards) {
		flipping = true
		flipped = true
	} else {
		*cardIndex = -1
	}
}

func updateCardIndex(cardIndex *int) {
	if *cardIndex >= len(cards) {
		*cardIndex = 0
	}

	if *cardIndex != -1 {
		card = cards[*cardIndex]
	} else {
		card = Card{Front: "There are no more cards.", Back: "Congratulations!!!", LastReview: -1, Interval: -1}
	}
}

func draw() {
	if flipping {
		updateFlipPercentage()
	}

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// Draw the review count at the top-left corner
	drawReviewCount()

	if flipped {
		drawFlippedCard()
	} else {
		drawUnflippedCard()
	}

	rl.EndDrawing()
}

func drawReviewCount() {
	count := 0
	for i := 0; i < len(cards); i++ {
		if cards[i].Interval == -1 || time.Since(time.Unix(cards[i].LastReview, 0)).Minutes() > float64(cards[i].Interval) {
			count++
		}
	}

	countText := fmt.Sprintf("Reviews waiting: %d", count)
	rl.DrawText(countText, 10, 10, fontSize, rl.White)
}

func updateFlipPercentage() {
	if flipPercentage > 0 {
		flipPercentage -= 0.015 * delta_time
	} else {
		flipPercentage = 1
		flipping = false
		flipped = !flipped
	}
}

func drawFlippedCard() {
	if flipping {
		drawCard(card.Back, flipPercentage, colorBrightness(rl.White, -flipPercentage))
		drawCard(card.Front, 1-flipPercentage, colorBrightness(rl.White, flipPercentage))
	} else {
		drawCard(card.Back, flipPercentage, rl.White)
	}
}

func drawUnflippedCard() {
	if flipping {
		drawCard(card.Front, flipPercentage, colorBrightness(rl.White, -flipPercentage))
		drawCard(card.Back, 1-flipPercentage, colorBrightness(rl.White, flipPercentage))
	} else {
		drawCard(card.Front, flipPercentage, rl.White)
	}
}

func colorBrightness(color rl.Color, brightness float32) rl.Color {
	brightness = brightness * 150
	return rl.NewColor(color.R*uint8(brightness), color.G*uint8(brightness), color.B*uint8(brightness), color.A)
}

func drawCard(text string, xPercentage float32, cardColor rl.Color) {
	// Ensure xPercentage is within the range [0.0, 1.0]
	if xPercentage < 0.0 {
		xPercentage = 0.0
	} else if xPercentage > 1.0 {
		xPercentage = 1.0
	}

	// Scale the card width with xPercentage
	scaledCardWidth := float32(cardWidth) * xPercentage

	rl.DrawRectanglePro(rl.NewRectangle(float32(width)/2, float32(height)/2, scaledCardWidth, float32(cardHeight)), rl.NewVector2(scaledCardWidth/2, float32(cardHeight)/2), 0, cardColor)

	lines := wrapText(text, cardWidth-float32(fontSize), fontSize)

	// Calculate the vertical space between lines
	lineSpacing := float32(fontSize) + 5 // Convert fontSize to float32

	// Draw each line
	for i, line := range lines {
		textWidth := rl.MeasureText(line, fontSize)
		textX := (float32(width) - float32(textWidth)) / 2 // Convert width to float32
		textY := (float32(height)-float32(len(lines))*lineSpacing)/2 + float32(i)*lineSpacing

		// Convert textY to int32
		rl.DrawText(line, int32(textX), int32(textY), fontSize, rl.Black)
	}
}

func wrapText(text string, maxWidth float32, fontSize int32) []string {
	words := strings.Fields(text)
	var lines []string
	currentLine := ""

	for _, word := range words {
		// Check if adding the word to the current line exceeds the maximum width
		if rl.MeasureText(currentLine+" "+word, fontSize) <= int32(maxWidth) {
			// Add the word to the current line
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		} else {
			// Start a new line with the current word
			lines = append(lines, currentLine)
			currentLine = word
		}
	}

	// Add the last line
	lines = append(lines, currentLine)

	return lines
}
