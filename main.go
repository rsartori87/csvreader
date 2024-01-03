package main

import (
	"fmt"
	"sort"

	"com.example/csvreader/loader"
)

func main() {
	texts := loader.ReadTexts("resources")

	sort.Slice(texts, func(i, j int) bool {
		return texts[i].Title() < texts[j].Title()
	})
	
	for _, t := range texts {
		t.Print()
		fmt.Println("")
	}

}
