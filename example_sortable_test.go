// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package fuzzy

import (
	"fmt"
	"strings"
)

// Player is a very simple data model.
type Player struct {
	Firstname string
	Lastname  string
}

// Name returns the full name of the Player.
func (p *Player) Name() string {
	return strings.TrimSpace(p.Firstname + " " + p.Lastname)
}

// Team is a collection of Player items. This is where fuzzy.Sortable
// must be implemented to enable fuzzy sorting.
type Team []*Player

// Default sort.Interface methods
func (t Team) Len() int      { return len(t) }
func (t Team) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

// Less is used as a tie-breaker when fuzzy match score is the same.
func (t Team) Less(i, j int) bool { return t[i].Name() < t[j].Name() }

// Keywords implements Sortable.
// Comparisons are based on the the full name of the player.
func (t Team) Keywords(i int) string { return t[i].Name() }

// Fuzzy sort players by name.
func ExampleSort() {
	var t = Team{
		&Player{"Alisson", "Becker"},
		&Player{Firstname: "Adrián"},
		&Player{"Andy", "Lonergan"},
		&Player{"Caoimhín", "Kelleher"},
		&Player{"Virgil", "van Dijk"},
		&Player{"Joe", "Gomez"},
		&Player{"Andy", "Robertson"},
		&Player{"Joel", "Matip"},
		&Player{"Ki-Jana", "Hoever"},
		&Player{"Trent", "Alexander-Arnold"},
		&Player{"Sepp", "van den Berg"},
		&Player{"Neco", "Williams"},
		&Player{Firstname: "Fabinho"},
		&Player{"Georginio", "Wijnaldum"},
		&Player{"James", "Milner"},
		&Player{"Naby", "Keita"},
		&Player{"Jordan", "Henderson"},
		&Player{"Alex", "Oxlade-Chamberlain"},
		&Player{"Xherdan", "Shaqiri"},
		&Player{"Curtis", "Jones"},
		&Player{"Harvel", "Elliott"},
		&Player{"Roberto", "Firmino"},
		&Player{"Sadio", "Mané"},
		&Player{"Mohamed", "Salah"},
		&Player{"Takumi", "Minamino"},
		&Player{"Divock", "Origi"},
	}
	// Unsorted
	fmt.Println(t[0].Name())

	// Initials
	Sort(t, "taa")
	fmt.Println(t[0].Name())

	// Initials beat start of string
	Sort(t, "al")
	fmt.Println(t[0].Name())

	// Start of word
	Sort(t, "ox")
	fmt.Println(t[0].Name())

	// Earlier in string = better match
	Sort(t, "x")
	fmt.Println(t[0].Name())

	// Diacritics ignored if query is ASCII
	Sort(t, "mane")
	fmt.Println(t[0].Name())

	// But not if query isn't
	Sort(t, "né")
	fmt.Println(t[0].Name())
	Sort(t, "ne")
	fmt.Println(t[0].Name())

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
