package loader

import "testing"

const BASE_PATH = "../resources"

func TestReadTexts(t *testing.T) {
	texts := ReadTexts(BASE_PATH)
	length := len(texts)
	if length != 14 {
		t.Fatalf("Expected slice with length 14 got one with length %d", length)
	}
}

func TestItPopulateAllAuthors(t *testing.T) {
	texts := ReadTexts(BASE_PATH)
	for _, txt := range texts {
		authorsCount := len(txt.Authors())
		if authorsCount == 0 {
			t.Fatalf("Found empty author for text %s", txt.Title())
		}
	}
}
