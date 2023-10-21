package commands

import (
	"flag"
	"fmt"
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
)

const OneAwayCommandName = "OneAway"

var _ CommandRun = (*OneAwayCommand)(nil)

type OneAwayCommand struct {
	fs     *flag.FlagSet
	str1   string
	str2   string
	output bool
}

func NewOneAwayCommand() *OneAwayCommand {
	o := &OneAwayCommand{
		fs: flag.NewFlagSet(OneAwayCommandName, flag.ExitOnError),
	}

	o.fs.StringVar(&o.str1, "str1", "", "The string 1")
	o.fs.StringVar(&o.str2, "str2", "", "The string 2")
	return o
}

func (o *OneAwayCommand) Init(args []string) error {
	return o.fs.Parse(args)
}

func (o *OneAwayCommand) Run() error {
	o.output = chapter_01.OneAway(o.str1, o.str2)
	fmt.Printf("The result for %v, %v is %v \n", o.str1, o.str2, o.output)
	return nil
}

func (o *OneAwayCommand) Name() string {
	return o.fs.Name()
}
