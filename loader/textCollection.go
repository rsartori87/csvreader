package loader

import "strings"

type textCollection []SearchableText

func (t textCollection) SearchByTitle(title *string) textCollection {
	if title == nil || len(strings.TrimSpace(*title)) == 0 {
		return t
	}
	result := textCollection{}
	for _, txt := range t {
		if txt.Title() == *title {
			result = append(result, txt)
		}
	}
	return result
}

func (t textCollection) SearchByIsbn(isbn *string) textCollection {
	if isbn == nil || len(strings.TrimSpace(*isbn)) == 0 {
		return t
	}
	results := textCollection{}
	for _, txt := range t {
		if txt.Isbn() == *isbn {
			results = append(results, txt)
		}
	}
	return results
}
