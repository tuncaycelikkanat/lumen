package report

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"lumen/internal/monitor"
	"os"
	"path/filepath"
)

// ensureOutputDir ensures the output directory exists.
func ensureOutputDir() (string, error) {
	outputDir := "outputs"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return "", fmt.Errorf("Error Message: %v", err)
		}
	}
	return outputDir, nil
}

// ExportToCSV saves the caught events to a CSV file.
func ExportToCSV(filename string, events []monitor.Event) error {
	outputDir, err := ensureOutputDir()
	if err != nil {
		return err
	}

	filePath := filepath.Join(outputDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error Message: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Time", "File", "RuleName", "Severity", "Message"})
	for _, e := range events {
		writer.Write([]string{e.Time, e.File, e.RuleName, e.Severity, e.Message})
	}

	return nil
}

// ExportToJSON saves the caught events to a JSON file.
func ExportToJSON(filename string, events []monitor.Event) error {
	outputDir, err := ensureOutputDir()
	if err != nil {
		return err
	}

	filePath := filepath.Join(outputDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error Message: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(events); err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	return nil
}
