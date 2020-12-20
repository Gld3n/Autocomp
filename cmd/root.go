//Package cmd provides the commands logic
// Copyright © 2020 Gld3n gld3ndev30@hotmail.com
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "autocomp <command>",
	Short: "Autocomp will build your files automatically!",
	Long: `Autocomp is a package designed to build .go files in real time
without the need to do it manually. Autocomp will search for changing
files into your directory to save you some time!`,
}

// Execute executes the root command.
func Execute() {
	// call the root command and verify there's no errors
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	/* author := rootCmd.PersistentFlags().BoolP("author", "a", false,
		"author name for copyright attribution")
	err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(author)
	if *author {
		fmt.Println("CLI made by Gld3n")
	} */
	rootCmd.AddCommand(watchCmd)
}
