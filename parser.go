package main

// ASTNode represents a node in an abstract syntax tree
type ASTNode struct {
	nodeType string      // Node type (literal, call expression, program)
	value    interface{} // Node value (literal value, function name)
	params   []ASTNode   // Slice of parameters to function call
	body     []ASTNode   // Slice of nodes that make up a program
}

var walk func() ASTNode

// Parser converts a slice of tokens into an abstract syntax tree
func Parser(tokens []Token) ASTNode {
	current := 0

	walk = func() ASTNode {
		token := tokens[current]

		// Number literals
		if token.tokenType == "number" {
			current++

			return ASTNode{
				"NumberLiteral",
				token.value,
				nil,
				nil,
			}
		}

		// String literals
		if token.tokenType == "string" {
			current++

			return ASTNode{
				"StringLiteral",
				token.value,
				nil,
				nil,
			}
		}

		// Call expressions
		if token.tokenType == "paren" && token.value == "(" {
			current++
			token = tokens[current]

			node := ASTNode{
				"CallExpression",
				token.value,
				make([]ASTNode, 0, 10),
				nil,
			}

			current++
			token = tokens[current]

			for token.tokenType != "paren" || token.tokenType == "paren" && token.value != ")" {
				node.params = append(node.params, walk())
				token = tokens[current]
			}

			current++

			return node
		}

		panic(token.tokenType)
	}

	ast := ASTNode{
		"Program",
		nil,
		nil,
		make([]ASTNode, 0, 10),
	}

	for current < len(tokens) {
		ast.body = append(ast.body, walk())
	}

	return ast
}
