package main

import (
	// 임베딩 패키지 직접 이용하진 않으나, 사용하므로 _로 참조
	_ "embed"
	"fmt"
	"os"
)

// 경로에 해당하는 파일을 임베딩
//
//go:embed static/image.jpg
var f1 []byte

//go:embed static/textfile
var f2 string

func writeToFile(s []byte, path string) error {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()

	n, err := fd.Write(s)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n)
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Print select 1|2")
		return
	}

	fmt.Println("f1:", len(f1), "f2:", len(f2))

	switch arguments[1] {
	case "1":
		curPath, _ := os.Getwd()
		filename := curPath + "\\static\\result.png"
		err := writeToFile(f1, filename)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "2":
		fmt.Print(f2)
	default:
		fmt.Println("Not a valid option!")
	}
}
