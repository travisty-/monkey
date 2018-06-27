package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/travisty-/monkey/lexer"
	"github.com/travisty-/monkey/token"
)

const PROMPT = ">> "

// Start initializes the REPL that tokenizes Monkey
// source code, and prints the tokens.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
