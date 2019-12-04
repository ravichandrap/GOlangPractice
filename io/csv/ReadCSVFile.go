package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
)

func readCSVFile(filePath string) [][] string {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	csvRecord := csv.NewReader(f)
	records, err := csvRecord.ReadAll()
	fmt.Println(records)
	fmt.Println("------------")
	if err != nil {
		fmt.Println(err)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i][0] < records[j][0]
	})
	return records
}

func main()  {
	recovers := readCSVFile("io/csv/myfile.csv")
	fmt.Println(recovers)
}
