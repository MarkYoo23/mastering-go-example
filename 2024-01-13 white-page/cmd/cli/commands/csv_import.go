package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"white-page/internal/di"
	"white-page/internal/entries"
	"white-page/internal/envs"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: commandCsvImport,
		Run: func(cmd *cobra.Command, args []string) {

			service, err := di.GetService[entries.EntryRepository]()
			if err != nil {
				panic(fmt.Errorf("error : %v", err))
			}

			openFile, err := os.Open(envs.PhoneBookCsvLocation)
			if err != nil {
				panic(err)
			}

			var buf = make([]byte, 1024)

			_, err = openFile.Read(buf)
			if err != nil {
				panic(err)
			}

			err = openFile.Close()
			if err != nil {
				panic(err)
			}

			err = service.DeleteAll()
			if err != nil {
				panic(fmt.Errorf("db exec fail : %v", err))
			}

			line := strings.Split(string(buf), "\n")

			for _, v := range line {
				split := strings.Split(v, ",")
				if len(split) == 1 {
					break
				}

				_, err = service.Add(split[0], split[1], split[2])
				if err != nil {
					panic(fmt.Errorf("db error : %v", err))
				}
			}

			fmt.Println("import success")
		}})
}
