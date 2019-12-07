package main

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
)

func main() {
	path := "."
	export(readFiles(readDir(path)))
}

func readDir(path string) *os.File  {
	file, error := os.Open(path)
	if error != nil {
		log.Fatal("Error reading dir")
		return file
	}
	return file
}

func readFiles(file *os.File)   (names []string, err error) {
	if file == nil {
		log.Fatal("Error files")
		return nil, errors.New("os file is nil")
	}
	names, err = file.Readdirnames(0)
	return
}

func export(names []string, err error)  {

	if err != nil {
		log.Fatal("Error reading files ")
	}
	fileName := "./fileNames.csv"
	f, err := os.Open(fileName)
	if err != nil {

		f , err = os.Create(fileName)
		if err != nil {
			log.Fatal("ERROR while csv file creating")
		}
	}
	defer f.Close()
	fCSV := csv.NewWriter(f)
	defer fCSV.Flush()

	for _, value := range names {
		err:= fCSV.Write([]string{value})
		if err != nil {
			log.Fatal("Error while writing file", err)
		}
	}








}