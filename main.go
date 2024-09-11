package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func validate(tokens []string) bool {
	if len(tokens) < 2 || tokens[0] != `{` || tokens[len(tokens) - 1] != `}` {
		return false
	}
	curr := tokens[0]
	i := 1
	for i < len(tokens) {
		token := tokens[i]
		if curr == `{` || curr == `,`{
			if string(token) == `"` && i+1 < len(tokens) {
				// Validating the key
				i++
				for i < len(tokens) && tokens[i] != `"` {
					i++
				}
				if tokens[i] != `"` {
					return false
				}

				// Validating the seperator
				i++
				if i >= len(tokens) || tokens[i] != `:` {
					return false
				}

				// Validating the value
				i++
				_, err := strconv.Atoi(tokens[i])
				if i < len(tokens) && tokens[i] == `t` {

				} else if i < len(tokens) && tokens[i] == `f` {

				} else if i < len(tokens) && tokens[i] == `n` {

				} else if i < len(tokens) && tokens[i] == `"` {

				} else if i < len(tokens) && err == nil {

				}
			}
		} 

		curr = tokens[i]
		i++
		
		
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

func step2(jsonStep2I1 []string, jsonStep2I2 []string, jsonStep2V1 []string, jsonStep2V2 []string) {
	if validate(jsonStep2I1) {
		fmt.Println("Step 2 invalid.json: Failed")
	} else {
		fmt.Println("Step 2 invalid.json: Passed")
	}

	if validate(jsonStep2I2) {
		fmt.Println("Step 2 invalid2.json: Failed")
	} else {
		fmt.Println("Step 2 invalid2.json: Passed")
	}

	if validate(jsonStep2V1) {
		fmt.Println("Step 2 valid.json: Passed")
	} else {
		fmt.Println("Step 2 valid.json: Failed")
	}

	if validate(jsonStep2V2) {
		fmt.Println("Step 2 valid.json: Passed")
	} else {
		fmt.Println("Step 2 valid.json: Failed")
	}
}

func main() {
	jsonStep1I1 := readingFile("tests/step1/invalid.json")
	jsonStep1V1 := readingFile("tests/step1/valid.json")

	step1(jsonStep1I1, jsonStep1V1)

	jsonStep2I1 := readingFile("tests/step2/invalid.json")
	jsonStep2I2 := readingFile("tests/step2/invalid2.json")
	jsonStep2V1 := readingFile("tests/step2/valid.json")
	jsonStep2V2 := readingFile("tests/step2/valid2.json")

	step2(jsonStep2I1, jsonStep2I2, jsonStep2V1, jsonStep2V2)

}