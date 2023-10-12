package commands

import (
	"flag"
	"fmt"
	"github.com/GerardoHP/cracking_the_coding_interview/pkg/chapter_01"
)

const UrlifyName = "Urlify"

var _ CommandRun = (*UrlifyCommand)(nil)

type UrlifyCommand struct {
	fs     *flag.FlagSet
	input  string
	size   int
	output string
}

func NewUrlifyComamnd() *UrlifyCommand {
	uc := &UrlifyCommand{
		fs: flag.NewFlagSet(UrlifyName, flag.ExitOnError),
	}

	uc.fs.StringVar(&uc.input, "input", "", "The string to parse")
	uc.fs.IntVar(&uc.size, "size", -1, "The size of the original string")
	return uc
}

func (u *UrlifyCommand) Init(args []string) error {
	return u.fs.Parse(args)
}

func (u *UrlifyCommand) Run() error {
	u.output = chapter_01.Urlify([]rune(u.input), u.size)
	fmt.Printf("The result for %v, is %v \n", u.input, u.output)
	return nil
}

func (u *UrlifyCommand) Name() string {
	return u.fs.Name()
}
