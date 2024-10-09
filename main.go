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
	// Shortest valid JSON is {} which is length 2 so anything less is INVALID
	if len(tokens) < 2 {
		return false
	}
	// Returning validity of smallest JSON 
	if tokens[0] == `{` && tokens[1] == `}` {
		return true
	}
	i := 0
	
	// Main JSON Parse
	for i < len(tokens) {
		for i < len(tokens) && tokens[i] == ` ` { // Skipping whitespace
			i++
		}

		if i == len(tokens) - 1 && tokens[i] == `}` { // Reached the end of the entire JSON
			return true
		} 

		if tokens[i] == `}` { // Reaching the end of a JSON Object value so move on to next token
			i++
		} else if tokens[i] == `{` || tokens[i] == `,`{ // Starting to parse the Entire JSON or a JSON Object value
			// or continuing to the next JSON key for parsing
			i++
			for i < len(tokens) && tokens[i] == ` ` { // Skipping Whitespace
				i++
			}
			if i < len(tokens) && tokens[i] == `"`  { // Starting to parse the JSON KEY

				// Validating the key
				i++
				for i < len(tokens) && tokens[i] == ` ` { // Skipping Whitespace
					i++
				}
				for i < len(tokens) && tokens[i] != `"` { // Iterating through the entire JSON key
					i++
					for i < len(tokens) && tokens[i] == ` ` { // Skipping Whitespace
						i++
					}
				}
				if i >= len(tokens) { // Checking if the JSON key is in the right format
					return false
				}

				// Validating the seperator
				i++
				for i < len(tokens) && tokens[i] == ` ` { // Skipping Whitespace
					i++
				}
				if i >= len(tokens) || tokens[i] != `:` {
					return false
				}

				// Validating the value
				i++
				for i < len(tokens) && tokens[i] == ` ` { // Skipping Whitespace
					i++
				}
				_, err := strconv.Atoi(tokens[i]) // Using this assignment to check if the token is a number
				if i+3 < len(tokens) && tokens[i] == `t` && tokens[i+1] == `r` && tokens[i+2] == `u` && tokens[i+3] == `e` { // Checking if the value is the boolean true
					i += 4

				} else if i+4 < len(tokens) && tokens[i] == `f` && tokens[i+1] == `a` && tokens[i+2] == `l` && tokens[i+3] == `s` && tokens[i+4] == `e` { // Checking if the value is the boolean false
					i += 5
			
				} else if i+3 < len(tokens) && tokens[i] == `n` && tokens[i+1] == `u` && tokens[i+2] == `l` && tokens[i+3] == `l`  { // Checking if the value is null
					i += 4
				} else if i < len(tokens) && tokens[i] == `"` { // Checking if the value is a string
					i++
					for i < len(tokens) && tokens[i] == ` ` {
						i++
					}
					for i < len(tokens) && tokens[i] != `"` {
						i++
						for i < len(tokens) && tokens[i] == ` ` {
							i++
						}
					}
					if i >= len(tokens) {
						return false
					}
					i++
					for i < len(tokens) && tokens[i] == ` ` {
						i++
					}
				} else if i < len(tokens) && err == nil { // Checking if the value is a number
					i++
					for i < len(tokens) && tokens[i] == ` ` {
						i++
					}
					_, err := strconv.Atoi(tokens[i])
					for i < len(tokens) && err == nil {
						i++
						for i < len(tokens) && tokens[i] == ` ` {
							i++
						}
						_, err = strconv.Atoi(tokens[i])
					}
					if i >= len(tokens) {
						return false
					}

				} else if i < len(tokens) && tokens[i] == `{` && validate(tokens[i:]) { // Checking if the value is a JSON Object
					return true
				} else if i < len(tokens) && tokens[i] == `[` { // Checking if the value is a list
					i++
					for {
						for i < len(tokens) && tokens[i] == ` ` {
							i++
						}
						_, err := strconv.Atoi(tokens[i])
						if i >= len(tokens) {
							return false
						}
						if tokens[i] == `"` { // Checking if the current list value is a string
							i++
							for i < len(tokens) && tokens[i] == ` ` {
								i++
							}
							for i < len(tokens) && tokens[i] != `"` {
								i++
								for i < len(tokens) && tokens[i] == ` ` {
									i++
								}
							}
							if i >= len(tokens) {
								return false
							}
							i++
							for i < len(tokens) && tokens[i] == ` ` {
								i++
							}
	
						} else if err == nil { // Checking if the current list value is a number
							i++
							for i < len(tokens) && tokens[i] == ` ` {
								i++
							}
							_, err := strconv.Atoi(tokens[i])
							for i < len(tokens) && err == nil {
								i++
								for i < len(tokens) && tokens[i] == ` ` {
									i++
								}
								_, err = strconv.Atoi(tokens[i])
							}
							if i >= len(tokens) {
								return false
							}
	
						} else if i+3 < len(tokens) && tokens[i] == `t` && tokens[i+1] == `r` && tokens[i+2] == `u` && tokens[i+3] == `e` { // Checking if the current list value is true
							i += 4
		
						} else if i+4 < len(tokens) && tokens[i] == `f` && tokens[i+1] == `a` && tokens[i+2] == `l` && tokens[i+3] == `s` && tokens[i+4] == `e` { // Checking if the current list value is false
							i += 5
					
						} else if i+3 < len(tokens) && tokens[i] == `n` && tokens[i+1] == `u` && tokens[i+2] == `l` && tokens[i+3] == `l`  { // Checking if the current list value is null
							i += 4
						}

						if tokens[i] == `,` { // Checking if there is more elements in the list
							continue
						} else if tokens[i] == `]` { // Checking if we're at the end of the list
							i++
							break
						} else { // List is not in the right format
							return false
						}

					}
					
				} else { // JSON Value is not in the right format
					return false
				}
			} else { // JSON Key is not in the right format
				return false
			}
		}  
		
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

	result := strings.Split(tokens, ``)
	
	
	return result

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
		fmt.Println("Step 2 valid2.json: Passed")
	} else {
		fmt.Println("Step 2 valid2.json: Failed")
	}
}

func step3(jsonStep3I1 []string, jsonStep3V1 []string) {
	if validate(jsonStep3I1) {
		fmt.Println("Step 3 invalid.json: Failed")
	} else {
		fmt.Println("Step 3 invalid.json: Passed")
	}

	if validate(jsonStep3V1) {
		fmt.Println("Step 3 valid.json: Passed")
	} else {
		fmt.Println("Step 3 valid.json: Failed")
	}

}

func step4(jsonStep4I1 []string, jsonStep4V1 []string, jsonStep4V2 []string) {
	if validate(jsonStep4I1) {
		fmt.Println("Step 4 invalid.json: Failed")
	} else {
		fmt.Println("Step 4 invalid.json: Passed")
	}

	if validate(jsonStep4V1) {
		fmt.Println("Step 4 valid.json: Passed")
	} else {
		fmt.Println("Step 4 valid.json: Failed")
	}

	if validate(jsonStep4V2) {
		fmt.Println("Step 4 valid2.json: Passed")
	} else {
		fmt.Println("Step 4 valid2.json: Failed")
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

	jsonStep3I1 := readingFile("tests/step3/invalid.json")
	jsonStep3V1 := readingFile("tests/step3/valid.json")

	step3(jsonStep3I1, jsonStep3V1)

	jsonStep4I1 := readingFile("tests/step4/invalid.json")
	jsonStep4V1 := readingFile("tests/step4/valid.json")
	jsonStep4V2 := readingFile("tests/step4/valid2.json")

	step4(jsonStep4I1, jsonStep4V1, jsonStep4V2)


}

// Sam Bakri