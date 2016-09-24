package main

import (
	"encoding/json"
	"fmt"
	"os"

	c "github.com/hstove/gender/classifier"
)

type pixel struct {
	Name string `json: name`
}

type pixels []*pixel

type jsonPixels struct {
	Users pixels `json: users`
}

// loadPixelsFromJSONFile loads PixelsCamp users from a JSON file
func loadPixelsFromJSONFile(filepath string) (pixels, error) {
	// Open JSON file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	// Load JSON
	jsonParser := json.NewDecoder(file)

	var jsonPixels jsonPixels
	if err := jsonParser.Decode(&jsonPixels); err != nil {
		return nil, err
	}

	return jsonPixels.Users, nil
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
		gender, _ := c.Classify(classifier, pixel.Name)
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
	// Retrieve first names
	filepath := "users.json"
	pixels, err := loadPixelsFromJSONFile(filepath)
	if err != nil {
		fmt.Errorf("There was an error while loading JSON from %s, %v", filepath, err)
	}

	// Print stats
	fmt.Println("Pixels:")
	printStats(pixels)
}
