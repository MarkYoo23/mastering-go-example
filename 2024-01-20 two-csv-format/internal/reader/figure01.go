package reader

import (
	"fmt"
	"sort"
	"strings"
)

type Figure01 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type Figure01s []Figure01

func (fs Figure01s) Len() int {
	return len(fs)
}

func (fs Figure01s) Less(i, j int) bool {
	return strings.Compare(fs[i].Name, fs[j].Name) < 0
}

func (fs Figure01s) Swap(i, j int) {
	fs[i], fs[j] = fs[j], fs[i]
}

func (fs Figure01s) Sort() {
	sort.Sort(fs)
}

func (fs Figure01s) PrintConsole() {
	for i, f := range fs {
		fmt.Printf("%d : %s, %s, %s, %s\n", i, f.Name, f.Surname, f.Tel, f.LastAccess)
	}
}
