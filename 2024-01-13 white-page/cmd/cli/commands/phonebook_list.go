package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"white-page/internal/di"
	"white-page/internal/entries"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: commandList,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := di.GetService[entries.EntryRepository]()
		if err != nil {
			panic(fmt.Errorf("error : %v", err))
		}

		results, err := service.GetAll()
		if err != nil {
			panic(fmt.Errorf("not found error : %v", err))
		}

		fmt.Println(fmt.Sprintf("entries : %v", results))
	},
}
