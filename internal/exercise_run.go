package internal

type ExerciseRun interface {
	Init([]string) error
	Run() error
	Name() string
}
