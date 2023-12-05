package main

type Cards struct {
}

func NewCards() *Cards {
	return &Cards{}
}

func (c *Cards) GetCardFront(text string) string {
	return "LOREM"
}

func (c *Cards) GetCardBack(text string) string {
	return "IPSUM"
}
