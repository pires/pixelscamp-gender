package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	c "github.com/hstove/gender/classifier"
)

type pixel struct {
	Name string `json: name`
}

type pixels []*pixel

type jsonPixels struct {
	Users pixels `json: users`
}

func percent(num, numTotal int) float64 {
	return (float64(num) / float64(numTotal) * 100)
}

// printStats calculates and prints out gender stats
func printStats(pixels pixels) {
	classifier := c.Classifier()
	var numFemale int
	var numMale int

	// Perform the magic stuff
	for _, pixel := range pixels {
		gender, _ := c.Classify(classifier, strings.Split(pixel.Name, " ")[0])
		switch gender {
		case string(c.Boy):
			numMale += 1
		case string(c.Girl):
			numFemale += 1
		}
	}

	// Make the math
	numTotal := len(pixels)
	numUnknown := numTotal - numFemale - numMale
	percentFemale := percent(numFemale, numTotal)
	percentMale := percent(numMale, numTotal)
	percentUnknown := (100 - percentFemale - percentMale)

	// Print stats
	fmt.Printf(" > Total: %d\n", numTotal)
	fmt.Printf(" > Female: %d (%.2f%%)\n", numFemale, percentFemale)
	fmt.Printf(" > Male: %d (%.2f%%)\n", numMale, percentMale)
	fmt.Printf(" > Unknown: %d (%.2f%%)\n", numUnknown, percentUnknown)
}

func main() {
	// Retrieve Pixels Camp atendees from API
	url := "https://api.pixels.camp/users/?count=1000"
	res, err := http.Get(url)
	if err != nil {
		fmt.Errorf("There was an error while querying the API: %v", err)
	}

	// Decode response
	jsonParser := json.NewDecoder(res.Body)
	var pixels jsonPixels
	if err := jsonParser.Decode(&pixels); err != nil {
		fmt.Errorf("There was an error while reading the response: %v", err)
	}

	// Print stats
	fmt.Println("Pixels:")
	printStats(pixels.Users)
}
