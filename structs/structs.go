// TODO: 1) add a func that will add a new page
// 2) add a func that will delete a page

package main

import (
	"bufio"
	"fmt"
	"os"
)

type StoryBook struct {
	page     string
	nextPage *StoryBook
}

func printStory(book *StoryBook) {
	fmt.Println(book.page)
	if book.nextPage == nil {
		fmt.Println("There are no more pages.")
		fmt.Println("Would you like to add a new page? y/n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response := scanner.Text()
		if response == "y" {
			addPage(book)
		} else {
			return
		}
	}
	printStory(book.nextPage)
}

func addPage(book *StoryBook) {
	fmt.Println("Add content for the page")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	content := scanner.Text()
	newPage := StoryBook{content, nil}
	book.nextPage = &newPage
}

func deletePage() {

}

func main() {
	page1 := StoryBook{"This is the first page of the book.", nil}
	page2 := StoryBook{"A knight doned in black aromor comes across a stranded child in the forest.", nil}
	page3 := StoryBook{"The lost child is overjoyed at the sight of her saviour.", nil}
	page1.nextPage = &page2
	page2.nextPage = &page3

	printStory(&page1)
}
