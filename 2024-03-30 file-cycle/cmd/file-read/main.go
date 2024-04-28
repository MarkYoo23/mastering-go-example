package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetSize(path string) (int64, error) {
	// 하위 폴더 DIR 및 파일 리스트를 을 조회한다.
	// 이때 os.[]DirEntry 를 가져온다
	contents, err := os.ReadDir(path)
	if err != nil {
		return -1, err
	}

	// 하우
	var total int64
	for _, entry := range contents {
		// Visit directory entries
		if entry.IsDir() {
			// 폴더일 경우 재귀 호출하여 하위 폴더의 데이터를 계산해온다.
			temp, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return -1, err
			}
			total += temp
			// Get size of each non-directory entry
		} else {
			// 일반 파일일 경우 파일 값을 계산한다.
			info, err := entry.Info()
			if err != nil {
				return -1, err
			}
			// Returns an int64 value
			total += info.Size()
		}
	}
	return total, nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a <Directory>")
		return
	}

	// 파일의 진짜 위치를 가져온다. (심볼릭 링크일경우)
	root, err := filepath.EvalSymlinks(arguments[1])

	// 파일 정보를 가져온다.
	fileInfo, err := os.Stat(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 파일 정보로 디렉터리인지 확인한다.
	fileInfo, _ = os.Lstat(root)
	mode := fileInfo.Mode()
	if !mode.IsDir() {
		fmt.Println(root, "not a directory!")
		return
	}

	// 디렉터리 전체 크기를 계산한다.
	i, err := GetSize(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total Size:", i)
}
