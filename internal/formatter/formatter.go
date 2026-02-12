package formatter

import (
	"fmt"
	"strings"
)

// ConventionalCommit represents a structured commit message
type ConventionalCommit struct {
	Type           string
	Scope          string
	Title          string
	Body           []string
	BreakingChange bool
}

// Format returns the formatted conventional commit message
func (c *ConventionalCommit) Format() string {
	var result strings.Builder

	// Build header: type(scope): title
	if c.Scope != "" {
		result.WriteString(fmt.Sprintf("%s(%s): %s", c.Type, c.Scope, c.Title))
	} else {
		result.WriteString(fmt.Sprintf("%s: %s", c.Type, c.Title))
	}

	// Add body if present
	if len(c.Body) > 0 {
		result.WriteString("\n\n")
		for _, line := range c.Body {
			result.WriteString(fmt.Sprintf("- %s\n", line))
		}
	}

	// Add breaking change footer if needed
	if c.BreakingChange {
		result.WriteString("\nBREAKING CHANGE: This commit contains breaking changes")
	}

	return result.String()
}

// ParseCommitType validates and returns a conventional commit type
func ParseCommitType(typeStr string) string {
	validTypes := map[string]bool{
		"feat":     true,
		"fix":      true,
		"docs":     true,
		"style":    true,
		"refactor": true,
		"test":     true,
		"chore":    true,
		"perf":     true,
		"ci":       true,
		"build":    true,
		"revert":   true,
	}

	typeStr = strings.ToLower(strings.TrimSpace(typeStr))

	if validTypes[typeStr] {
		return typeStr
	}

	// Default to "chore" if unknown
	return "chore"
}

// TruncateTitle ensures the title is under max length
func TruncateTitle(title string, maxLen int) string {
	if len(title) <= maxLen {
		return title
	}
	return title[:maxLen-3] + "..."
}
