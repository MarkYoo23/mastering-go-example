package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

// 위 폴더 지원을 안함.
//
//go:embed static
var f embed.FS

var searchString string

func walkFunction(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Printf("Path=%q, isDir=%v\n", path, d.IsDir())
	return nil
}

func walkSearch(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if d.Name() == searchString {
		fileInfo, err := fs.Stat(f, path)
		if err != nil {
			return err
		}
		fmt.Println("Found", path, "with size", fileInfo.Size())
		return nil
	}
	return nil
}

func list(f embed.FS) error {
	return fs.WalkDir(f, ".", walkFunction)
}

func search(f embed.FS) error {
	return fs.WalkDir(f, ".", walkSearch)
}

func extract(f embed.FS, filepath string) ([]byte, error) {
	return fs.ReadFile(f, filepath)
}

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
	// At this point we do not know what is included in ./static

	// 임베딩과 결합하여, 파일 리스트 호출이 가능합니다.
	err := list(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 임베딩과 결합하여, 파일 검색을 파일 명으로 가능합니다.
	searchString = "file.txt"
	err = search(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 파일을 불러올 수 있습니다.
	buffer, err := extract(f, "static/file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Save it to an actual file.txt
	curPath, _ := os.Getwd()

	// 실제로 os.File 을 사용하는 경우... windows/linux 기반 시스템에서 같이 사용하지 못함
	err = writeToFile(buffer, curPath+"\\static\\IOFS.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}
