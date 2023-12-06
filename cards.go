package main

import (
	"fmt"
	"time"
)

type Cards struct {
}

func NewCards() *Cards {
	return &Cards{}
}

func (c *Cards) GetCardFront(id int) string {
	return dueCards[id].Front
}

func (c *Cards) GetCardBack(id int) string {
	return dueCards[id].Back
}

func (c *Cards) UpdateCard(id int, correct bool) {
	if id < 0 || id >= len(dueCards) {
		fmt.Println("Invalid card ID")
		return
	}

	// Update LastReview based on the current time
	dueCards[id].LastReview = int(time.Now().Unix())

	if correct {
		// If the user answered correctly, update the interval based on the Leitner Box algorithm
		switch dueCards[id].Interval {
		case 0:
			dueCards[id].Interval = 86400 // 1 day
		default:
			dueCards[id].Interval *= 2 // Double the interval
		}
	} else {
		// If the user answered incorrectly, reset the interval to the initial value
		dueCards[id].Interval = 0
	}

	// Reflect the changes in the original flashcards slice
	updateOriginalFlashcard(&dueCards[id])

	// Save flashcards to the file
	err := saveFlashCards("./config.json")
	if err != nil {
		fmt.Println("Error saving flash cards:", err)
	}
}

func updateOriginalFlashcard(updatedCard *FlashCard) {
	// Find the corresponding card in the original flashcards slice and update it
	for i, card := range flashcards {
		if card.Front == updatedCard.Front && card.Back == updatedCard.Back {
			flashcards[i] = *updatedCard
			break
		}
	}
}

func (c *Cards) GetTotalCards() int {
	return TotalCards
}
