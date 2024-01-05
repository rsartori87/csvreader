package loader

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

const AUTHORS_FILE = "/authors.csv"
const BOOKS_FILE = "/books.csv"
const MAGAZINES_FILE = "/magazines.csv"

func ReadTexts(basePath string) textCollection {
	results := make(chan textCollection, 2)
	defer close(results)
	
	authors := readAuthors(basePath)
	
	output := textCollection{}

	go readBooks(basePath, authors, results)
	go readMagazines(basePath, authors, results)

	for i := 0; i < 2; i++ {
		r := <-results
		output = append(output, r...)
	}

	return output
}

func readAuthors(basePath string) []Author {
	f, err := os.Open(basePath + AUTHORS_FILE)
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

func readMagazines(basePath string, authors []Author, results chan textCollection) {
	f, err := os.Open(basePath + MAGAZINES_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := readLines(f)
	result := textCollection{}

	for _, line := range lines {
		m := megazine{
			SearchableText: text{
				title: line[0],
				isbn: line[1],
				authors: findAuthors(authors, line[2]),
			},
			PublishedAt: line[3],
		}

		result = append(result, m)
	}
	
	results <- result
}

func readBooks(basePath string, authors []Author, results chan textCollection) {
	f, err := os.Open(basePath + BOOKS_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lines := readLines(f)
	result := textCollection{}

	for _, line := range lines {
		b := book{
			SearchableText: text{
				title: line[0],
				isbn: line[1],
				authors: findAuthors(authors, line[2]),
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
