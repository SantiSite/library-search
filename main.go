package main

import (
	"fmt"
	"strings"
)

/**
Given a flat file of book metadata, write a Library function that parses the book data and provides an API that lets you
search for all books containing a word.

Your solution must have a runtime of faster than O(n) for the search api, with n representing the number of words in
the library.

API:

Library
  - <constructor>(input) -> returns a Library object
  - search(word) -> returns book titles that contain the word anywhere in the
    title, author, or description fields. Only matches *whole* words.
    E.g. Searching for "My" or "book" would match a book containing
    "My book", but searching for "My b" or "boo" would *not* match.
*/

type book struct {
	title, author, description string
}

type library struct {
	books []book
}

// replacePunctuationMarks returns the text string without punctuation marks.
func replacePunctuationMarks(world string) string {
	world = strings.Replace(world, ".", "", 1)
	world = strings.Replace(world, ",", "", 1)
	return world
}

// search find the keyword in the library
func (l *library) search(keyword string) []string {
	keyword = replacePunctuationMarks(keyword)
	var titles []string
	for _, book := range l.books {
		title := replacePunctuationMarks(book.title)
		author := replacePunctuationMarks(book.author)
		description := replacePunctuationMarks(book.description)
		if strings.Contains(strings.ToLower(title), " "+keyword+" ") ||
			strings.Contains(strings.ToLower(author), " "+keyword+" ") ||
			strings.Contains(strings.ToLower(description), " "+keyword+" ") {
			titles = append(titles, book.title)
		}
	}
	if len(titles) > 0 {
		return titles
	}
	return nil
}

// getBooks get the books from the data strings.
func getBooks(libraryData string) []book {
	// Separate each book and store it in a slice.
	listOfRawBooks := strings.Split(libraryData, "\n\n")

	var books []book

	for _, item := range listOfRawBooks {
		books = append(books, book{
			title:       strings.TrimSpace(item[7:strings.Index(item, "\nA")]),
			author:      strings.TrimSpace(item[strings.Index(item, "AUTHOR: ")+8 : strings.Index(item, "\nD")]),
			description: strings.TrimSpace(item[strings.Index(item, "DESCRIPTION: ")+13:]),
		})
	}
	return books
}

// printTitles displays search results.
func printTitles(titles []string, keyword string) {
	defer fmt.Println("")
	if titles == nil {
		fmt.Println("No books were found matching the word " + `"` + keyword + `"`)
	} else {
		fmt.Println("Books containing the word " + `"` + keyword + `"` + ":")
		for _, title := range titles {
			fmt.Println("-", title)
		}
	}
}

func main() {
	libraryData := `
TITLE: Hitchhiker's Guide to the Galaxy
AUTHOR: Douglas Adams
DESCRIPTION: Seconds before the Earth is demolished to make way for a galactic freeway,
Arthur Dent is plucked off the planet by his friend Ford Prefect, a researcher for the
revised edition of The Hitchhiker's Guide to the Galaxy who, for the last fifteen years,
has been posing as an out-of-work actor.

TITLE: Dune
AUTHOR: Frank Herbert
DESCRIPTION: The troubles begin when stewardship of Arrakis is transferred by the
Emperor from the Harkonnen Noble House to House Atreides. The Harkonnens don't want to
give up their privilege, though, and through sabotage and treachery they cast young
Duke Paul Atreides out into the planet's harsh environment to die. There he falls in
with the Fremen, a tribe of desert dwellers who become the basis of the army with which
he will reclaim what's rightfully his. Paul Atreides, though, is far more than just a
usurped duke. He might be the end product of a very long-term genetic experiment
designed to breed a super human; he might be a messiah. His struggle is at the center
of a nexus of powerful people and events, and the repercussions will be felt throughout
the Imperium.

TITLE: A Song Of Ice And Fire Series
AUTHOR: George R.R. Martin
DESCRIPTION: As the Seven Kingdoms face a generation-long winter, the noble Stark family
confronts the poisonous plots of the rival Lannisters, the emergence of the
White Walkers, the arrival of barbarian hordes, and other threats.
`

	// Test words
	keyword1 := "Arrakis"
	keyword2 := "winter"
	keyword3 := "demolished"
	keyword4 := "the"
	keyword5 := "A S"

	// Creation of the library object
	books := getBooks(libraryData)
	newLibrary := library{books: books}

	// Search for the word
	titles1 := newLibrary.search(strings.ToLower(keyword1))
	titles2 := newLibrary.search(strings.ToLower(keyword2))
	titles3 := newLibrary.search(strings.ToLower(keyword3))
	titles4 := newLibrary.search(strings.ToLower(keyword4))
	titles5 := newLibrary.search(strings.ToLower(keyword5))

	printTitles(titles1, keyword1)
	printTitles(titles2, keyword2)
	printTitles(titles3, keyword3)
	printTitles(titles4, keyword4)
	printTitles(titles5, keyword5)
}
