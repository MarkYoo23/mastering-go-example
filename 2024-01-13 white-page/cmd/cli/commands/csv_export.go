package commands

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"white-page/internal/di"
	"white-page/internal/entries"
	"white-page/internal/envs"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: commandExport,
		Run: func(cmd *cobra.Command, args []string) {

			service, err := di.GetService[entries.EntryRepository]()
			if err != nil {
				panic(fmt.Errorf("error : %v", err))
			}

			results, err := service.GetAll()
			if err != nil {
				panic(fmt.Errorf("not found error : %v", err))
			}

			// CSV 파일 생성
			file, err := os.Open(envs.PhoneBookCsvLocation)
			if err != nil {
				panic(fmt.Errorf("could not create CSV file: %v", err))
			}

			defer func(file *os.File) {
				_ = file.Close()
			}(file)

			writer := csv.NewWriter(file)
			defer writer.Flush()

			for _, result := range results {
				if err = writer.Write([]string{result.Name, result.Surname, result.Tel}); err != nil {
					panic(fmt.Errorf("could not write to CSV file: %v", err))
				}
			}

			fmt.Println("export success")
		}})
}
