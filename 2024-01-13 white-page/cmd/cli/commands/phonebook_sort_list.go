package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"white-page/internal/di"
	"white-page/internal/entries"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: commandSortList,
		Run: func(cmd *cobra.Command, args []string) {
			service, err := di.GetService[entries.EntryService]()
			if err != nil {
				panic(fmt.Errorf("error : %v", err))
			}

			results, err := service.GetSortByFullName()
			if err != nil {
				panic(fmt.Errorf("not found error : %v", err))
			}

			for _, result := range results {
				fmt.Println(fmt.Sprintf("entry : %v, %v, %v", result.Name, result.Surname, result.Tel))
			}
		}})
}
