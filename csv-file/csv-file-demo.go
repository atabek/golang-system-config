package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"os"
)

func main() {
	csvFile, err := ioutil.ReadFile("./file1.csv")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	r := csv.NewReader(strings.NewReader(string(csvFile)))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

