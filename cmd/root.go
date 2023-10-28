package cmd

import (
	"bufio"
	"github.com/kyrgyz-nlp/qtr/text"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = cobra.Command{
	Use:   "qtr",
	Long:  "cat <FILE> | qtr",
	Short: "Kirill tekstterdi Qyrğyz-latyn tamğasyna transliterlöö",
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(text.ScanLinesWithLF)
		converter := text.NewConverter(scanner)
		converter.Convert(os.Stdout)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
