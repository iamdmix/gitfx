/*
Copyright © 2025 Dharmik Vivek Shinde <dharmikvs26@gmail.com>
*/

package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var version = "0.1.0" // Current GitFx CLI version

var rootCmd = &cobra.Command{
    Use:     "gix",
    Short:   "GitFx - A friendlier Git CLI",
    Long: `GitFx is a modern, user-friendly command-line interface for Git.
Powered by GitFx, it streamlines your workflow and enhances productivity.

Use 'gix help' to explore available commands and options.`,
    Version: version,
    Run: func(cmd *cobra.Command, args []string) {
        displayWelcomeMessage()
    },
}

func Execute() {
    rootCmd.SetVersionTemplate("🎯 GitFx version: {{.Version}}\n")

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

func displayWelcomeMessage() {
    asciiArt := `
 $$$$$$\  $$\   $$\     $$$$$$$$\       
$$  __$$\ \__|  $$ |    $$  _____|      
$$ /  \__|$$\ $$$$$$\   $$ |  $$\   $$\ 
$$ |$$$$\ $$ |\_$$  _|  $$$$$\\$$\ $$  |
$$ |\_$$ |$$ |  $$ |    $$  __|\$$$$  / 
$$ |  $$ |$$ |  $$ |$$\ $$ |   $$  $$<  
\$$$$$$  |$$ |  \$$$$  |$$ |  $$  /\$$\ 
 \______/ \__|   \____/ \__|  \__/  \__|
`
    orange := "\033[38;5;208m"
    reset := "\033[0m"
    fmt.Println(orange + asciiArt + reset)
    fmt.Println("Welcome to GitFx — a modern, friendlier Git CLI.")
    fmt.Println("Type 'gix help' to view available commands and options.")
}
