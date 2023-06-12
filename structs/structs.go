package main

import "fmt"

type StoryBook struct {
	page string 
	nextPage *StoryBook
}

func printStory(book *StoryBook) {
	fmt.Println(book.page)
	if(book.nextPage == nil) {
		return
	}
	printStory(book.nextPage)
}

func main() {
	page1 := StoryBook{"This is the first page of the book.", nil}
	page2 := StoryBook{"A knight doned in black aromor comes across a stranded child in the forest.", nil}
	page3 := StoryBook{"The lost child is overjoyed at the sight of her saviour.", nil}
	page1.nextPage = &page2
	page2.nextPage = &page3

	printStory(&page1)
}
