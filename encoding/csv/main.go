package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	saveCSV("test.csv", records)

	rs := new([][]string)
	loadCSV("test.csv", rs)
	for _, r := range *rs {
		fmt.Println(r)
	}
}

func saveCSV(fileName string, records [][]string) {
	outFile, err := os.Create(fileName)
	checkError(err)
	w := csv.NewWriter(outFile)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func loadCSV(fileName string, records *[][]string)  {
	inFile, err := os.Open(fileName)
	checkError(err)

	r := csv.NewReader(inFile)
	*records, err = r.ReadAll()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}
