package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JOSIAHTHEPROGRAMMER/Smart-Commit/internal/ai"
	"github.com/JOSIAHTHEPROGRAMMER/Smart-Commit/internal/config"
	"github.com/JOSIAHTHEPROGRAMMER/Smart-Commit/internal/git"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var version = "0.1.0"

// Flags
var (
	dryRun      bool
	copyToClip  bool
	commitType  string
	commitScope string
)

// Color definitions
var (
	successColor = color.New(color.FgGreen, color.Bold)
	errorColor   = color.New(color.FgRed, color.Bold)
	infoColor    = color.New(color.FgCyan)
	warningColor = color.New(color.FgYellow)
	headerColor  = color.New(color.FgMagenta, color.Bold)
)

var rootCmd = &cobra.Command{
	Use:   "smartcommit",
	Short: "AI-powered git commit message generator",
	Long:  `SmartCommit uses AI to generate conventional commit messages based on your staged changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Load environment variables
		if err := config.LoadEnv(); err != nil {
			warningColor.Printf("Warning: %v\n", err)
		}

		// Load config file
		cfg, err := config.LoadConfig()
		if err != nil {
			warningColor.Printf("Warning: %v\n", err)
			cfg = config.DefaultConfig()
		}

		// Check if in git repo
		if !git.IsGitRepo() {
			errorColor.Println("Error: Not a git repository")
			os.Exit(1)
		}

		// Check for staged changes
		if !git.HasStagedChanges() {
			errorColor.Println("No staged changes found")
			infoColor.Println("Use 'git add <files>' to stage changes first")
			os.Exit(1)
		}

		// Get the diff
		diff, err := git.GetStagedDiff()
		if err != nil {
			errorColor.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		// Filter sensitive data before sending to AI
		diff = git.FilterSensitiveData(diff)

		// Preview diff stats
		successColor.Printf("✓ Found staged changes (%d characters)\n", len(diff))
		infoColor.Println("\nGenerating commit message...")

		// Generate commit message using AI
		message, err := ai.GenerateCommitMessage(diff, cfg)
		if err != nil {
			errorColor.Printf("Error generating commit message: %v\n", err)
			os.Exit(1)
		}

		// Use default type from config if no flag provided
		if commitType == "" && cfg.DefaultType != "" {
			commitType = cfg.DefaultType
		}

		// Override type if flag is set or config has default
		if commitType != "" {
			message = overrideCommitType(message, commitType)
		}

		// Override scope if flag is set
		if commitScope != "" {
			message = overrideCommitScope(message, commitScope)
		}

		// Pretty print the commit message
		fmt.Println()
		headerColor.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		headerColor.Println("Generated Commit Message")
		headerColor.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

		// Format the message with colors
		lines := strings.Split(message, "\n")
		for i, line := range lines {
			if i == 0 {
				// First line (header) in bold green
				successColor.Println(line)
			} else if strings.TrimSpace(line) == "" {
				fmt.Println()
			} else {
				// Body lines in normal color
				fmt.Println(line)
			}
		}

		headerColor.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Println()

		// Handle --copy flag
		if copyToClip {
			err = clipboard.WriteAll(message)
			if err != nil {
				warningColor.Printf("Warning: Failed to copy to clipboard: %v\n", err)
			} else {
				successColor.Println("✓ Commit message copied to clipboard!")
			}
		}

		// Handle --dry-run flag
		if dryRun {
			infoColor.Println("Dry run mode - no commit created.")
			return
		}

		// Ask for confirmation
		fmt.Print("Commit with this message? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			errorColor.Printf("Error reading input: %v\n", err)
			os.Exit(1)
		}

		response = strings.TrimSpace(strings.ToLower(response))

		if response == "y" || response == "yes" {
			err = git.Commit(message)
			if err != nil {
				errorColor.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			successColor.Println("\n✓ Commit successful!")
		} else {
			warningColor.Println("\nCommit cancelled.")
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Version = version

	// Add flags
	rootCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Print commit message without committing")
	rootCmd.Flags().BoolVarP(&copyToClip, "copy", "c", false, "Copy commit message to clipboard")
	rootCmd.Flags().StringVarP(&commitType, "type", "t", "", "Force commit type (feat, fix, chore, etc.)")
	rootCmd.Flags().StringVarP(&commitScope, "scope", "s", "", "Force commit scope")
}

// overrideCommitType replaces the commit type in the message
func overrideCommitType(message, newType string) string {
	lines := strings.Split(message, "\n")
	if len(lines) == 0 {
		return message
	}

	// Parse first line: type(scope): title or type: title
	firstLine := lines[0]

	if strings.Contains(firstLine, "(") {
		// Has scope: type(scope): title
		scopeStart := strings.Index(firstLine, "(")
		rest := firstLine[scopeStart:]
		lines[0] = newType + rest
	} else if strings.Contains(firstLine, ":") {
		// No scope: type: title
		colonIdx := strings.Index(firstLine, ":")
		lines[0] = newType + firstLine[colonIdx:]
	}

	return strings.Join(lines, "\n")
}

// overrideCommitScope adds or replaces the scope in the message
func overrideCommitScope(message, newScope string) string {
	lines := strings.Split(message, "\n")
	if len(lines) == 0 {
		return message
	}

	firstLine := lines[0]

	if strings.Contains(firstLine, "(") {
		// Already has scope, replace it
		typeEnd := strings.Index(firstLine, "(")
		scopeEnd := strings.Index(firstLine, ")")
		commitType := firstLine[:typeEnd]
		rest := firstLine[scopeEnd+1:]
		lines[0] = fmt.Sprintf("%s(%s)%s", commitType, newScope, rest)
	} else if strings.Contains(firstLine, ":") {
		// No scope, add it
		colonIdx := strings.Index(firstLine, ":")
		commitType := firstLine[:colonIdx]
		rest := firstLine[colonIdx:]
		lines[0] = fmt.Sprintf("%s(%s)%s", commitType, newScope, rest)
	}

	return strings.Join(lines, "\n")
}
