package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/lsnow99/conductorr/csl"
)

func main() {
	fmt.Println("Welcome to Conductorr REPL. Enter a csl expression and press enter for evaluation. Outermost parentheses will be inserted automatically if omitted. Ctrl+C to exit.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println()
			os.Exit(0)
		}
	}()

	env := make(map[string]interface{})

	for {
		fmt.Print(">> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		script := strings.TrimSpace(scanner.Text())
		if script[0] != '(' {
			script = "(" + script + ")"
		}
		exprs, err := csl.Parse(script)
		if err != nil {
			fmt.Println(err)
			continue
		}
		res, trace := csl.Eval(exprs, env)
		if trace.Err != nil {
			fmt.Println(trace.Err)
			continue
		}
		if res != nil {
			fmt.Println(res)
		}
	}
}
