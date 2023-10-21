package commands

import (
	"flag"
	"fmt"
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
)

const StringCompressionCommandName = "StringCompression"

var _ CommandRun = (*StringCompressionCommand)(nil)

type StringCompressionCommand struct {
	fs     *flag.FlagSet
	input  string
	output string
}

func NewStringCompressionCommand() *StringCompressionCommand {
	s := &StringCompressionCommand{
		fs: flag.NewFlagSet(StringCompressionCommandName, flag.ExitOnError),
	}

	s.fs.StringVar(&s.input, "input", "", "The string input")
	return s
}
func (s *StringCompressionCommand) Init(args []string) error {
	return s.fs.Parse(args)
}

func (s *StringCompressionCommand) Run() error {
	s.output = chapter_01.StringCompression(s.input)
	fmt.Printf("The result for %v is %v \n", s.input, s.output)
	return nil
}

func (s *StringCompressionCommand) Name() string {
	return s.fs.Name()
}
