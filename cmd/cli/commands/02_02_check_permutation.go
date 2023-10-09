package commands

import (
	"flag"
	"fmt"
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
)

const CheckPermutationName = "CheckPermutation"

var _ CommandRun = (*CheckPermutationCommand)(nil)

type CheckPermutationCommand struct {
	fs     *flag.FlagSet
	str1   string
	str2   string
	result bool
}

func NewCheckPermutationCommand() *CheckPermutationCommand {
	cp := &CheckPermutationCommand{
		fs: flag.NewFlagSet(CheckPermutationName, flag.ExitOnError),
	}

	cp.fs.StringVar(&cp.str1, "string1", "", "The first string to compare")
	cp.fs.StringVar(&cp.str2, "string2", "", "The second string to compare")
	return cp
}

func (c CheckPermutationCommand) Init(args []string) error {
	return c.fs.Parse(args)
}

func (c CheckPermutationCommand) Run() error {
	result := chapter_01.CheckPermutation(c.str1, c.str2)
	fmt.Printf("result for %v and %v is %v \n", c.str1, c.str2, result)
	c.result = result
	return nil
}

func (c CheckPermutationCommand) Name() string {
	return c.fs.Name()
}
