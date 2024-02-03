package main

import (
	"fmt"
)

func main() {
	readAndSort("./.resource/figure01.csv")
	readAndSort("./.resource/figure02.csv")
	readAndSort("./.resource/figureX.csv")
}

func readAndSort(path string) {
	csvReader := NewCsvReader()

	f, err := csvReader.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	f.Sort()
	f.PrintConsole()
}
