package report

import (
	"encoding/csv"
	"fmt"
	"lumen/internal/monitor"
	"os"
	"path/filepath"
)

func ExportToCSV(filename string, events []monitor.Event) error {
	outputDir := "outputs"

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, 0755)
		if err != nil {
			return fmt.Errorf("Error Message: %v", err)
		}
	}

	filePath := filepath.Join(outputDir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Error Message: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Time", "File", "RuleName", "Message"})
	for _, e := range events {
		writer.Write([]string{e.Time, e.File, e.RuleName, e.Message})
	}

	return nil
}
