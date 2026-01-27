package monitor

import (
	"bufio"
	"fmt"
	"io"
	"lumen/internal/config"
	"os"
	"strings"
	"time"
)

type Event struct {
	Time     string
	File     string
	RuleName string
	Message  string
}

func StartTailing(filePath string, rules []config.Rule, results chan<- Event) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Hata: %s açılamadı\n", filePath)
		return
	}
	defer file.Close()

	// Dosyanın sonuna atla
	file.Seek(0, io.SeekEnd)
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}

		// Her satırı kurallarla kıyasla
		for _, rule := range rules {
			if strings.Contains(line, rule.Keyword) {
				results <- Event{
					Time:     time.Now().Format("15:04:05"),
					File:     filePath,
					RuleName: rule.Name,
					Message:  strings.TrimSpace(line),
				}
				break
			}
		}
	}
}
