package main

import (
	"flag"
	"fmt"
	"sort"

	"com.example/csvreader/loader"
)

func main() {

	title := flag.String("title", "", "title to be searched")
	isbn := flag.String("isbn", "", "isbn to be searched")
	flag.Parse()
	

	texts := loader.ReadTexts("resources")

	sort.Slice(texts, func(i, j int) bool {
		return texts[i].Title() < texts[j].Title()
	})
	
	for _, t := range texts.SearchByTitle(title).SearchByIsbn(isbn) {
		t.Print()
		fmt.Println("")
	}

}
