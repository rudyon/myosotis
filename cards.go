package main

type Cards struct {
}

func NewCards() *Cards {
	return &Cards{}
}

func (c *Cards) Get(text string) string {
	return text + "aaa"
}
