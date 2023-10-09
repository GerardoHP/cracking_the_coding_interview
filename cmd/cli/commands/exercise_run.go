package commands

type CommandRun interface {
	Init([]string) error
	Run() error
	Name() string
}
