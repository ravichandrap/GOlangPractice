package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	f,err := os.Open("./fileNames.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csv := csv.NewWriter(f)
	defer csv.Flush()

	dir, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := dir.Readdir(0)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range fileInfo {
		line := "name: "+info.Name()
		csv.Write([]string{line})
	}
}
