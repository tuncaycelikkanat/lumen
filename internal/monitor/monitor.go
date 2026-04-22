package monitor

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"lumen/internal/config"

	"github.com/nxadm/tail"
)

// Event represents a single detected log anomaly.
type Event struct {
	Time     string
	File     string
	RuleName string
	Severity string
	Message  string
}

// compiledRule holds a rule along with its compiled regex for performance.
type compiledRule struct {
	config.Rule
	RegexPattern *regexp.Regexp
}

// StartTailing begins monitoring the specified log file using nxadm/tail.
func StartTailing(filePath string, rules []config.Rule, results chan<- Event) {
	// Pre-compile regex patterns for efficiency
	var compiledRules []compiledRule
	for _, r := range rules {
		cr := compiledRule{Rule: r}
		if r.Regex != "" {
			re, err := regexp.Compile(r.Regex)
			if err != nil {
				fmt.Printf("Warning: Failed to compile regex for rule %s: %v\n", r.Name, err)
			} else {
				cr.RegexPattern = re
			}
		}
		compiledRules = append(compiledRules, cr)
	}

	// Tail the file. Follow = true keeps tailing. ReOpen = true handles log rotation.
	t, err := tail.TailFile(filePath, tail.Config{
		Follow: true,
		ReOpen: true,
		Poll:   true, // Useful for systems where inotify doesn't work well
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2, // io.SeekEnd
		},
	})

	if err != nil {
		fmt.Printf("Error: Cannot start tailing file %s: %v\n", filePath, err)
		return
	}

	// Read lines as they come
	for line := range t.Lines {
		if line.Err != nil {
			continue // Skip errors
		}

		text := line.Text
		for _, cr := range compiledRules {
			matched := false

			// Check Regex if present
			if cr.RegexPattern != nil {
				if cr.RegexPattern.MatchString(text) {
					matched = true
				}
			} else if cr.Keyword != "" {
				// Fallback to Keyword
				if strings.Contains(text, cr.Keyword) {
					matched = true
				}
			}

			if matched {
				results <- Event{
					Time:     time.Now().Format("15:04:05"),
					File:     filePath,
					RuleName: cr.Name,
					Severity: cr.Severity,
					Message:  strings.TrimSpace(text),
				}
				break // Only trigger one rule per line to avoid spam
			}
		}
	}
}
