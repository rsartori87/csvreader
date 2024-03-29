package loader

import (
	"fmt"
)

type text struct {
	title string
	isbn string
	authors []Author
}

type book struct {
	SearchableText
	Description string
}

type megazine struct {
	SearchableText
	PublishedAt string
}

type SearchableText interface {
	Title() string
	Authors() []Author
	Print()
	Isbn() string
}

func (t text) Title() string {
	return t.title
}

func (t text) Authors() []Author {
	return t.authors
}

func (t text) Isbn() string {
	return t.isbn
}

func (t text) Print() {
	fmt.Printf("Title: %s\n", t.title)
	fmt.Printf("Isbn: %s\n", t.isbn)
	for _, a := range t.authors {
		a.Print()
	}
}

func (m megazine) Print() {
	m.SearchableText.Print()
	fmt.Printf("Published at: %s\n", m.PublishedAt)
}

func (b book) Print() {
	b.SearchableText.Print()
	fmt.Printf("Description:\n%s\n", b.Description)
}
