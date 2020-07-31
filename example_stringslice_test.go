// Copyright (c) 2020 Dean Jackson <deanishe@deanishe.net>
// MIT Licence applies http://opensource.org/licenses/MIT
// Created on 2020-07-31

package fuzzy

import "fmt"

// Sort a slice of strings by fuzzy match.
func ExampleSortStrings() {
	squad := []string{
		"Alisson Becker",
		"Adrián",
		"Andy Lonergan",
		"Caoimhín Kelleher",
		"Virgil van Dijk",
		"Joe Gomez",
		"Andy Robertson",
		"Joel Matip",
		"Ki-Jana Hoever",
		"Trent Alexander-Arnold",
		"Sepp van den Berg",
		"Neco Williams",
		"Fabinho",
		"Georginio Wijnaldum",
		"James Milner",
		"Naby Keita",
		"Jordan Henderson",
		"Alex Oxlade-Chamberlain",
		"Xherdan Shaqiri",
		"Curtis Jones",
		"Harvel Elliott",
		"Roberto Firmino",
		"Sadio Mané",
		"Mohamed Salah",
		"Takumi Minamino",
		"Divock Origi",
	}

	// Unsorted
	fmt.Println(squad[0])

	// Initials
	SortStrings(squad, "taa")
	fmt.Println(squad[0])

	// Initials beat start of string
	SortStrings(squad, "al")
	fmt.Println(squad[0])

	// Start of word
	SortStrings(squad, "ox")
	fmt.Println(squad[0])

	// Earlier in string = better match
	SortStrings(squad, "x")
	fmt.Println(squad[0])

	// Diacritics ignored when query is ASCII
	SortStrings(squad, "mane")
	fmt.Println(squad[0])

	// But not if query isn't
	SortStrings(squad, "né")
	fmt.Println(squad[0])
	SortStrings(squad, "ne")
	fmt.Println(squad[0])

	// Output:
	// Alisson Becker
	// Trent Alexander-Arnold
	// Andy Lonergan
	// Alex Oxlade-Chamberlain
	// Xherdan Shaqiri
	// Sadio Mané
	// Sadio Mané
	// Neco Williams
}
