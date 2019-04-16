package main

import (
	"fmt"
)

type storyPage struct {
	text string
	// nextPage storyPage // cannot have a recursive type in struct
	nextPage *storyPage // instead, we use a pointer
}

// functions can have receivers
func (page *storyPage) playStory() {
	if page == nil {
		return // no return type just exits
	}
	fmt.Println(page.text)
	// playStory(page.nextPage) // receivers - instead of this
	page.nextPage.playStory()
}

/* example receiver
func (num int) square() {
	return num*num
}
5.square()
*/

func (prevPage *storyPage) insertPage(newPage *storyPage) {
	if prevPage == nil {
		return
	}
	newPage.nextPage = prevPage.nextPage
	prevPage.nextPage = newPage
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	page1 := storyPage{"Adventure start-o!", nil}
	page2 := storyPage{"We begin the story with our princess, walking down the street.", nil}
	page3 := storyPage{"She is on the quest to find the mythical 'pou-teen'.", nil}
	page4 := storyPage{"Suddenly, a wo pops out of the wild grass!", nil}
	
	page1.nextPage = &page2 // nextPage references the address of the next page
	page2.nextPage = &page3 // nextPage references the address of the next page
	page3.nextPage = &page4 // nextPage references the address of the next page

	// playStory(&page1)
	page1.playStory()
	// scanner.Scan()

	page5 := storyPage{"Wo attacks you for 5 damage", &page4}
	page2.insertPage(&page5)
	// page4.nextPage = nil

	// playStory(&page1)
	page1.playStory()
}