package commands

import (
	"flag"
	"fmt"
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
)

const IsUniqueCommandName = "IsUnique"

var _ CommandRun = (*IsUniqueCommand)(nil)

type IsUniqueCommand struct {
	fs     *flag.FlagSet
	entry  string
	result bool
}

func NewIsUniqueCommand() *IsUniqueCommand {
	iu := &IsUniqueCommand{
		fs: flag.NewFlagSet(IsUniqueCommandName, flag.ExitOnError),
	}

	iu.fs.StringVar(&iu.entry, "entry", "", "The string to test")
	return iu
}

func (i *IsUniqueCommand) Init(args []string) error {
	return i.fs.Parse(args)
}

func (i *IsUniqueCommand) Run() error {
	i.result = chapter_01.IsUnique(i.entry)
	fmt.Printf("result for %v is %v \n", i.entry, i.result)
	return nil
}

func (i *IsUniqueCommand) Name() string {
	return i.fs.Name()
}
