package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"white-page/internal/di"
	"white-page/internal/entries"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use: commandSearch,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := di.GetService[entries.EntryRepository]()
		if err != nil {
			panic(fmt.Errorf("error : %v", err))
		}

		surName := args[0]

		result, err := service.FindByPhone(surName)
		if err != nil {
			panic(fmt.Errorf("not found error : %v", err))
		}

		fmt.Println(fmt.Sprintf("entry : %v", result))
	},
}
