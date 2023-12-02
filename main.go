package main

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width, height           int32   = 600, 600
	font_size               int32   = 32
	card_width, card_height float32 = float32(width - 200), float32(height - 200)
)

var flipped bool
var flipping bool
var flip_percentage float32 = 1

type Card struct {
	Front string
	Back  string
}

func main() {

	rl.InitWindow(width, height, "leitner")
	rl.SetTargetFPS(60)

	// Example usage
	myCard := Card{
		Front: "The pioneer of the experimental science of memory was...",
		Back:  "Hermann Ebbinghaus",
	}

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) {
			flipping = true
		}

		if flipping {
			if flip_percentage > 0 {
				flip_percentage -= 0.1
			} else {
				flip_percentage = 1
				flipping = false
				flipped = !flipped
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if flipped {
			if flipping {
				drawCard(myCard.Back, flip_percentage, colorBrightness(rl.White, -flip_percentage))
				drawCard(myCard.Front, 1-flip_percentage, colorBrightness(rl.White, flip_percentage))
			} else {
				drawCard(myCard.Back, flip_percentage, rl.White)
			}
		} else {
			if flipping {
				drawCard(myCard.Front, flip_percentage, colorBrightness(rl.White, -flip_percentage))
				drawCard(myCard.Back, 1-flip_percentage, colorBrightness(rl.White, flip_percentage))
			} else {
				drawCard(myCard.Front, flip_percentage, rl.White)
			}
		}

		rl.EndDrawing()
	}
}

func colorBrightness(color rl.Color, brightness float32) rl.Color {
	brightness = brightness * 150
	return rl.NewColor(color.R*uint8(brightness), color.G*uint8(brightness), color.B*uint8(brightness), color.A)
}

func drawCard(text string, x_percentage float32, card_color rl.Color) {
	// Ensure x_percentage is within the range [0.0, 1.0]
	if x_percentage < 0.0 {
		x_percentage = 0.0
	} else if x_percentage > 1.0 {
		x_percentage = 1.0
	}

	// Scale the card width with x_percentage
	scaledCardWidth := float32(card_width) * x_percentage

	rl.DrawRectanglePro(rl.NewRectangle(float32(width)/2, float32(height)/2, scaledCardWidth, float32(card_height)), rl.NewVector2(scaledCardWidth/2, float32(card_height)/2), 0, card_color)

	lines := wrapText(text, card_width, font_size)

	// Calculate the vertical space between lines
	line_spacing := float32(font_size) + 5 // Convert font_size to float32

	// Draw each line
	for i, line := range lines {
		text_width := rl.MeasureText(line, font_size)
		text_x := (float32(width) - float32(text_width)) / 2 // Convert width to float32
		text_y := (float32(height)-float32(len(lines))*line_spacing)/2 + float32(i)*line_spacing

		// Convert text_y to int32
		rl.DrawText(line, int32(text_x), int32(text_y), font_size, rl.Black)
	}
}

func wrapText(text string, max_width float32, font_size int32) []string {
	words := strings.Fields(text)
	var lines []string
	current_line := ""

	for _, word := range words {
		// Check if adding the word to the current line exceeds the maximum width
		if rl.MeasureText(current_line+" "+word, font_size) <= int32(max_width) {
			// Add the word to the current line
			if current_line != "" {
				current_line += " "
			}
			current_line += word
		} else {
			// Start a new line with the current word
			lines = append(lines, current_line)
			current_line = word
		}
	}

	// Add the last line
	lines = append(lines, current_line)

	return lines
}
