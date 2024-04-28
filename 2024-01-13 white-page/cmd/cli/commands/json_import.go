package commands

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"white-page/internal/di"
	"white-page/internal/entries"
)

//go:embed static
var f embed.FS

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: commandJsonImport,
		Run: func(cmd *cobra.Command, args []string) {
			file, err := fs.ReadFile(f, "phonebook.json")
			if err != nil {
				return
			}

			var es []entries.Entry
			err = json.Unmarshal(file, &es)
			if err != nil {
				return
			}

			service, err := di.GetService[entries.EntryRepository]()
			if err != nil {
				return
			}

			for _, element := range es {
				_, err = service.Add(element.Name, element.Surname, element.Tel)
				if err != nil {
					panic(fmt.Errorf("db error : %v", err))
				}
			}
		}})
}
