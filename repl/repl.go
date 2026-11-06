package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey/lexer"
	"monkey/token"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	for {
		_, _ = fmt.Fprint(out, Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return scanner.Err()
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, _ = fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
