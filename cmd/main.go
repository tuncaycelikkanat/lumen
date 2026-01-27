package main

import (
	"fmt"
	"lumen/internal/config"
	"lumen/internal/monitor"
	"lumen/internal/report"
	"time"
)

func main() {

	fmt.Println(`
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓██████████████▓▒░░▒▓████████▓▒░▒▓███████▓▒░  
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓██████▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓████████▓▒░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░                                              
		Log Unification and Monitoring Engine`)

	cfg, _ := config.LoadConfig("rules.yaml")
	events := make(chan monitor.Event)
	var captured []monitor.Event

	go func() {
		for e := range events {
			captured = append(captured, e)
			fmt.Printf("\n[WARN] %s is detected! Message: %s\n", e.RuleName, e.Message)
		}
	}()

	for {
		fmt.Println("\n--- LUMEN CLI ---")
		fmt.Println("1. Start Monitoring")
		fmt.Println("2. Take Report (CSV)")
		fmt.Println("3. Exit")
		fmt.Print("Your input: ")

		var input int
		fmt.Scanln(&input)

		switch input {
		case 1:
			go monitor.StartTailing("auth.log", cfg.Rules, events)
			fmt.Println("Monitoring started...")
		case 2:
			timestamp := time.Now().Format("20060102_150405")
			fileName := fmt.Sprintf("report_%s.csv", timestamp)
			report.ExportToCSV(fileName, captured)
			fmt.Println("Raport saved.")
		case 3:
			return
		}
	}

}
