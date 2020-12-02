// Copyright © 2020 Gld3m gld3ndev30@hotmail.com
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var watchCmd = &cobra.Command{
	Use:     "watch [--run] <file>",
	Aliases: []string{"w"},
	Short:   "Start the auto-compiling process",
	Example: "watch -R main.go",
	Run: func(cmd *cobra.Command, filename []string) {
		// get the first argument
		f := filename[0]

		fmt.Printf("Started watching at %s\n", f)
		// Establish a watcher for the file we want to observe for changes
		watch(f)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ended running")
	},
}

// all the logic behind changing file monitoring
func watch(filename string) {
	fmt.Printf("watching: %s\n", filename)
}
