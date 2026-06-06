package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/titu/file-date-lunar-corr/internal/lunar"
	"github.com/titu/file-date-lunar-corr/internal/scanner"
	"github.com/titu/file-date-lunar-corr/internal/stats"
)

func main() {
	dir := flag.String("dir", ".", "Directory to scan")
	output := flag.String("output", "file_dates.csv", "Output CSV file")
	flag.Parse()

	fmt.Printf("Scanning directory: %s\n", *dir)
	files, err := scanner.Scan(*dir)
	if err != nil {
		log.Fatalf("Error scanning directory: %v", err)
	}

	fmt.Printf("Found %d files. Calculating lunar phases...\n", len(files))

	// Group counts by phase
	counts := make(map[lunar.Phase]int)
	for _, p := range lunar.PhaseNames {
		counts[p] = 0
	}

	// Prepare CSV data
	csvData := [][]string{
		{"Path", "Created Date", "Lunar Phase Value", "Lunar Phase Name"},
	}

	for _, f := range files {
		phaseName := lunar.GetPhase(f.Created)
		phaseValue := lunar.GetPhaseFloat(f.Created)
		counts[phaseName]++

		csvData = append(csvData, []string{
			f.Path,
			f.Created.Format(time.RFC3339),
			fmt.Sprintf("%.4f", phaseValue),
			string(phaseName),
		})
	}

	// Write CSV
	err = writeCSV(*output, csvData)
	if err != nil {
		log.Fatalf("Error writing CSV: %v", err)
	}
	fmt.Printf("Results saved to %s\n", *output)

	// Perform Statistical Analysis
	var observed []int
	for _, p := range lunar.PhaseNames {
		observed = append(observed, counts[p])
	}

	result := stats.PerformChiSquareTest(observed)

	fmt.Println("\n--- Lunar Phase Distribution ---")
	for _, p := range lunar.PhaseNames {
		fmt.Printf("%-20s: %d\n", p, counts[p])
	}

	fmt.Printf("\nChi-squared Statistic: %.4f\n", result.Statistic)
	fmt.Printf("P-value: %.4f\n", result.PValue)

	if result.PValue <= 0.05 {
		fmt.Println("Result: Statistically non-uniform — possible pattern!")
	} else {
		fmt.Println("Result: Likely random.")
	}
}

func writeCSV(path string, data [][]string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.WriteAll(data)
}
