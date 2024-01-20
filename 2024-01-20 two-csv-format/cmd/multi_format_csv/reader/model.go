package reader

import "strings"

type Figure01 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type Figure01s []Figure01

func (f Figure01s) Len() int {
	return len(f)
}

func (f Figure01s) Less(i, j int) bool {
	return strings.Compare(f[i].Name, f[j].Name) < 0
}

func (f Figure01s) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

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
