// A small compiler that converts Lisp into C syntax
// e.g. (add 1 2) -> add(1,2)
package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Welcome to the tiny compiler!")
}

// Token represents a token type and value
type Token struct {
	tokenType string // The token type
	value     string // The token value
}

func tokenizer(input string) []Token {

	var tokens []Token

	WHITESPACE := regexp.MustCompile("\\s")
	NUMBERS := regexp.MustCompile("[0-9")

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
		}
	}

	return tokens
}
