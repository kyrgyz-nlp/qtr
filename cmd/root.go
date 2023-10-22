package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:   "qtr",
	Short: "Kirill tekstterdi Qyrgyzdyn latyn tamgasyna transliterlöö",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {

}
