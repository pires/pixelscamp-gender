package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	c "github.com/hstove/gender/classifier"
	b "github.com/jbrukh/bayesian"
)

func worker(classifier *b.Classifier, filepath string) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		count, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			panic(err.Error())
		}

		name := strings.ToLower(record[0])
		idx := 0
		for idx <= int(count) {
			c.Learn(classifier, name, record[1])
			idx++
		}

	}
	fmt.Printf("Completed learning from %s.", filepath)
}

func main() {
	classifier := c.NewClassifier()
	worker(classifier, "names_pt_2014.csv")
	worker(classifier, "names_pt_2015.csv")
	worker(classifier, "names_uk.csv")
	worker(classifier, "names_us.csv")
	classifier.WriteToFile("../classifier.serialized")
}
