/*
Copyright © 2025 Dharmik Vivek Shinde dharmikvs26@gmail.com
*/

package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Constants for git config scopes
const (
	scopeGlobal = "--global"
	scopeLocal  = "--local"
	scopeSystem = "--system"
)

// configCmd allows users to set their Git identity interactively.
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure your Git identity (name and email) interactively",
	Long: `Interactively set your Git username and email for a chosen configuration scope.
Scopes:
  • Global - affects all repositories for the current user
  • Local  - affects only the current repository
  • System - affects all users on the system (rarely used)
`,
	Run: runConfigCmd,
}

func init() {
	rootCmd.AddCommand(configCmd)
}

// runConfigCmd is the main logic for the config command.
func runConfigCmd(cmd *cobra.Command, args []string) {
	fmt.Println("🔧 GitFx Config Mode — Set your Git identity")

	scope, scopeLabel, ok := promptForScope()
	if !ok {
		fmt.Println("❌ Configuration cancelled.")
		return
	}

	// Early check: if user wants local and not in a repo, abort before prompting further
	if scope == scopeLocal && !isInsideGitRepo() {
		fmt.Println("❌ You are not inside a Git repository. Local Git configuration requires an initialized repository.")
		fmt.Println("💡 Tip: Run `git init` first if you want to use local config here.")
		return
	}

	// Only prompt for values if repo state is valid
	username := promptForValue("Username (user.name)", "")
	email := promptForValue("Email (user.email)", "")

	if username == "" && email == "" {
		fmt.Println("⚠️  No values entered. Git config was not changed.")
		return
	}

	if username != "" {
		if err := setGitConfig(scope, "user.name", username); err != nil {
			fmt.Printf("❌ Failed to set user.name: %v\n", err)
		} else {
			fmt.Printf("✅ Set user.name = %s\n", username)
		}
	} else {
		fmt.Println("⚠️  Skipped setting user.name")
	}

	if email != "" {
		if err := setGitConfig(scope, "user.email", email); err != nil {
			fmt.Printf("❌ Failed to set user.email: %v\n", err)
		} else {
			fmt.Printf("✅ Set user.email = %s\n", email)
		}
	} else {
		fmt.Println("⚠️  Skipped setting user.email")
	}

	fmt.Printf("✅ Git identity updated using %s scope.\n", scopeLabel)
	fmt.Printf("🔍 You can verify with: git config %s --list\n", scope)
}

// promptForScope prompts the user to select a git config scope.
func promptForScope() (scopeFlag string, scopeLabel string, ok bool) {
	prompt := promptui.Select{
		Label: "Choose Git config scope",
		Items: []string{
			"Global (applies to all repos)",
			"Local (applies only to this repo)",
			"System (rarely used, for all users)",
		},
		Size: 3,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return "", "", false
	}
	switch result {
	case "Local (applies only to this repo)":
		return scopeLocal, "local", true
	case "System (rarely used, for all users)":
		return scopeSystem, "system", true
	default:
		return scopeGlobal, "global", true
	}
}

// promptForValue prompts the user for a value and trims whitespace.
func promptForValue(label, defaultValue string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
		Validate: func(input string) error {
			return nil
		},
	}
	value, err := prompt.Run()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(value)
}

// isInsideGitRepo checks if the current directory is inside a git repository.
func isInsideGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	out, err := cmd.Output()
	if err != nil || strings.TrimSpace(string(out)) != "true"{
		return false;
	}
	return true
}

// setGitConfig sets a git config value for the given scope and returns an error if any.
func setGitConfig(scope, key, value string) error {
	cmd := exec.Command("git", "config", scope, key, value)
	return cmd.Run()
}
