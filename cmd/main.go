package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/sgrilux/twitapp/commands"
)

func main() {

	parser := flags.NewParser(&commands.TwitApp, flags.HelpFlag|flags.PassDoubleDash)
	parser.NamespaceDelimiter = "-"

	helpParser := flags.NewParser(&commands.TwitApp, flags.HelpFlag)
	helpParser.NamespaceDelimiter = "-"

	_, err := parser.Parse()
	if err != nil {
		if err2 := handleError(helpParser, err); err2 != nil {
			os.Exit(0)
		}
	}
}

func handleError(helpParser *flags.Parser, err error) error {
	if err == commands.ErrShowHelpMessage {
		return showHelp(helpParser)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	return nil
}

func showHelp(helpParser *flags.Parser) error {
	_, err := helpParser.ParseArgs([]string{"-h"})
	if err != nil {
		return fmt.Errorf("error parsing help")
	}
	helpParser.WriteHelp(os.Stdout)
	return nil
}
