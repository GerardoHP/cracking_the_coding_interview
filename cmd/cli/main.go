package main

import (
	"fmt"

	"github.com/GerardoHP/cracking_the_coding_interview/cmd/cli/commands"
	"os"
)

// todo: Pending all the validation errors

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		_, _ = fmt.Fprintf(os.Stderr, "You must pass a sub-command")
		os.Exit(1)
	}

	var cmd commands.CommandRun
	switch args[0] {
	case commands.IsUniqueCommandName:
		cmd = commands.NewIsUniqueCommand()
		_ = cmd.Init(args[1:])
		break
	case commands.CheckPermutationName:
		cmd = commands.NewCheckPermutationCommand()
		_ = cmd.Init(args[1:])
		break
	case commands.UrlifyName:
		cmd = commands.NewUrlifyComamnd()
		_ = cmd.Init(args[1:])
		break
	case commands.PalindromePermutationName:
		cmd = commands.NewPalindromePermutationCommand()
		_ = cmd.Init(args[1:])
		break
	case commands.OneAwayCommandName:
		cmd = commands.NewOneAwayCommand()
		_ = cmd.Init(args[1:])
	case commands.StringCompressionCommandName:
		cmd = commands.NewStringCompressionCommand()
		_ = cmd.Init(args[1:])
	default:
		cmd = nil
	}

	if cmd == nil {
		_, _ = fmt.Fprintf(os.Stderr, "You must pass a valid command")
		os.Exit(1)
	}

	err := cmd.Run()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "There was an error of the execution of the command")
		os.Exit(1)
	}
}
