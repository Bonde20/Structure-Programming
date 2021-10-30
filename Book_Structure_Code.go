
package main

import (
	"fmt"
)

func main() {
	b := &Book{
		235612,
		"janzo",
		"James",
		"1-5682-654-40",
		"Cinderella",
	}
	b.BookReplacement("Snow-white")

	b.PrintBookInfos()
}

type Book struct {
	BookId                          int
	AuthorFirstName, AuthorLastName string
	ISBN                            string
	InStock                         string
}

func (b *Book) PrintBookInfos() {
	fmt.Printf("Bookid is:%d |AuthorFirstName is:%s | AuthorLastName is:%s | The ISBN is: %s |Book inStock:%s",
		b.BookId, b.AuthorFirstName, b.AuthorLastName, b.ISBN, b.InStock)
}

func (b *Book) BookReplacement(newBook string) {

	b.InStock = newBook
}
