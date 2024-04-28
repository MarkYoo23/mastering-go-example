package envs

import "os"

var (
	PhoneBookCsvLocation  string = ".bin/phonebook.csv"
	PhoneBookJsonLocation string = ".bin/phonebook.json"
)

func init() {
	// 환경 변수가 없다면, 기본값으로 설정
	phonebookPath := os.Getenv("PHONEBOOK_CSV_PATH")
	if phonebookPath != "" {
		PhoneBookCsvLocation = phonebookPath
	}
}
