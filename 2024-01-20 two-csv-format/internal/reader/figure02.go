package reader

import (
	"fmt"
	"sort"
	"strings"
)

type Figure02 struct {
	Name       string
	Surname    string
	Areacode   string
	Tel        string
	LastAccess string
}

type Figure02s []Figure02

func (f Figure02s) Len() int {
	return len(f)
}

func (f Figure02s) Less(i, j int) bool {
	return strings.Compare(f[i].Name, f[j].Name) < 0
}

func (f Figure02s) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (fs Figure02s) Sort() {
	sort.Sort(fs)
}

func (fs Figure02s) PrintConsole() {
	for i, f := range fs {
		fmt.Printf("%d : %s, %s, %s, %s, %s\n", i, f.Name, f.Surname, f.Areacode, f.Tel, f.LastAccess)
	}
}
