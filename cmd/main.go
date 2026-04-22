package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"lumen/internal/config"
	"lumen/internal/monitor"
	"lumen/internal/report"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	configPath string
	targetLogs []string
	format     string
)

var rootCmd = &cobra.Command{
	Use:   "lumen",
	Short: "Lumen is a lightweight log rule evaluation engine",
	Long: `
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓██████████████▓▒░░▒▓████████▓▒░▒▓███████▓▒░  
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓██████▓▒░ ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓█▓▒░     ░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
░▒▓████████▓▒░▒▓██████▓▒░░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░                                              
		Log Unification and Monitoring Engine

Lumen monitors log files based on YAML rules.`,
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start monitoring logs",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig(configPath)
		if err != nil {
			color.Red("Error loading config: %v", err)
			os.Exit(1)
		}

		events := make(chan monitor.Event)
		var captured []monitor.Event

		// Start tailing for each target log file
		for _, target := range targetLogs {
			go monitor.StartTailing(target, cfg.Rules, events)
			color.Green("Started monitoring: %s", target)
		}

		// Handle graceful shutdown to save reports
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		color.Cyan("Lumen is running. Press Ctrl+C to stop and save report.")

		for {
			select {
			case e := <-events:
				captured = append(captured, e)
				printEvent(e)
			case <-sigs:
				color.Yellow("\nShutting down Lumen...")
				saveReport(captured)
				os.Exit(0)
			}
		}
	},
}

func printEvent(e monitor.Event) {
	// Format output based on severity
	severity := strings.ToLower(e.Severity)
	var printFunc func(format string, a ...interface{})

	switch severity {
	case "critical":
		printFunc = color.New(color.FgHiRed, color.Bold).PrintfFunc()
	case "high":
		printFunc = color.New(color.FgRed).PrintfFunc()
	case "warn", "warning":
		printFunc = color.New(color.FgYellow).PrintfFunc()
	case "info":
		printFunc = color.New(color.FgCyan).PrintfFunc()
	default:
		printFunc = color.New(color.FgWhite).PrintfFunc()
	}

	printFunc("[%s] [%s] %s - %s\n", e.Time, strings.ToUpper(e.Severity), e.RuleName, e.Message)
}

func saveReport(events []monitor.Event) {
	if len(events) == 0 {
		color.Yellow("No events captured. Skipping report generation.")
		return
	}

	timestamp := time.Now().Format("20060102_150405")
	var err error
	var fileName string

	if format == "json" {
		fileName = fmt.Sprintf("report_%s.json", timestamp)
		err = report.ExportToJSON(fileName, events)
	} else {
		fileName = fmt.Sprintf("report_%s.csv", timestamp)
		err = report.ExportToCSV(fileName, events)
	}

	if err != nil {
		color.Red("Failed to save report: %v", err)
	} else {
		color.Green("Report saved successfully: outputs/%s", fileName)
	}
}

func init() {
	startCmd.Flags().StringVarP(&configPath, "config", "c", "rules.yaml", "Path to rules configuration file")
	startCmd.Flags().StringSliceVarP(&targetLogs, "log", "l", []string{"auth.log"}, "Log file(s) to monitor. Can specify multiple.")
	startCmd.Flags().StringVarP(&format, "format", "f", "csv", "Report export format (csv or json)")

	rootCmd.AddCommand(startCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
