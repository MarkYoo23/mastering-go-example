package main

import (
	"fmt"
	"mastering-go-example/cmd/multi_format_csv/reader"
	"sort"
)

func main() {
	readAndSort("./.resource/figure01.csv")
	readAndSort("./.resource/figure02.csv")
	readAndSort("./.resource/figureX.csv")
}

func readAndSort(path string) {
	csvReader := reader.NewCsvReader()

	figure, err := csvReader.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	// type 을 스위치로 사용하는 이유가 뭘까? 내부 소스를 좀 볼까?
	switch figure.(type) {
	case reader.Figure01s:
		sort.Sort(figure.(reader.Figure01s))
	case reader.Figure02s:
		sort.Sort(figure.(reader.Figure02s))
	}

	// 그냥 SORT 하면 안되나?

	fmt.Println(figure)
}
