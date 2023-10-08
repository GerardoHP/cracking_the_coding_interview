package internal

type StringBuilder struct {
	array []byte
}

func (builder *StringBuilder) Append(str string) {
	builder.array = append(builder.array, []byte(str)...)
}

func (builder *StringBuilder) String() string {
	return string(builder.array)
}
