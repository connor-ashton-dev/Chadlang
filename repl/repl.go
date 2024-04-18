package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/connor-ashton-dev/chad/evaluator"
	"github.com/connor-ashton-dev/chad/lexer"
	"github.com/connor-ashton-dev/chad/object"
	"github.com/connor-ashton-dev/chad/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			fmt.Println("didn't scan")
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			_, err := io.WriteString(out, evaluated.Inspect())
			if err != nil {
				return
			}
			_, err = io.WriteString(out, "\n")
			if err != nil {
				return
			}
		}
	}
}

const CHAD_FACE = ` 

   .------\ /------.
   |       -       |
   |               |
   |               |
   |               |
_______________________
===========.===========
  / ~~~~~     ~~~~~ \
 /|     |     |\
 W   ---  / \  ---   W
 \.      |o o|      ./
  |                 |
  \    #########    /
   \  ## ----- ##  /
    \##         ##/
     \_____v_____/

`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, CHAD_FACE)
	io.WriteString(out, "Shnikey! We ran into some errors here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
