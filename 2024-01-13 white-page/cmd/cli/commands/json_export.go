package commands

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"white-page/internal/di"
	"white-page/internal/entries"
	"white-page/internal/envs"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: commandJsonExport,
		Run: func(cmd *cobra.Command, args []string) {
			service, err := di.GetService[entries.EntryRepository]()
			if err != nil {
				panic(fmt.Errorf("error : %v", err))
			}

			results, err := service.GetAll()
			if err != nil {
				panic(fmt.Errorf("not found error : %v", err))
			}

			marshal, err := json.Marshal(results)
			if err != nil {
				return
			}

			file, err := os.OpenFile(envs.PhoneBookJsonLocation, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
			if err != nil {
				panic(fmt.Errorf("could not create phonebook.json: %v", err))
			}

			defer func(file *os.File) {
				_ = file.Close()
			}(file)

			_, err = file.Write(marshal)
			if err != nil {
				return
			}

			fmt.Println("export success")
		}})
}
