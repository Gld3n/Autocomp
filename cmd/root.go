// Copyright © 2020 Gld3n gld3ndev30@hotmail.com
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// rootCmd represents the base command when called without any subcommands.
var (
	// Used for flags.
	author string

	rootCmd = &cobra.Command{
		Use:   "autocomp <command>",
		Short: "Autocomp will build your files automatically!",
		Long: `Autocomp is a package designed to build .go files in real time
without the need to do it manually. Autocomp will search for changing
files into your directory to save you some time!`,
	}
)

// Execute executes the root command.
func Execute() {
	// call the root command and verify there's no errors
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "Robert Vale",
		"author name for copyright attribution")
	err := viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	if err != nil {
		log.Fatal(err)
	}
	rootCmd.AddCommand(watchCmd)
}
