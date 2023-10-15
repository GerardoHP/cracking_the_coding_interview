package commands

import (
	"flag"
	"fmt"
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
)

const PalindromePermutationName = "PalindromePermutation"

var _ CommandRun = (*PalindromePermutationCommand)(nil)

type PalindromePermutationCommand struct {
	fs     *flag.FlagSet
	input  string
	output bool
}

func NewPalindromePermutationCommand() *PalindromePermutationCommand {
	p := &PalindromePermutationCommand{
		fs: flag.NewFlagSet(PalindromePermutationName, flag.ExitOnError),
	}

	p.fs.StringVar(&p.input, "input", "", "The string to test")
	return p
}

func (p *PalindromePermutationCommand) Init(args []string) error {
	return p.fs.Parse(args)
}

func (p *PalindromePermutationCommand) Run() error {
	p.output = chapter_01.PalindromePermutation(p.input)
	fmt.Printf("The result for %v,  is %v \n", p.input, p.output)
	return nil
}

func (p *PalindromePermutationCommand) Name() string {
	return p.fs.Name()
}
