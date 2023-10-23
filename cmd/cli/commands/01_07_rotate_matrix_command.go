package commands

import "flag"

const RotateMatrixCommandName = "RotateMatrix"

var _ CommandRun = (*RotateMatrixCommand)(nil)

type RotateMatrixCommand struct {
	fs     *flag.FlagSet
	input  [][]int
	output [][]int
}

func NewRotateMatrixCommand() *RotateMatrixCommand {
	rm := &RotateMatrixCommand{
		fs:    flag.NewFlagSet(RotateMatrixCommandName, flag.ExitOnError),
		input: make([][]int, 0),
	}

	return rm
}

func (r *RotateMatrixCommand) Init(args []string) error {
	//TODO implement me, implement to read from stdin a file with the slices
	panic("implement me")
}

func (r *RotateMatrixCommand) Run() error {
	//TODO implement me
	panic("implement me")
}

func (r *RotateMatrixCommand) Name() string {
	//TODO implement me
	panic("implement me")
}
