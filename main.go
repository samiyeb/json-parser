package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func validate(tokens []string) bool {
	if len(tokens) < 2 || tokens[0] != `{` || tokens[len(tokens) - 1] != `}` {
		return false
	}

	return true
}

func readingFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(file)
	tokens := ""
	for reader.Scan() {
		tokens += reader.Text()
	}

	return strings.Split(tokens, "")

}

func step1(jsonStep1I1 []string, jsonStep1V1 []string) {
	if validate(jsonStep1I1) {
		fmt.Println("Step 1 invalid.json: Failed")
	} else {
		fmt.Println("Step 1 invalid.json: Passed")
	}

	if validate(jsonStep1V1) {
		fmt.Println("Step 1 valid.json: Passed")
	} else {
		fmt.Println("Step 1 valid.json: Failed")
	}

}

func main() {
	jsonStep1I1 := readingFile("tests/step1/invalid.json")
	jsonStep1V1 := readingFile("tests/step1/valid.json")

	step1(jsonStep1I1, jsonStep1V1)

	

}