package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func RunREPL() {
	// Close the REPL on Ctrl-C
	SetupCloseHandler()

	fmt.Println("TinyDB REPL: Press ctrl+c to exit")
	printPrompt()

	var reader = bufio.NewReader(os.Stdin)
	var text = read(reader)

	for {
		fmt.Println(text)
		printPrompt()
		text = read(reader)
	}
}

func printPrompt() {
	fmt.Print("BabyDB> ")
}

func read(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nBye!")
		os.Exit(0)
	}()
}
