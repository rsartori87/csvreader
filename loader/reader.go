package loader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func ReadTexts() []SearchableText {
	results := make(chan []SearchableText, 2)
	defer close(results)
	
	authors := readAuthors()
	
	output := []SearchableText{}

	go readBooks(authors, results)
	go readMagazines(authors, results)

	for i := 0; i < 2; i++ {
		r := <-results
		output = append(output, r...)
	}

	return output
}

func readAuthors() []Author {
	f, err := os.Open("resources/authors.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := readLines(f)
	result := []Author{}

	for _, line := range lines {
		a := Author{
			Email: line[0],
			Firstname: line[1],
			Lastname: line[2],
		}
		result = append(result, a)
	}
	
	return result
}

func findAuthors(authors []Author, rawEmails string) []Author {
	emails := strings.Split(rawEmails, ",")
	result := []Author{}

	for _, e := range emails {
		aPtr := findAuthor(authors, e)
		if aPtr != nil {
			result = append(result, *aPtr)
		}
	}

	return result
}

func findAuthor(authors []Author, email string) *Author {
	i := slices.IndexFunc(authors, func(a Author) bool {
		return a.Email == email
	})
	if i == -1 {
		return nil
	}

	return &authors[i]
}

func readMagazines(authors []Author, results chan []SearchableText) {
	f, err := os.Open("resources/magazines.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := readLines(f)
	result := []SearchableText{}

	for _, line := range lines {
		m := megazine{
			SearchableText: text{
				title: line[0],
				Isbn: line[1],
				Authors: findAuthors(authors, line[2]),
			},
			PublishedAt: line[3],
		}

		result = append(result, m)
	}
	
	results <- result
}

func readBooks(authors []Author, results chan []SearchableText) {
	f, err := os.Open("resources/books.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := readLines(f)
	result := []SearchableText{}

	for _, line := range lines {
		b := book{
			SearchableText: text{
				title: line[0],
				Isbn: line[1],
				Authors: findAuthors(authors, line[2]),
			},
			Description: line[3],
		}

		result = append(result, b)
	}

	results <- result
}

func readLines(file *os.File) [][]string {
	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'
	lines := [][]string{}
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, rec)
	}
	return lines[1:]
}
