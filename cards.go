package main

type Cards struct {
}

func NewCards() *Cards {
	return &Cards{}
}

func (c *Cards) GetCardFront(id int) string {
	return flashcards[id].Front
}

func (c *Cards) GetCardBack(id int) string {
	return flashcards[id].Back
}
