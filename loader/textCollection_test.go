package loader

import "testing"

func TestSearchByTitleReturnEmptyIfNoResultsAreFound(t *testing.T) {
	aText := text{
		title: "a title",
	}
	texts := textCollection{aText}
	title := "another title"
	results := texts.SearchByTitle(&title)
	if len(results) != 0 {
		t.Fatalf("Expected empty slice got a slice with len %d", len(results))
	}
}

func TestSearchByTitleReturnSearchedElements(t *testing.T) {
	aText := text{
		title: "a title",
	}
	texts := textCollection{aText}
	title := "a title"
	results := texts.SearchByTitle(&title)
	if len(results) != 1 {
		t.Fatalf("Expected a slice with one element got a slice of length %d", len(results))
	}
}

func TestSearchByTitleReturnAllElementsForEmptyTitle(t *testing.T) {
	aText := text{
		title: "a title",
	}
	texts := textCollection{aText}
	title := "    "
	results := texts.SearchByTitle(&title)
	if len(results) != 1 {
		t.Fatalf("Expected a slice with one element got a slice of length %d", len(results))
	}
}

func TestSearchByTitleReturnAllElementsForNilTitle(t *testing.T) {
	aText := text{
		title: "a title",
	}
	texts := textCollection{aText}
	results := texts.SearchByTitle(nil)
	if len(results) != 1 {
		t.Fatalf("Expected a slice with one element got a slice of length %d", len(results))
	}	
}

func TestSearchByIsbnReturnEmptyIfNoResultsAreFound(t *testing.T) {
	aText := text{
		isbn: "42",
	}
	texts := textCollection{aText}
	isbn := "22"
	results := texts.SearchByIsbn(&isbn)
	if len(results) != 0 {
		t.Fatalf("Expected empty slice got a slice with len %d", len(results))
	}
}

func TestSearchByIsbnReturnSearchedElements(t *testing.T) {
	aText := text{
		isbn: "42",
	}
	texts := textCollection{aText}
	isbn := "42"
	results := texts.SearchByIsbn(&isbn)
	if len(results) != 1 {
		t.Fatalf("Expected empty slice with one element got a slice og length %d", len(results))
	}
}

func TestSearchByIsbnReturnAllElementForEmptyIsbn(t *testing.T) {
	aText := text{
		isbn: "42",
	}
	texts := textCollection{aText}
	isbn := "    "
	results := texts.SearchByIsbn(&isbn)
	if len(results) != 1 {
		t.Fatalf("Expected empty slice with one element got a slice og length %d", len(results))
	}
}

func TestSearchByIsbnReturnAllElementForNilIsbn(t *testing.T) {
	aText := text{
		isbn: "42",
	}
	texts := textCollection{aText}
	results := texts.SearchByIsbn(nil)
	if len(results) != 1 {
		t.Fatalf("Expected empty slice with one element got a slice og length %d", len(results))
	}
}
