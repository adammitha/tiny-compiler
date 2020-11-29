package main

import (
	"errors"
	"regexp"
)

// Token represents a token type and value
type Token struct {
	tokenType string // The token type
	value     string // The token value
}

// Tokenizer converts source code into tokens
func Tokenizer(input string) ([]Token, error) {

	var tokens []Token

	WHITESPACE := regexp.MustCompile("\\s")
	NUMBERS := regexp.MustCompile("[0-9]")
	LETTERS := regexp.MustCompile("[a-z]")

	for current := 0; current < len(input); {
		char := string(input[current])

		// Open parenthesis
		if char == "(" {
			tokens = append(tokens, Token{"paren", "("})
			current++
			continue
		}

		// Close parenthesis
		if char == ")" {
			tokens = append(tokens, Token{"paren", ")"})
			current++
			continue
		}

		// Whitespace
		if WHITESPACE.Match([]byte(char)) {
			current++
			continue
		}

		// Numbers
		if NUMBERS.Match([]byte(char)) {
			// Keep a value variable for building up our number token
			var value string

			for NUMBERS.Match([]byte(char)) {
				value += char
				current++
				char = string(input[current])
			}

			tokens = append(tokens, Token{"number", value})

			continue
		}

		// Strings
		if char == "\"" {
			// Keep a value variable for building up our string token
			var value string

			current++
			char = string(input[current])

			for char != "\"" {
				value += char
				current++
				char = string(input[current])
			}

			tokens = append(tokens, Token{"string", value})

			continue
		}

		// Function names
		// (add 2 4)
		//  ^^^
		if LETTERS.Match([]byte(char)) {
			// Keep a value variable for building up our string token
			var value string

			for LETTERS.Match([]byte(char)) {
				value += char
				current++
				char = string(input[current])
			}

			tokens = append(tokens, Token{"name", value})

			continue
		}

		// If we have not matched a character by now, return an error
		return nil, errors.New("I don't know what this character is: " + char)

	}

	return tokens, nil
}
